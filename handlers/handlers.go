package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jordan-wright/email"
	_ "github.com/snowflakedb/gosnowflake"
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

type HostInfo struct {
	Folder            string
	HostWhitelist     []string
	AllowTimezonesApi bool
	AllowCalculateApi bool
	AllowKneeboardApi bool
}

type Server struct {
	CertFolder  string
	ServeFolder string
	Emailer     Smtp
	DbConnStr   string
	Db          *sql.DB
	MapsApiKey  string
	Hosts       []HostInfo
	env         string
}

func LoadServerFromSettings(settingsFilePath string, environment string) (*Server, error) {
	settings := &Server{env: environment}

	content, err := os.ReadFile(settingsFilePath)
	if err != nil {
		log.Printf(`error reading settings file: %v`, err.Error())
		return nil, err
	}
	err = json.Unmarshal(content, settings)
	if err != nil {
		log.Printf(`error parsing settings file: %v`, err.Error())
		return nil, err
	}
	settings.Emailer.Address = fmt.Sprintf(`%v:%v`, settings.Emailer.Host, settings.Emailer.Port)
	settings.Emailer.Auth = smtp.PlainAuth(``, settings.Emailer.Username, settings.Emailer.Password, settings.Emailer.Host)
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

			if info.AllowCalculateApi {
				sub.Path(`/api/calculate`).Methods(`POST`).HandlerFunc(s.handleCalculateApi)
			}
			if info.AllowTimezonesApi {
				sub.Path(`/api/timezones`).Methods(`POST`).HandlerFunc(s.handleTimezoneApi)
			}
			if info.AllowKneeboardApi {
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

func (s *Server) generateIcs(params CalcPayload, langValue lang.Lang) string {
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
	return calculator.BuildIcsFile(steps, langValue)
}

func (s *Server) emailPlan(ics string, to string) error {
	e := email.NewEmail()
	e.From = s.Emailer.From
	e.To = []string{to}
	e.Subject = `Timegiver Plan`
	e.Text = []byte("Attached is your Timegiver plan to beat jet lag on your upcoming trip!\n\n")
	_, err := e.Attach(strings.NewReader(ics), `Timegiver Plan.ics`, `text/calendar`)
	if err != nil {
		return err
	}
	err = e.Send(s.Emailer.Address, s.Emailer.Auth)
	if err != nil {
		return err
	}
	return nil
}

const insert = `INSERT INTO TIMEGIVER.PUBLIC.API_USAGE (EXECUTED,DEPARTURE_OFFSET,ARRIVAL_OFFSET,ARRIVAL_TIME,LANG,DEPARTURE_LOC,ARRIVAL_LOC,WAKE,BREAKFAST,LUNCH,DINNER,SLEEP,ENVIRONMENT,ERRORS) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)`

func (s *Server) insertApiRequest(params CalcPayload, langValue lang.Lang, handleErr error) error {
	now := time.Now()
	handleErrStr := ``
	if handleErr != nil {
		handleErrStr = handleErr.Error()
	}
	_, err := s.Db.Exec(insert, now, params.DepartureOffset, params.ArrivalOffset, params.Arrival, langValue.String(), params.DepartureLoc, params.ArrivalLoc, params.Wake, params.Breakfast, params.Lunch, params.Dinner, params.Sleep, s.env, handleErrStr)
	return err
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

func getLangFromRequest(r *http.Request) lang.Lang {
	langStr := r.Header.Get(http.CanonicalHeaderKey(`accept-language`))
	return lang.ParseLang(langStr)
}
