package main

import (
	"crypto/tls"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/acme/autocert"
	"io/fs"
	"net/http"
	"os"
	"time"
	"timegiverserver/handlers"
)

func main() {
	logger, err := os.OpenFile(`./Log.txt`, os.O_CREATE|os.O_TRUNC, fs.ModePerm)
	if err != nil {
		println(err.Error())
		return
	}
	settings, err := handlers.LoadServerFromSettings(`settings.json`, logger, `PROD`)
	if err != nil {
		return
	}
	e := generateRouter(settings)
	m := &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		Cache:      autocert.DirCache(settings.CertFolder),
		HostPolicy: autocert.HostWhitelist(`timegiver.app`, `www.timegiver.app`),
	}
	serveTls := &http.Server{
		Addr:         `:443`,
		Handler:      e,
		TLSConfig:    &tls.Config{GetCertificate: m.GetCertificate},
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	serve := &http.Server{
		Addr:         `:80`,
		Handler:      m.HTTPHandler(nil),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	go func() {
		redirectErr := serve.ListenAndServe()
		settings.Log(redirectErr.Error())
	}()
	err = serveTls.ListenAndServeTLS(``, ``)
	settings.Log(err.Error())
}

func generateRouter(settings *handlers.Server) *mux.Router {
	e := mux.NewRouter()
	e.HandleFunc(`/`, settings.HandleHomepage)
	e.HandleFunc(`/api/calculate`, settings.HandleCalculateApi).Methods(`POST`)
	e.PathPrefix(`/`).HandlerFunc(settings.HandleFile).Methods(`GET`)
	return e
}
