package kneeboard

import (
	"fmt"
	"os"
	"testing"
)

func TestFromFrequencyHtml(t *testing.T) {
	kb := testKneeboard()

	freqStr := kb.fromFrequenciesHtml()
	expected := `            <div class="frequency">
                <div class="frequency-name">ATIS</div>
                <div class="frequency-code">123.800</div>
                <div class="frequency-comment"></div>
            </div>
            <div class="frequency">
                <div class="frequency-name">Ground</div>
                <div class="frequency-code">121.900</div>
                <div class="frequency-comment">EAST, RWY 05R/23L, 14/32</div>
            </div>
`
	if freqStr != expected {
		t.Fatalf("expected\n\n%v\n\nbut got\n\n%v", expected, freqStr)
	}
}

func TestToFrequencyHtml(t *testing.T) {
	kb := testKneeboard()

	freqStr := kb.toFrequenciesHtml()
	expected := `            <div class="frequency">
                <div class="frequency-name">Unicom</div>
                <div class="frequency-code">123.075</div>
                <div class="frequency-comment"></div>
            </div>
`
	if freqStr != expected {
		t.Fatalf("expected\n\n%v\n\nbut got\n\n%v", expected, freqStr)
	}
}

func TestFromRwyEndsHtml(t *testing.T) {
	kb := testKneeboard()

	rwyStr := kb.fromRwyEndsHtml()
	expected := fmt.Sprintf(`            <div class="runway-id" style="left: %.fpx;top: %.fpx">05L</div>
            <div class="runway-id" style="left: %.fpx;top: %.fpx">23R</div>
            <div class="runway-id" style="left: %.fpx;top: %dpx">05R</div>
            <div class="runway-id" style="left: %.fpx;top: %.fpx">23L</div>
`, (rwyImageSize*0.024)+rwyImagePadding, (rwyImageSize*0.6636)+rwyImagePadding, (rwyImageSize*0.8439)+rwyImagePadding, 0.0+rwyImagePadding, (rwyImageSize*0.1826)+rwyImagePadding, rwyImageSize+rwyImagePadding, (rwyImageSize*0.79727)+rwyImagePadding, (rwyImageSize*0.5022)+rwyImagePadding)
	if rwyStr != expected {
		t.Fatalf("expected\n\n%v\n\nbut got\n\n%v", expected, rwyStr)
	}
}

func TestToRwyEndsHtml(t *testing.T) {
	kb := testKneeboard()

	rwyStr := kb.toRwyEndsHtml()
	expected := fmt.Sprintf(`            <div class="runway-id" style="left: %.fpx;top: %.dpx">03</div>
            <div class="runway-id" style="left: %.fpx;top: %.fpx">21</div>
`, (rwyImageSize*0.26497)+rwyImagePadding, rwyImageSize+rwyImagePadding, (rwyImageSize*0.73502)+rwyImagePadding, 0.0+rwyImagePadding)
	if rwyStr != expected {
		t.Fatalf("expected\n\n%v\n\nbut got\n\n%v", expected, rwyStr)
	}
}

func TestFullPageHtml(t *testing.T) {
	kb := testKneeboard()

	content := kb.BuildHtml()

	err := os.WriteFile(`kneeboard.html`, content, os.ModePerm)
	if err != nil {
		t.Fatalf(`got error %v`, err.Error())
	}
}

func TestImage(t *testing.T) {
	f, err := os.OpenFile(`RDU.png`, os.O_CREATE, os.ModePerm)
	if err != nil {
		t.Fatalf(`got error %v`, err.Error())
	}

	kb := testKneeboard()

	err = kb.CreateRwyImage(`RDU`, f)
	if err != nil {
		t.Fatalf(`got error %v`, err.Error())
	}
}

func testKneeboard() *Kneeboard {
	return &Kneeboard{
		From: `RDU`,
		To:   `TTA`,
		Info: map[string]Info{
			"RDU": {`Raleigh International`, 435.2},
			"TTA": {`Raleigh Executive`, 246.6},
		},
		Frequencies: []Frequency{
			{
				Facility:  "RDU",
				Label:     "ATIS",
				Frequency: "123.800",
				Sort:      0,
				Comment:   "",
			},
			{
				Facility:  "RDU",
				Label:     "Ground",
				Frequency: "121.900",
				Sort:      3,
				Comment:   "EAST, RWY 05R/23L, 14/32",
			},
			{
				Facility:  "TTA",
				Label:     "Unicom",
				Frequency: "123.075",
				Sort:      1,
				Comment:   "",
			},
		},
		Runways: []Runway{
			{
				Facility:   "RDU",
				Rwy:        "05L/23R",
				Identifier: "05L",
				X:          0.02415932204550982,
				Y:          0.6636671610655724,
			},
			{
				Facility:   "RDU",
				Rwy:        "05L/23R",
				Identifier: "23R",
				X:          0.8438669987017747,
				Y:          -0,
			},
			{
				Facility:   "RDU",
				Rwy:        "05R/23L",
				Identifier: "05R",
				X:          0.18259761523890428,
				Y:          1,
			},
			{
				Facility:   "RDU",
				Rwy:        "05R/23L",
				Identifier: "23L",
				X:          0.7972713290881251,
				Y:          0.5022247950756499,
			},
			{
				Facility:   "TTA",
				Rwy:        "03/21",
				Identifier: "03",
				X:          0.2649715684157489,
				Y:          1,
			},
			{
				Facility:   "TTA",
				Rwy:        "03/21",
				Identifier: "21",
				X:          0.735028431584251,
				Y:          -0,
			},
		},
	}
}
