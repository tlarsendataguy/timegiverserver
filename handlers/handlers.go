package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jordan-wright/email"
	_ "github.com/snowflakedb/gosnowflake"
	"github.com/stripe/stripe-go/v76"
	"log"
	"mime"
	"net/http"
	"net/smtp"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
	"timegiverserver/calculator"
)

type Smtp struct {
	Host    string
	Port    string
	From    string
	Address string
	Auth    smtp.Auth
}

type HostInfo struct {
	Folder          string
	HostWhitelist   []string
	HasTimegiverApi bool
	HasKneeboardApi bool
}

type Server struct {
	CertFolder           string
	ServeFolder          string
	Emailer              Smtp
	DbConnStr            string
	Db                   *sql.DB
	MapsApiKey           string
	Hosts                []HostInfo
	Test                 bool
	StripeKey            string
	WebhookSigningSecret string
}

func LoadServerFromSettings(filename string) (*Server, error) {
	dbUser := os.Getenv(`DB_USER`)
	dbPassword := os.Getenv(`DB_PASSWORD`)
	smtpUser := os.Getenv(`SMTP_USER`)
	smtpPassword := os.Getenv(`SMTP_PASSWORD`)
	stripeKey := os.Getenv(`STRIPE_KEY`)
	mapsApiKey := os.Getenv(`MAPS_API_KEY`)
	webhookSigningSecret := os.Getenv(`WEBHOOK_SIGNING_SECRET`)

	stripe.Key = stripeKey

	content, err := os.ReadFile(filename)
	if err != nil {
		log.Printf(`error reading settings file: %v`, err.Error())
		return nil, err
	}
	settings := &Server{}
	err = json.Unmarshal(content, settings)
	if err != nil {
		log.Printf(`error parsing settings file: %v`, err.Error())
		return nil, err
	}
	settings.DbConnStr = fmt.Sprintf(`%v:%v@%v`, dbUser, dbPassword, settings.DbConnStr)
	settings.StripeKey = stripeKey
	settings.Emailer.Address = fmt.Sprintf(`%v:%v`, settings.Emailer.Host, settings.Emailer.Port)
	settings.Emailer.Auth = smtp.PlainAuth(``, smtpUser, smtpPassword, settings.Emailer.Host)
	settings.MapsApiKey = mapsApiKey
	settings.WebhookSigningSecret = webhookSigningSecret
	if settings.DbConnStr != `` {
		log.Printf(`persisting to Snowflake`)
		settings.Db, err = sql.Open(`snowflake`, settings.DbConnStr)
		go settings.keepSnowflakeAlive()
	}
	return settings, err
}

func (s *Server) homepageHandler(hostFolder string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fullPath := path.Join(s.ServeFolder, hostFolder, `index.html`)
		content, err := os.ReadFile(fullPath)
		if err != nil {
			s.handle404(hostFolder, w)
			return
		}
		w.Header().Add("Content-Type", `text/html`)
		_, _ = w.Write(content)
	}
}

func (s *Server) fileHandler(hostFolder string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fullPath := path.Join(s.ServeFolder, hostFolder, r.URL.Path)
		content, err := os.ReadFile(fullPath)
		if err != nil {
			s.handle404(hostFolder, w)
			return
		}
		mimeType := mime.TypeByExtension(filepath.Ext(fullPath))
		w.Header().Add("Content-Type", mimeType)
		_, _ = w.Write(content)
	}
}

func (s *Server) handle404(hostFolder string, w http.ResponseWriter) {
	err404, _ := os.ReadFile(path.Join(s.ServeFolder, hostFolder, `404.html`))
	w.WriteHeader(404)
	_, _ = w.Write(err404)
}

func (s *Server) notFoundHandler(hostFolder string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		s.handle404(hostFolder, w)
	}
}

func (s *Server) methodNotAllowedHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(405)
}

func (s *Server) CollectHostWhitelist() []string {
	whitelist := make([]string, 0, 10)
	for _, hosts := range s.Hosts {
		whitelist = append(whitelist, hosts.HostWhitelist...)
	}
	return whitelist
}

func (s *Server) GenerateRouter() *mux.Router {
	e := mux.NewRouter()

	for _, info := range s.Hosts {
		for _, host := range info.HostWhitelist {
			sub := e.Host(host).Subrouter()
			sub.Path(`/`).Methods(`GET`).HandlerFunc(s.homepageHandler(info.Folder))

			if info.HasTimegiverApi {
				sub.Path(`/webhook`).Methods(`POST`).HandlerFunc(s.handleSessionWebhook)
				sub.Path(`/checkout`).Methods(`POST`).HandlerFunc(s.handleCheckout)
				sub.Path(`/api/timezones`).Methods(`POST`).HandlerFunc(s.handleTimezoneApi)
			}
			if info.HasKneeboardApi {
				sub.Path(`/kneeboard`).Methods(`GET`).HandlerFunc(s.handleKneeboardApi)
			}

			sub.PathPrefix(`/.git`).Methods(`GET`).HandlerFunc(s.notFoundHandler(info.Folder))
			sub.PathPrefix(`/`).Methods(`GET`).HandlerFunc(s.fileHandler(info.Folder))
			sub.NotFoundHandler = http.HandlerFunc(s.notFoundHandler(info.Folder))
			sub.MethodNotAllowedHandler = http.HandlerFunc(s.methodNotAllowedHandler)
		}
	}

	return e
}

func (s *Server) requestTimezone(coordinates Coordinates, timestamp time.Time) (*googleTimezoneResponse, error) {
	url := s.buildTimezoneUrl(coordinates, timestamp)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	googleResponse := &googleTimezoneResponse{}
	d := json.NewDecoder(response.Body)
	err = d.Decode(googleResponse)
	_ = response.Body.Close()
	if err != nil {
		return nil, err
	}
	return googleResponse, nil
}

func (s *Server) buildTimezoneUrl(coords Coordinates, timestamp time.Time) string {
	return fmt.Sprintf(`https://maps.googleapis.com/maps/api/timezone/json?location=%v%%2C%v&timestamp=%v&key=%v`, coords.Lat, coords.Lng, timestamp.Unix(), s.MapsApiKey)
}

func (s *Server) generateIcs(params CalcPayload) calculator.PlanIcs {
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
	return calculator.BuildIcsFiles(steps)
}

func (s *Server) emailPlan(ics calculator.PlanIcs, to string) error {
	body, err := os.ReadFile(`./templates/timegiver.html`)
	if err != nil {
		return err
	}
	e := email.NewEmail()
	e.From = s.Emailer.From
	e.To = []string{to}
	e.Subject = `Timegiver Plan`
	e.HTML = body
	e.Text = []byte{}
	_, err = e.Attach(strings.NewReader(ics.Caffeine), `Caffeine Steps.ics`, `text/calendar`)
	if err != nil {
		return err
	}
	_, err = e.Attach(strings.NewReader(ics.Meals), `Meal Steps.ics`, `text/calendar`)
	if err != nil {
		return err
	}
	_, err = e.Attach(strings.NewReader(ics.Sleep), `Sleep Steps.ics`, `text/calendar`)
	if err != nil {
		return err
	}
	_, err = e.Attach(strings.NewReader(ics.Events), `Set Watch.ics`, `text/calendar`)
	if err != nil {
		return err
	}
	err = e.Send(s.Emailer.Address, s.Emailer.Auth)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) keepSnowflakeAlive() {
	for {
		time.Sleep(1 * time.Hour)
		rows, err := s.Db.Query(`SELECT 1`)
		if err != nil {
			log.Printf(err.Error())
			continue
		}
		_ = rows.Close()
	}
}

func writeError(w http.ResponseWriter, err error) {
	writeErrorMsg(w, err.Error())
}

func writeErrorMsg(w http.ResponseWriter, msg string) {
	w.WriteHeader(500)
	_, _ = w.Write([]byte(msg))
}
