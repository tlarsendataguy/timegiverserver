package handlers

import (
	"bytes"
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestCalculatorPayload(t *testing.T) {
	body := `{"DepartureOffset":0,"ArrivalOffset":-4.5,"Email":"me@me.com","Arrival":"20220102T030400","Wake":"06:00","Breakfast":"07:00","Lunch":"12:00","Dinner":"17:00","Sleep":"22:00"}`
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
	if payload.Email != `me@me.com` {
		t.Fatalf(`expected 'me@me.com' but got '%v'`, payload.Email)
	}
	if !payload.Arrival.Equal(time.Date(2022, 1, 2, 3, 4, 0, 0, time.UTC)) {
		t.Fatalf(`expected '2022-01-02 03:04:00' but got '%v'`, payload.Arrival)
	}
}

func TestInvalidJson(t *testing.T) {
	body := `hello world`
	_, err := checkPayload(body)
	if test := checkError(err, `invalid JSON`); test != nil {
		t.Fatalf(test.Error())
	}
}

func TestEmailNotProvided(t *testing.T) {
	body := `{"DepartureOffset":0,"ArrivalOffset":-4.5,"Arrival":"20220102T030400","Wake":"06:00","Breakfast":"07:00","Lunch":"12:00","Dinner":"17:00","Sleep":"22:00"}`
	_, err := checkPayload(body)
	if test := checkError(err, `a valid e-mail address was not provided`); test != nil {
		t.Fatalf(test.Error())
	}
}

func TestDepartureAndArrivalBothZero(t *testing.T) {
	body := `{"Email":"me@me.com","Arrival":"20220102T030400","Wake":"06:00","Breakfast":"07:00","Lunch":"12:00","Dinner":"17:00","Sleep":"22:00"}`
	_, err := checkPayload(body)
	if test := checkError(err, `no plan is needed when departure offset is equivalent to arrival offset`); test != nil {
		t.Fatalf(test.Error())
	}
}

func TestInvalidEmail(t *testing.T) {
	body := `{"DepartureOffset":0,"ArrivalOffset":-4.5,"Email":"hello world","Arrival":"20220102T030400","Wake":"06:00","Breakfast":"07:00","Lunch":"12:00","Dinner":"17:00","Sleep":"22:00"}`
	_, err := checkPayload(body)
	if test := checkError(err, `a valid e-mail address was not provided`); test != nil {
		t.Fatalf(test.Error())
	}
}

func TestArrivalNotProvided(t *testing.T) {
	body := `{"DepartureOffset":0,"ArrivalOffset":-4.5,"Email":"me@me.com","Wake":"06:00","Breakfast":"07:00","Lunch":"12:00","Dinner":"17:00","Sleep":"22:00"}`
	_, err := checkPayload(body)
	if test := checkError(err, `invalid date provided for Arrival`); test != nil {
		t.Fatalf(test.Error())
	}
}

func TestInvalidArrival(t *testing.T) {
	body := `{"DepartureOffset":0,"ArrivalOffset":-4.5,"Email":"me@me.com","Arrival":"2022-01-02T03:04:00Z","Wake":"06:00","Breakfast":"07:00","Lunch":"12:00","Dinner":"17:00","Sleep":"22:00"}`
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

func checkPayload(body string) (CalcPayload, error) {
	reader := bytes.NewReader([]byte(body))
	return ValidateCalcPayload(reader)
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
