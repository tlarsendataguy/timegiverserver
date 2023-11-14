package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"testing"
)

func TestLoadHandler(t *testing.T) {
	server, err := LoadServerFromSettings(`settings.json`)
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
	server, _ := LoadServerFromSettings(`settings.json`)
	router := server.GenerateRouter()
	w := &testWriter{}
	r := getRequestFor(`https://www.host1.com/`)
	router.ServeHTTP(w, r)

	err := checkResponse(w, 200, `./serveTest/host1/index.html`)
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestLoadFile(t *testing.T) {
	server, _ := LoadServerFromSettings(`settings.json`)
	router := server.GenerateRouter()
	w := &testWriter{}
	r := getRequestFor(`https://www.host1.com/scripts.js`)
	router.ServeHTTP(w, r)

	err := checkResponse(w, 200, `./serveTest/host1/scripts.js`)
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func Test404Response(t *testing.T) {
	server, _ := LoadServerFromSettings(`settings.json`)
	router := server.GenerateRouter()
	w := &testWriter{}
	r := getRequestFor(`https://www.host1.com/invalid_file`)
	router.ServeHTTP(w, r)

	err := checkResponse(w, 404, `./serveTest/host1/404.html`)
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestCalculateEmail(t *testing.T) {
	t.Skip(`Skipped by default. This test will send an e-mail and requires a non-tracked server file be created with the necessary SMTP authorization fields`)
	server, _ := LoadServerFromSettings(`settings_test.json`)
	router := server.GenerateRouter()
	w := &testWriter{}
	r := getRequestFor(`https://www.host1.com/api/calculate`)
	r.Method = `POST`
	body, _ := json.Marshal(MetadataPayload{
		DepartureOffset: -4,
		ArrivalOffset:   2,
		Arrival:         `2020-03-04T08:30`,
		Wake:            `06:00`,
		Breakfast:       `07:00`,
		Lunch:           `12:00`,
		Dinner:          `17:00`,
		Sleep:           `22:00`,
	})
	r.Body = io.NopCloser(bytes.NewReader(body))

	router.ServeHTTP(w, r)
	if w.status != 200 {
		t.Log(string(w.content))
		t.Fatalf(`expected 200 but got %v`, w.status)
	}
}

func TestTimezoneApi(t *testing.T) {
	t.Skip(`Skipped by default. This test calls the Google Maps Timezone API`)
	server, err := LoadServerFromSettings(`settings_test.json`)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	router := server.GenerateRouter()
	w := &testWriter{}
	r := getRequestFor(`https://www.host1.com/api/timezones`)
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

	router.ServeHTTP(w, r)
	t.Log(string(w.content))
	if w.status != 200 {
		t.Fatalf(`expected 200 but got %v`, w.status)
	}
}

func TestHostWhitelist(t *testing.T) {
	server, err := LoadServerFromSettings(`settings.json`)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	whitelist := server.CollectHostWhitelist()
	expected := []string{"host1.com", "www.host1.com", "something.somewhere.com"}
	if !reflect.DeepEqual(whitelist, expected) {
		t.Fatalf(`expected %v but got %v`, expected, whitelist)
	}
}

func TestRouter(t *testing.T) {
	server, err := LoadServerFromSettings(`settings.json`)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	router := server.GenerateRouter()
	if err = checkRoute(router, `https://www.host1.com/`); err != nil {
		t.Fatalf(err.Error())
	}
	if err = checkRoute(router, `https://www.host1.com/index.html`); err != nil {
		t.Fatalf(err.Error())
	}
	if err = checkApiRoute(router, `https://www.host1.com/api/timezones`); err != nil {
		t.Fatalf(err.Error())
	}
	if err = checkRoute(router, `https://host1.com/`); err != nil {
		t.Fatalf(err.Error())
	}
	if err = checkApiRoute(router, `https://host1.com/api/timezones`); err != nil {
		t.Fatalf(err.Error())
	}
	if err = checkRoute(router, `https://host1.com/index.html`); err != nil {
		t.Fatalf(err.Error())
	}
	if err = checkRoute(router, `https://something.somewhere.com/`); err != nil {
		t.Fatalf(err.Error())
	}
	if err = checkRoute(router, `https://something.somewhere.com/index.html`); err != nil {
		t.Fatalf(err.Error())
	}

	w := &testWriter{}
	r := getRequestFor(`https://something.somewhere.com/api/timezones`)
	r.Method = `POST`
	router.ServeHTTP(w, r)
	if w.status != 405 {
		t.Fatalf(`expected status 405 but got %v`, w.status)
	}

	r = getRequestFor(`https://something.somewhere.com/api/calculate`)
	r.Method = `POST`
	router.ServeHTTP(w, r)
	if w.status != 405 {
		t.Fatalf(`expected status 405 but got %v`, w.status)
	}
}

func TestKneeboard(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf(`got error %v`, err.Error())
	}
	server, err := LoadServerFromSettings(`settings_test.json`)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	server.ServeFolder = wd
	router := server.GenerateRouter()
	w := &testWriter{}
	r := getRequestFor(`https://something.somewhere.com/kneeboard?from=ORD&to=TTA`)
	r.Method = `GET`
	router.ServeHTTP(w, r)
	t.Log(string(w.content))
	if w.status != 200 {
		t.Fatalf(`expected 200 but got %v`, w.status)
	}
}

func checkRoute(router *mux.Router, url string) error {
	match := &mux.RouteMatch{}
	r := getRequestFor(url)
	if success := router.Match(r, match); !success {
		return match.MatchErr
	}
	return nil
}

func checkApiRoute(router *mux.Router, url string) error {
	match := &mux.RouteMatch{}
	r := requestFor(url, `POST`)
	router.Match(r, match)
	return match.MatchErr
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
	return requestFor(testUrl, `GET`)
}

func requestFor(testUrl string, method string) *http.Request {
	u, _ := url.Parse(testUrl)
	return &http.Request{
		Method: method,
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
