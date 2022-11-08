package main

import (
	"crypto/tls"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/acme/autocert"
	"net/http"
	"time"
)

func main() {
	e := mux.NewRouter()
	e.HandleFunc(`/api/calculate`, nil).Methods(`POST`)
	m := &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		Cache:      autocert.DirCache(`certs`),
		HostPolicy: autocert.HostWhitelist(`www.timegiver.app`),
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
		err := serve.ListenAndServe()
		println(err.Error())
	}()
	err := serveTls.ListenAndServeTLS(``, ``)
	println(err.Error())
}
