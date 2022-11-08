package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestLoadHandler(t *testing.T) {
	settings, err := LoadSettings()
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if value := settings.CertFolder; value != `./certs` {
		t.Fatalf(`expected './certs' but got '%v'`, value)
	}
	if value := settings.ServeFolder; value != `./serveTest` {
		t.Fatalf(`expected './serveTest' but got '%v'`, value)
	}
}

func TestLoadHomepage(t *testing.T) {
	settings, _ := LoadSettings()
	w := &testWriter{}
	settings.HandleHomepage(w, nil)

	err := checkResponse(w, 200, `./serveTest/index.html`)
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestLoadFile(t *testing.T) {
	settings, _ := LoadSettings()
	w := &testWriter{}
	r := getRequestFor(`https://www.timegiver.app/scripts.js`)
	settings.HandleFile(w, r)

	err := checkResponse(w, 200, `./serveTest/scripts.js`)
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func checkResponse(w *testWriter, expectedStatus int, expectedFile string) error {
	if w.status != expectedStatus {
		return fmt.Errorf(`expected status %v but got %v`, expectedStatus, w.status)
	}
	expected, _ := ioutil.ReadFile(expectedFile)
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
}

func (w *testWriter) Header() http.Header {
	return nil
}

func (w *testWriter) Write(content []byte) (int, error) {
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
