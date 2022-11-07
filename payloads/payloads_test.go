package payloads

import (
	"bytes"
	"errors"
	"fmt"
	"testing"
)

func TestCalculatorPayload(t *testing.T) {
	body := `{"DepartureOffset":0,"ArrivalOffset":-4.5,"Email":"me@me.com"}`
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
}

func TestInvalidJson(t *testing.T) {
	body := `hello world`
	_, err := checkPayload(body)
	if test := checkError(err, `invalid JSON`); test != nil {
		t.Fatalf(test.Error())
	}
}

func TestEmailNotProvided(t *testing.T) {
	body := `{"DepartureOffset":0,"ArrivalOffset":-4.5}`
	_, err := checkPayload(body)
	if test := checkError(err, `a valid e-mail address was not provided`); test != nil {
		t.Fatalf(test.Error())
	}
}

func TestDepartureAndArrivalBothZero(t *testing.T) {
	body := `{"Email":"me@me.com"}`
	_, err := checkPayload(body)
	if test := checkError(err, `no plan is needed when departure offset is equivalent to arrival offset`); test != nil {
		t.Fatalf(test.Error())
	}
}

func TestInvalidEmail(t *testing.T) {
	body := `{"DepartureOffset":0,"ArrivalOffset":-4.5,"Email":"hello world"}`
	_, err := checkPayload(body)
	if test := checkError(err, `a valid e-mail address was not provided`); test != nil {
		t.Fatalf(test.Error())
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
