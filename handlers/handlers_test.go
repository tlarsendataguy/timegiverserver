package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"testing"
	"time"
	"timegiverserver/lang"
)

type noWriter struct {
}

func (n *noWriter) Write(value []byte) (int, error) {
	return len(value), nil
}

func TestLoadHandler(t *testing.T) {
	server, err := LoadServerFromSettings(`settings.json`, &noWriter{}, `TEST`)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if value := server.CertFolder; value != `./certs` {
		t.Fatalf(`expected './certs' but got '%v'`, value)
	}
	if value := server.ServeFolder; value != `./serveTest` {
		t.Fatalf(`expected './serveTest' but got '%v'`, value)
	}
}

func TestLoadHomepage(t *testing.T) {
	server, _ := LoadServerFromSettings(`settings.json`, &noWriter{}, `TEST`)
	w := &testWriter{}
	server.HandleHomepage(w, nil)

	err := checkResponse(w, 200, `./serveTest/index.html`)
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestLoadFile(t *testing.T) {
	server, _ := LoadServerFromSettings(`settings.json`, &noWriter{}, `TEST`)
	w := &testWriter{}
	r := getRequestFor(`https://www.timegiver.app/scripts.js`)
	server.HandleFile(w, r)

	err := checkResponse(w, 200, `./serveTest/scripts.js`)
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func Test404Response(t *testing.T) {
	server, _ := LoadServerFromSettings(`settings.json`, &noWriter{}, `TEST`)
	w := &testWriter{}
	r := getRequestFor(`https://www.timegiver.app/invalid_file`)
	server.HandleFile(w, r)

	err := checkResponse(w, 404, `./serveTest/404.html`)
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestCalculateEmail(t *testing.T) {
	t.Skip(`Skipped by default. This test will send an e-mail and requires a non-tracked server file be created with the necessary SMTP authorization fields`)
	server, _ := LoadServerFromSettings(`smtp_auth.json`, &noWriter{}, `TEST`)
	w := &testWriter{}
	r := getRequestFor(`https://www.timegiver.app/api/calculate`)
	r.Method = `POST`
	body, _ := json.Marshal(tempPayload{
		DepartureOffset: -4,
		ArrivalOffset:   2,
		Email:           "larsenthomasj@gmail.com",
		Arrival:         `2020-03-04T08:30`,
		Wake:            `06:00`,
		Breakfast:       `07:00`,
		Lunch:           `12:00`,
		Dinner:          `17:00`,
		Sleep:           `22:00`,
	})
	r.Body = io.NopCloser(bytes.NewReader(body))

	server.HandleCalculateApi(w, r)
	if w.status != 200 {
		t.Log(string(w.content))
		t.Fatalf(`expected 200 but got %v`, w.status)
	}
}

func TestDb(t *testing.T) {
	t.Skip(`Skipped by default. This test will insert a record into snowflake and requires a non-tracked server file be created with the necessary connection string`)
	server, err := LoadServerFromSettings(`conn_str.json`, &noWriter{}, `TEST`)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	params := CalcPayload{
		DepartureOffset: -4,
		ArrivalOffset:   2,
		DepartureLoc:    `Raleigh, NC`,
		ArrivalLoc:      `Krakow, PL`,
		Email:           "larsenthomasj@gmail.com",
		Arrival:         time.Date(2022, 3, 4, 8, 30, 0, 0, time.UTC),
		Wake:            6 * time.Hour,
		Breakfast:       7 * time.Hour,
		Lunch:           12 * time.Hour,
		Dinner:          17 * time.Hour,
		Sleep:           22 * time.Hour,
	}
	err = server.insertApiRequest(params, lang.EN, nil)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
}

func TestTimezoneApi(t *testing.T) {
	t.Skip(`Skipped by default. This test calls the Google Maps Timezone API`)
	server, err := LoadServerFromSettings(`maps_api.json`, &noWriter{}, `TEST`)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	w := &testWriter{}
	r := getRequestFor(`https://www.timegiver.app/api/timezones`)
	r.Method = `POST`
	body, _ := json.Marshal(TimezoneRequestPayload{
		Timestamp: `2022-01-02T03:04`,
		From: Coordinates{
			Lat: 35.77664880968805,
			Lng: -78.64098235711558,
		},
		To: Coordinates{
			Lat: 51.51045976704988,
			Lng: -0.12275095972896906,
		},
	})
	r.Body = io.NopCloser(bytes.NewReader(body))

	server.HandleTimezoneApi(w, r)
	t.Log(string(w.content))
	if w.status != 200 {
		t.Fatalf(`expected 200 but got %v`, w.status)
	}
}

func checkResponse(w *testWriter, expectedStatus int, expectedFile string) error {
	if w.status != expectedStatus {
		return fmt.Errorf(`expected status %v but got %v`, expectedStatus, w.status)
	}
	expected, _ := os.ReadFile(expectedFile)
	if !reflect.DeepEqual(w.content, expected) {
		return fmt.Errorf("expected %v content but got:\n%v", expectedFile, w.content)
	}
	return nil
}

func getRequestFor(testUrl string) *http.Request {
	u, _ := url.Parse(testUrl)
	return &http.Request{
		Method: "GET",
		URL:    u,
	}
}

type testWriter struct {
	content []byte
	status  int
	header  http.Header
}

func (w *testWriter) Header() http.Header {
	if w.header == nil {
		w.header = make(http.Header)
	}
	return w.header
}

func (w *testWriter) Write(content []byte) (int, error) {
	if w.status == 0 {
		w.status = 200
	}
	if w.content == nil {
		w.content = make([]byte, len(content))
		copy(w.content, content)
		return len(content), nil
	}
	w.content = append(w.content, content...)
	return len(content), nil
}

func (w *testWriter) WriteHeader(status int) {
	w.status = status
}
