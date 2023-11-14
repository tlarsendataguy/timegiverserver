package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestCalculatorPayload(t *testing.T) {
	body := `{"DepartureOffset":0,"ArrivalOffset":-4.5,"Arrival":"2022-01-02T03:04","Wake":"06:00","Breakfast":"07:00","Lunch":"12:00","Dinner":"17:00","Sleep":"22:00"}`
	payload, err := checkPayload(body)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if payload.DepartureOffset != 0 {
		t.Fatalf(`expected 0 but got %v`, payload.DepartureOffset)
	}
	if payload.ArrivalOffset != -4.5 {
		t.Fatalf(`expected -4.5 but got %v`, payload.ArrivalOffset)
	}
	if !payload.Arrival.Equal(time.Date(2022, 1, 2, 3, 4, 0, 0, time.UTC)) {
		t.Fatalf(`expected '2022-01-02 03:04:00' but got '%v'`, payload.Arrival)
	}
}

func TestEmailNotProvided(t *testing.T) {
	body := `{"DepartureOffset":0,"ArrivalOffset":-4.5,"Arrival":"2022-01-02T03:04","Wake":"06:00","Breakfast":"07:00","Lunch":"12:00","Dinner":"17:00","Sleep":"22:00"}`
	_, err := checkPayload(body)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
}

func TestDepartureAndArrivalBothZero(t *testing.T) {
	body := `{"Email":"me@me.com","Arrival":"2022-01-02T03:04","Wake":"06:00","Breakfast":"07:00","Lunch":"12:00","Dinner":"17:00","Sleep":"22:00"}`
	_, err := checkPayload(body)
	if test := checkError(err, `no plan is needed when departure and arrival are in the same time zone`); test != nil {
		t.Fatalf(test.Error())
	}
}

func TestArrivalNotProvided(t *testing.T) {
	body := `{"DepartureOffset":0,"ArrivalOffset":-4.5,"Wake":"06:00","Breakfast":"07:00","Lunch":"12:00","Dinner":"17:00","Sleep":"22:00"}`
	_, err := checkPayload(body)
	if test := checkError(err, `invalid date provided for Arrival`); test != nil {
		t.Fatalf(test.Error())
	}
}

func TestInvalidArrival(t *testing.T) {
	body := `{"DepartureOffset":0,"ArrivalOffset":-4.5,"Arrival":"20220102T030400Z","Wake":"06:00","Breakfast":"07:00","Lunch":"12:00","Dinner":"17:00","Sleep":"22:00"}`
	_, err := checkPayload(body)
	if test := checkError(err, `invalid date provided for Arrival`); test != nil {
		t.Fatalf(test.Error())
	}
}

func TestParseTime(t *testing.T) {
	d, err := parseTime(`15:30`)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if d.Hours() != 15.5 {
		t.Fatalf(`expected 15.5 hours but got %v`, d.Hours())
	}
	_, err = parseTime(`30:30`)
	if err == nil {
		t.Fatalf(`expected an error but got none`)
	}
	_, err = parseTime(`-2:30`)
	if err == nil {
		t.Fatalf(`expected an error but got none`)
	}
	_, err = parseTime(``)
	if err == nil {
		t.Fatalf(`expected an error but got none`)
	}
}

func TestParseMetadata(t *testing.T) {
	source := map[string]string{
		"Wake":            "06:00",
		"ArrivalOffset":   "1",
		"Arrival":         "2023-12-30T20:07",
		"ArrivalLoc":      "Zürich, Switzerland",
		"Lunch":           "12:00",
		"Breakfast":       "07:00",
		"Dinner":          "18:00",
		"DepartureLoc":    "Raleigh, NC, USA",
		"DepartureOffset": "-5",
		"Sleep":           "22:00",
	}

	metadata, err := MetadataFromMap(source)
	if err != nil {
		t.Fatalf(`got error %v`, err.Error())
	}
	if metadata.ArrivalOffset != 1 {
		t.Fatalf(`expected 1 but got %v`, metadata.ArrivalOffset)
	}
}

func checkPayload(body string) (CalcPayload, error) {
	payload := MetadataPayload{}
	err := json.Unmarshal([]byte(body), &payload)
	if err != nil {
		return CalcPayload{}, err
	}
	return ValidateMetadataPayload(payload)
}

func checkError(err error, msg string) error {
	if err == nil {
		return errors.New(`expected a error but got none`)
	}
	if txt := err.Error(); txt != msg {
		return fmt.Errorf(`expected '%v' but got '%v'`, msg, txt)
	}
	return nil
}
