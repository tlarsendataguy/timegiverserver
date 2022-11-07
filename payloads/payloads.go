package payloads

import (
	"encoding/json"
	"errors"
	"io"
	"math"
	"net/mail"
)

type CalcPayload struct {
	DepartureOffset float64
	ArrivalOffset   float64
	Email           string
}

func ValidateCalcPayload(r io.Reader) (CalcPayload, error) {
	d := json.NewDecoder(r)
	payload := CalcPayload{}
	err := d.Decode(&payload)
	if err != nil {
		return CalcPayload{}, errors.New(`invalid JSON`)
	}
	if _, emailErr := mail.ParseAddress(payload.Email); emailErr != nil {
		return CalcPayload{}, errors.New(`a valid e-mail address was not provided`)
	}
	if math.Trunc(payload.DepartureOffset) == math.Trunc(payload.ArrivalOffset) {
		return CalcPayload{}, errors.New(`no plan is needed when departure offset is equivalent to arrival offset`)
	}
	return payload, err
}
