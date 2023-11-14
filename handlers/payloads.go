package handlers

import (
	"errors"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type CheckoutUrlPayload struct {
	CheckoutUrl string
}

type CalcPayload struct {
	DepartureOffset float64
	ArrivalOffset   float64
	Arrival         time.Time
	Wake            time.Duration
	Breakfast       time.Duration
	Lunch           time.Duration
	Dinner          time.Duration
	Sleep           time.Duration
}

type MetadataPayload struct {
	DepartureOffset float64
	ArrivalOffset   float64
	DepartureLoc    string
	ArrivalLoc      string
	Arrival         string
	Wake            string
	Breakfast       string
	Lunch           string
	Dinner          string
	Sleep           string
}

func (m MetadataPayload) ToMap() map[string]string {
	result := map[string]string{}
	result[`DepartureOffset`] = strconv.FormatFloat(m.DepartureOffset, 'f', -1, 64)
	result[`ArrivalOffset`] = strconv.FormatFloat(m.ArrivalOffset, 'f', -1, 64)
	result[`DepartureLoc`] = m.DepartureLoc
	result[`ArrivalLoc`] = m.ArrivalLoc
	result[`Arrival`] = m.Arrival
	result[`Wake`] = m.Wake
	result[`Breakfast`] = m.Breakfast
	result[`Lunch`] = m.Lunch
	result[`Dinner`] = m.Dinner
	result[`Sleep`] = m.Sleep
	return result
}

func MetadataFromMap(source map[string]string) (MetadataPayload, error) {
	var m MetadataPayload
	var err error
	m.DepartureOffset, err = strconv.ParseFloat(source[`DepartureOffset`], 64)
	if err != nil {
		return m, err
	}
	m.ArrivalOffset, err = strconv.ParseFloat(source[`ArrivalOffset`], 64)
	if err != nil {
		return m, err
	}
	m.DepartureLoc = source[`DepartureLoc`]
	m.ArrivalLoc = source[`ArrivalLoc`]
	m.Arrival = source[`Arrival`]
	m.Wake = source[`Wake`]
	m.Breakfast = source[`Breakfast`]
	m.Lunch = source[`Lunch`]
	m.Dinner = source[`Dinner`]
	m.Sleep = source[`Sleep`]
	return m, nil
}

func ValidateMetadataPayload(payload MetadataPayload) (CalcPayload, error) {
	if math.Trunc(payload.DepartureOffset) == math.Trunc(payload.ArrivalOffset) {
		return CalcPayload{}, errors.New(`no plan is needed when departure and arrival are in the same time zone`)
	}

	output := CalcPayload{
		DepartureOffset: payload.DepartureOffset,
		ArrivalOffset:   payload.ArrivalOffset,
	}

	var err error
	output.Arrival, err = time.Parse(`2006-01-02T15:04`, payload.Arrival)
	if err != nil {
		return CalcPayload{}, errors.New(`invalid date provided for Arrival`)
	}
	output.Wake, err = parseTime(payload.Wake)
	if err != nil {
		return CalcPayload{}, errors.New(`error parsing Wake: ` + err.Error())
	}
	output.Breakfast, err = parseTime(payload.Breakfast)
	if err != nil {
		return CalcPayload{}, errors.New(`error parsing Breakfast: ` + err.Error())
	}
	output.Lunch, err = parseTime(payload.Lunch)
	if err != nil {
		return CalcPayload{}, errors.New(`error parsing Lunch: ` + err.Error())
	}
	output.Dinner, err = parseTime(payload.Dinner)
	if err != nil {
		return CalcPayload{}, errors.New(`error parsing Dinner: ` + err.Error())
	}
	output.Sleep, err = parseTime(payload.Sleep)
	if err != nil {
		return CalcPayload{}, errors.New(`error parsing Sleep: ` + err.Error())
	}

	return output, nil
}

func parseTime(value string) (time.Duration, error) {
	if valid, _ := regexp.MatchString(`^[0-9]{2}:[0-9]{2}$`, value); !valid {
		return 0 * time.Minute, errors.New(`time is not formatted correctly`)
	}
	pieces := strings.Split(value, `:`)
	builder := strings.Builder{}
	builder.WriteString(pieces[0])
	builder.WriteString(`h`)
	builder.WriteString(pieces[1])
	builder.WriteString(`m`)
	d, err := time.ParseDuration(builder.String())
	if err != nil {
		return 0 * time.Minute, errors.New(`time is not formatted correctly`)
	}
	if hrs := d.Hours(); hrs < 0 || hrs >= 24 {
		return 0 * time.Minute, errors.New(`time is not between 00:00 and 23:59`)
	}
	return d, nil
}

type Coordinates struct {
	Lat float64
	Lng float64
}

type TimezoneRequestPayload struct {
	Timestamp string
	From      Coordinates
	To        Coordinates
}

type TimezoneResponsePayload struct {
	FromOffset float64
	ToOffset   float64
}

type googleTimezoneResponse struct {
	DstOffset float64 `json:"dstOffset"`
	RawOffset float64 `json:"rawOffset"`
}

func (g *googleTimezoneResponse) Offset() float64 {
	return (g.DstOffset + g.RawOffset) / 3600
}
