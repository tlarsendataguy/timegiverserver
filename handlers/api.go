package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session"
	"github.com/stripe/stripe-go/v76/webhook"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
	"timegiverserver/handlers/kneeboard"
)

func (s *Server) handleSessionWebhook(w http.ResponseWriter, r *http.Request) {
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		writeError(w, fmt.Errorf(`error reading request body: %v`, err.Error()))
		return
	}

	event, err := webhook.ConstructEvent(payload, r.Header.Get(`Stripe-Signature`), s.WebhookSigningSecret)
	if err != nil {
		writeError(w, fmt.Errorf(`error constructing event: %v`, err.Error()))
		return
	}

	if event.Type != `checkout.session.completed` {
		writeError(w, fmt.Errorf(`expecting 'checkout.session.completed' event but got '%v'`, event.Type))
		return
	}

	var sn stripe.CheckoutSession
	err = json.Unmarshal(event.Data.Raw, &sn)
	if err != nil {
		writeError(w, fmt.Errorf(`error parsing checkout session: %v`, err.Error()))
		return
	}

	if sn.Status != stripe.CheckoutSessionStatusComplete {
		writeError(w, errors.New(`checkout has not been completed`))
		return
	}

	metadata, err := MetadataFromMap(sn.Metadata)
	if err != nil {
		writeError(w, fmt.Errorf(`error parsing metadata: %v`, err.Error()))
		return
	}
	params, err := ValidateMetadataPayload(metadata)
	if err != nil {
		writeError(w, err)
		return
	}
	ics := s.generateIcs(params)

	err = s.emailPlan(ics, sn.CustomerDetails.Email)
	if err != nil {
		writeError(w, err)
		return
	}

	w.WriteHeader(200)
}

func (s *Server) handleCheckout(w http.ResponseWriter, r *http.Request) {
	metadata := MetadataPayload{}
	err := json.NewDecoder(r.Body).Decode(&metadata)
	if err != nil {
		writeError(w, fmt.Errorf(`error decoding payload: %v`, err.Error()))
		return
	}

	payload, err := ValidateMetadataPayload(metadata)
	if err != nil {
		writeError(w, fmt.Errorf(`error validating payload: %v`, err.Error()))
		return
	}

	productName := fmt.Sprintf(`Beat jet lag on a trip from %v to %v, arriving %v`, metadata.DepartureLoc, metadata.ArrivalLoc, payload.Arrival.Format(`January 2 at 3:04pm`))
	params := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String(`USD`),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Images: nil,
						Name:   stripe.String(productName),
					},
					UnitAmount: stripe.Int64(150),
				},
				Quantity: stripe.Int64(1),
			},
		},
		Metadata:   metadata.ToMap(),
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String("https://" + r.Host + `/success.html`),
	}

	se, err := session.New(params)
	if err != nil {
		writeError(w, fmt.Errorf(`error creating new checkout session: %v`, err.Error()))
		return
	}

	_, _ = w.Write([]byte(se.URL))
}

func (s *Server) handleTimezoneApi(w http.ResponseWriter, r *http.Request) {
	var payload TimezoneRequestPayload
	d := json.NewDecoder(r.Body)
	err := d.Decode(&payload)
	if err != nil {
		writeErrorMsg(w, `JSON payload is invalid`)
		return
	}
	timestamp, err := time.Parse(`2006-01-02T15:04`, payload.Timestamp)
	if err != nil {
		writeErrorMsg(w, `Timestamp is not formatted correctly`)
		return
	}

	fromOffset, err := s.requestTimezone(payload.From, timestamp)
	if err != nil {
		writeErrorMsg(w, fmt.Sprintf(`error obtaining departure timezone: %v`, err.Error()))
		return
	}
	toOffset, err := s.requestTimezone(payload.To, timestamp)
	if err != nil {
		writeErrorMsg(w, fmt.Sprintf(`error obtaining arrival timezone: %v`, err.Error()))
		return
	}

	response := TimezoneResponsePayload{
		FromOffset: fromOffset.Offset(),
		ToOffset:   toOffset.Offset(),
	}
	responseBytes, err := json.Marshal(response)
	if err != nil {
		writeErrorMsg(w, `the server did not marshall the JSON correctly`)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	_, _ = w.Write(responseBytes)
}

func (s *Server) handleKneeboardApi(w http.ResponseWriter, r *http.Request) {
	rootPath := filepath.Join(s.ServeFolder, `kneeboard`)
	variables := r.URL.Query()
	from := strings.ToUpper(variables.Get(`from`))
	to := strings.ToUpper(variables.Get(`to`))
	kb := &kneeboard.Kneeboard{From: from, To: to}

	result, err := s.Db.Query(`SELECT * FROM NASR.PUBLIC.FREQUENCIES WHERE FACILITY IN (?, ?)`, from, to)
	if err != nil {
		writeError(w, err)
		return
	}
	err = kb.LoadFrequencies(result)
	if err != nil {
		writeError(w, err)
		return
	}

	result, err = s.Db.Query(`SELECT * FROM NASR.PUBLIC.RUNWAY_BOXES WHERE FACILITY in (?, ?)`, from, to)
	if err != nil {
		writeError(w, err)
		return
	}
	err = kb.LoadRunways(result)
	if err != nil {
		writeError(w, err)
		return
	}

	result, err = s.Db.Query(`SELECT "Location Identifier", "Official Facility Name", "Airport Elevation" FROM NASR.PUBLIC.APT WHERE "Location Identifier" IN (?, ?)`, from, to)
	if err != nil {
		writeError(w, err)
		return
	}
	err = kb.LoadInfo(result)
	if err != nil {
		writeError(w, err)
		return
	}

	var f *os.File
	if from != `` {
		f, err = os.OpenFile(filepath.Join(rootPath, from+`.png`), os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
		if err != nil {
			writeError(w, err)
			return
		}
		err = kb.CreateRwyImage(from, f)
		if err != nil {
			writeError(w, err)
			return
		}
		_ = f.Close()
	}

	if to != `` {
		f, err = os.OpenFile(filepath.Join(rootPath, to+`.png`), os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
		if err != nil {
			writeError(w, err)
			return
		}
		err = kb.CreateRwyImage(to, f)
		if err != nil {
			writeError(w, err)
			return
		}
		_ = f.Close()
	}

	_, _ = w.Write(kb.BuildHtml())
}
