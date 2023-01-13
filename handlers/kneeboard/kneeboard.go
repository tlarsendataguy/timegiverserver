package kneeboard

import (
	"database/sql"
	"fmt"
	"github.com/fogleman/gg"
	"image"
	"image/color"
	"io"
	"strings"
)

type Info struct {
	Name      string
	Elevation float64
}

type Frequency struct {
	Facility  string
	Label     string
	Frequency string
	Sort      int
	Comment   string
}

type Runway struct {
	Facility   string
	Rwy        string
	Identifier string
	X          float64
	Y          float64
}

type Kneeboard struct {
	From        string
	To          string
	Info        map[string]Info
	Frequencies []Frequency
	Runways     []Runway
}

func (k *Kneeboard) LoadInfo(rows *sql.Rows) error {
	k.Info = make(map[string]Info)
	var facility, name string
	var elevation float64
	for rows.Next() {
		err := rows.Scan(&facility, &name, &elevation)
		if err != nil {
			_ = rows.Close()
			return err
		}
		k.Info[facility] = Info{name, elevation}
	}
	return nil
}

func (k *Kneeboard) LoadFrequencies(rows *sql.Rows) error {
	k.Frequencies = make([]Frequency, 0)

	for rows.Next() {
		freq := Frequency{}
		err := rows.Scan(&freq.Facility, &freq.Label, &freq.Frequency, &freq.Sort, &freq.Comment)
		if err != nil {
			_ = rows.Close()
			return err
		}
		k.Frequencies = append(k.Frequencies, freq)
	}
	return nil
}

func (k *Kneeboard) LoadRunways(rows *sql.Rows) error {
	k.Runways = make([]Runway, 0)
	for rows.Next() {
		rwy := Runway{}
		err := rows.Scan(&rwy.Facility, &rwy.Rwy, &rwy.Identifier, &rwy.X, &rwy.Y)
		if err != nil {
			_ = rows.Close()
			return err
		}
		k.Runways = append(k.Runways, rwy)
	}
	return nil
}

func (k *Kneeboard) BuildHtml() []byte {
	str := strings.Replace(htmlPage, `{AIRPORT 1 NAME}`, k.Info[k.From].Name, 1)
	str = strings.Replace(str, `{AIRPORT 1 FREQUENCIES}`, k.fromFrequenciesHtml(), 1)
	str = strings.Replace(str, `{AIRPORT 1 RUNWAYS}`, fmt.Sprintf(rwyImageTemplate, k.From), 1)
	str = strings.Replace(str, `{AIRPORT 1 ID}`, k.From, 1)
	str = strings.Replace(str, `{AIRPORT 1 ELEVATION}`, fmt.Sprintf(`%.f`, k.Info[k.From].Elevation), 1)
	str = strings.Replace(str, `{AIRPORT 1 RUNWAY ENDS}`, k.fromRwyEndsHtml(), 1)

	str = strings.Replace(str, `{AIRPORT 2 NAME}`, k.Info[k.To].Name, 1)
	str = strings.Replace(str, `{AIRPORT 2 FREQUENCIES}`, k.toFrequenciesHtml(), 1)
	str = strings.Replace(str, `{AIRPORT 2 RUNWAYS}`, fmt.Sprintf(rwyImageTemplate, k.To), 1)
	str = strings.Replace(str, `{AIRPORT 2 ID}`, k.To, 1)
	str = strings.Replace(str, `{AIRPORT 2 ELEVATION}`, fmt.Sprintf(`%.f`, k.Info[k.To].Elevation), 1)
	str = strings.Replace(str, `{AIRPORT 2 RUNWAY ENDS}`, k.toRwyEndsHtml(), 1)

	return []byte(str)
}

func (k *Kneeboard) fromFrequenciesHtml() string {
	return k.frequenciesHtml(k.From)
}

func (k *Kneeboard) toFrequenciesHtml() string {
	return k.frequenciesHtml(k.To)
}

func (k *Kneeboard) frequenciesHtml(facility string) string {
	builder := &strings.Builder{}
	for _, freq := range k.Frequencies {
		if freq.Facility != facility {
			continue
		}
		str := fmt.Sprintf(freqTemplate, freq.Label, freq.Frequency, freq.Comment)
		builder.WriteString(str)
	}
	return builder.String()
}

func (k *Kneeboard) fromRwyEndsHtml() string {
	return k.rwyEndsHtml(k.From)
}

func (k *Kneeboard) toRwyEndsHtml() string {
	return k.rwyEndsHtml(k.To)
}

func (k *Kneeboard) rwyEndsHtml(facility string) string {
	builder := &strings.Builder{}
	for _, rwy := range k.Runways {
		if rwy.Facility != facility {
			continue
		}
		str := fmt.Sprintf(rwyEndTemplate, (rwy.X*rwyImageSize)+rwyImagePadding, (rwy.Y*rwyImageSize)+rwyImagePadding, rwy.Identifier)
		builder.WriteString(str)
	}
	return builder.String()
}

func (k *Kneeboard) CreateRwyImage(facility string, w io.Writer) error {
	rgba := image.NewRGBA(image.Rect(0, 0, rwyImageSize, rwyImageSize))
	ctx := gg.NewContextForRGBA(rgba)
	ctx.SetRGBA255(100, 100, 100, 255)
	ctx.SetLineWidth(4.0)

	runways := make([]Runway, 0)
	for _, runway := range k.Runways {
		if runway.Facility != facility {
			continue
		}
		runways = append(runways, runway)
	}

	transparent := color.RGBA{R: 1, G: 1, B: 1, A: 0}
	for x := 0; x < rwyImageSize; x++ {
		for y := 0; y < rwyImageSize; y++ {
			rgba.SetRGBA(x, y, transparent)
		}
	}

	for index := 0; index < len(runways); index += 2 {
		startX := runways[index].X * rwyImageSize
		startY := runways[index].Y * rwyImageSize
		endX := runways[index+1].X * rwyImageSize
		endY := runways[index+1].Y * rwyImageSize
		ctx.DrawLine(startX, startY, endX, endY)
		ctx.Stroke()
	}

	return ctx.EncodePNG(w)
}

const freqTemplate = `            <div class="frequency">
                <div class="frequency-name">%v</div>
                <div class="frequency-code">%v</div>
                <div class="frequency-comment">%v</div>
            </div>
`

const rwyEndTemplate = `            <div class="runway-id" style="left: %.fpx;top: %.fpx">%v</div>
`

const rwyImageTemplate = `            <img src="%v.png" class="runway-image">
`

const rwyImageSize = 150
const rwyImagePadding = 25

const htmlPage = `<html>
<head>
    <style>
		@font-face {
		  font-family: B612;
		  src: url(B612Mono-Regular.ttf);
		}

        body {
            width: 11.0in;
            height: 8.0in;
            display: flex;
            flex-direction: row;
            font-size: 11pt;
            margin: 0;
			font-family: B612;
        }

        p {
            margin: 5px 0;
        }

        .half-sheet {
            width: 50%;
            height: 100%;
            display: flex;
            flex-direction: column;
        }

        .apt-info {
            display: flex;
            flex-direction: row;
        }

        .notes {
            flex-grow: 1;
            display: flex;
            flex-direction: column;
        }

        .note-row {
            border-top: 1px solid black;
            flex-grow: 1;
        }

        .frequencies {
            flex-grow: 1;
            display: flex;
            flex-direction: column;
        }

        .frequency {
            border: 1px solid black;
            padding: 3px;
            margin: 2px 0;
            font-size: 9pt;
            display: flex;
            flex-direction: row;
        }

        .frequency-code {
            width: 80px
        }

        .frequency-name {
            width: 80px;
        }

        .frequency-comment {
            flex-grow: 1;
        }

        .image-box {
            width: 200px;
            height: 200px;
            font-size: 10pt;
            padding: 0 0 5px 5px;
            position: relative;
        }
        
        .runway-image {
            position: absolute;
            left: 25px;
            top: 25px;
        }

        .right-side {
            padding-left: 20px;
        }

        .left-side {
            padding-right: 20px;
        }

        .identifier {
            position: absolute;
            top: 0;
            right: 0;
        }

        .elevation {
            position: absolute;
            bottom: 0;
            right: 0;
        }

        .runway-id {
            position: absolute;
            font-size: 8pt;
            padding: 1px;
            background-color: white;
            border: 1px solid black;
        }
    </style>
</head>
<body>
<div class="half-sheet left-side">
    <div class="apt-info">
        <div class="frequencies">
            <p>{AIRPORT 1 NAME}</p>
{AIRPORT 1 FREQUENCIES}
        </div>
        <div class="image-box">
            <img src="wind-ring.png">
{AIRPORT 1 RUNWAYS}
            <div class="identifier">{AIRPORT 1 ID}</div>
            <div class="elevation">{AIRPORT 1 ELEVATION}</div>
{AIRPORT 1 RUNWAY ENDS}
        </div>
    </div>
    <div class="notes">
        <div class="note-row"></div>
        <div class="note-row"></div>
        <div class="note-row"></div>
        <div class="note-row"></div>
        <div class="note-row"></div>
    </div>
</div>
<div class="half-sheet right-side">
    <div class="apt-info">
        <div class="frequencies">
            <p>{AIRPORT 2 NAME}</p>
{AIRPORT 2 FREQUENCIES}
        </div>
        <div class="image-box">
            <img src="wind-ring.png">
{AIRPORT 2 RUNWAYS}
            <div class="identifier">{AIRPORT 2 ID}</div>
            <div class="elevation">{AIRPORT 2 ELEVATION}</div>
{AIRPORT 2 RUNWAY ENDS}
        </div>
    </div>
    <div class="notes">
        <div class="note-row"></div>
        <div class="note-row"></div>
        <div class="note-row"></div>
        <div class="note-row"></div>
        <div class="note-row"></div>
    </div>
</div>
</body>
</html>
`
