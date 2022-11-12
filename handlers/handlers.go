package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/jordan-wright/email"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"path"
	"strings"
	"timegiverserver/calculator"
	"timegiverserver/lang"
)

type Smtp struct {
	Host     string
	Port     string
	Username string
	Password string
	From     string
	Address  string
	Auth     smtp.Auth
}

type Settings struct {
	CertFolder  string
	ServeFolder string
	Emailer     Smtp
}

func LoadSettings(settingsFilePath string) (*Settings, error) {
	content, err := ioutil.ReadFile(settingsFilePath)
	if err != nil {
		return nil, err
	}
	settings := &Settings{}
	err = json.Unmarshal(content, settings)
	if err != nil {
		return nil, err
	}
	settings.Emailer.Address = fmt.Sprintf(`%v:%v`, settings.Emailer.Host, settings.Emailer.Port)
	settings.Emailer.Auth = smtp.PlainAuth(``, settings.Emailer.Username, settings.Emailer.Password, settings.Emailer.Host)
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

func (s *Settings) HandleCalculateApi(w http.ResponseWriter, r *http.Request) {
	params, err := ValidateCalcPayload(r.Body)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		w.WriteHeader(400)
		return
	}
	langStr := r.Header.Get(http.CanonicalHeaderKey(`accept-language`))
	langValue := lang.ParseLang(langStr)
	calc := calculator.InitializeCalculator(calculator.Inputs{
		Arrival:         params.Arrival,
		DepartureOffset: params.DepartureOffset,
		ArrivalOffset:   params.ArrivalOffset,
		Routine: calculator.DailyRoutine{
			Wake:      params.Wake,
			Breakfast: params.Breakfast,
			Lunch:     params.Lunch,
			Dinner:    params.Dinner,
			Sleep:     params.Sleep,
		},
	})
	steps := calc.Plan()
	ics := calculator.BuildIcsFile(steps, langValue)
	e := email.NewEmail()
	e.From = s.Emailer.From
	e.To = []string{params.Email}
	e.Subject = `Timegiver Plan`
	e.Text = []byte("Attached is your Timegiver plan to beat jet lag on your upcoming trip!\n\n")
	_, err = e.Attach(strings.NewReader(ics), `plan.ics`, `text/calendar`)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		w.WriteHeader(500)
		return
	}
	err = e.Send(s.Emailer.Address, s.Emailer.Auth)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
}
