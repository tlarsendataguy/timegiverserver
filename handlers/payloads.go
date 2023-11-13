package handlers

import (
	"errors"
	"math"
	"net/mail"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type CalcPayload struct {
	DepartureOffset float64
	ArrivalOffset   float64
	Email           string
	Arrival         time.Time
	Wake            time.Duration
	Breakfast       time.Duration
	Lunch           time.Duration
	Dinner          time.Duration
	Sleep           time.Duration
}

type metadataPayload struct {
	DepartureOffset float64
	ArrivalOffset   float64
	DepartureLoc    string
	ArrivalLoc      string
	Email           string
	Arrival         string
	Wake            string
	Breakfast       string
	Lunch           string
	Dinner          string
	Sleep           string
}

func (m metadataPayload) ToMap() map[string]string {
	result := map[string]string{}
	result[`DepartureOffset`] = strconv.FormatFloat(m.DepartureOffset, 'f', -1, 64)
	result[`ArrivalOffset`] = strconv.FormatFloat(m.ArrivalOffset, 'f', -1, 64)
	result[`DepartureLoc`] = m.DepartureLoc
	result[`ArrivalLoc`] = m.ArrivalLoc
	result[`Email`] = m.Email
	result[`Arrival`] = m.Email
	result[`Wake`] = m.Wake
	result[`Breakfast`] = m.Breakfast
	result[`Lunch`] = m.Lunch
	result[`Dinner`] = m.Dinner
	result[`Sleep`] = m.Sleep
	return result
}

func (m metadataPayload) FromMap(source map[string]string) error {
	var err error
	m.DepartureOffset, err = strconv.ParseFloat(source[`DepartureOffset`], 64)
	if err != nil {
		return err
	}
	m.ArrivalOffset, err = strconv.ParseFloat(source[`ArrivalOffset`], 64)
	if err != nil {
		return err
	}
	m.DepartureLoc = source[`DepartureLoc`]
	m.ArrivalLoc = source[`ArrivalLoc`]
	m.Arrival = source[`Arrival`]
	m.Wake = source[`Wake`]
	m.Breakfast = source[`Breakfast`]
	m.Lunch = source[`Lunch`]
	m.Dinner = source[`Dinner`]
	m.Sleep = source[`Sleep`]
	return nil
}

func ValidateMetadataPayload(payload metadataPayload) (CalcPayload, error) {
	email := payload.Email
	if email != `` {
		if _, emailErr := mail.ParseAddress(email); emailErr != nil {
			return CalcPayload{}, errors.New(`a valid e-mail address was not provided`)
		}
	}

	if math.Trunc(payload.DepartureOffset) == math.Trunc(payload.ArrivalOffset) {
		return CalcPayload{}, errors.New(`no plan is needed when departure and arrival are in the same time zone`)
	}

	output := CalcPayload{
		DepartureOffset: payload.DepartureOffset,
		ArrivalOffset:   payload.ArrivalOffset,
		Email:           email,
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
