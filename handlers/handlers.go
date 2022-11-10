package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path"
)

type Settings struct {
	CertFolder  string
	ServeFolder string
}

func LoadSettings() (*Settings, error) {
	content, err := ioutil.ReadFile(`./settings.json`)
	if err != nil {
		return nil, err
	}
	settings := &Settings{}
	err = json.Unmarshal(content, settings)
	if err != nil {
		return nil, err
	}
	return settings, nil
}

func (s *Settings) HandleHomepage(w http.ResponseWriter, _ *http.Request) {
	fullPath := path.Join(s.ServeFolder, `index.html`)
	content, err := ioutil.ReadFile(fullPath)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	_, _ = w.Write(content)
	w.WriteHeader(200)
}

func (s *Settings) HandleFile(w http.ResponseWriter, r *http.Request) {
	fullPath := path.Join(s.ServeFolder, r.URL.Path)
	content, err := ioutil.ReadFile(fullPath)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	_, _ = w.Write(content)
	w.WriteHeader(200)
}
