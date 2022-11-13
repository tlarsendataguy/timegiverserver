package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/jordan-wright/email"
	_ "github.com/snowflakedb/gosnowflake"
	"io"
	"net/http"
	"net/smtp"
	"os"
	"path"
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

type Server struct {
	CertFolder  string
	ServeFolder string
	Emailer     Smtp
	DbConnStr   string
	Db          *sql.DB
	log         io.Writer
	env         string
}

func (s *Server) Log(format string, a ...interface{}) {
	entry := fmt.Sprintf(format, a...)
	entry = fmt.Sprintf(`%v: %v`, time.Now(), entry)
	_, _ = s.log.Write([]byte(entry))
}

func LoadServerFromSettings(settingsFilePath string, logger io.Writer, environment string) (*Server, error) {
	settings := &Server{log: logger, env: environment}

	content, err := os.ReadFile(settingsFilePath)
	if err != nil {
		settings.Log(`error reading settings file: %v`, err.Error())
		return nil, err
	}
	err = json.Unmarshal(content, settings)
	if err != nil {
		settings.Log(`error parsing settings file: %v`, err.Error())
		return nil, err
	}
	settings.Emailer.Address = fmt.Sprintf(`%v:%v`, settings.Emailer.Host, settings.Emailer.Port)
	settings.Emailer.Auth = smtp.PlainAuth(``, settings.Emailer.Username, settings.Emailer.Password, settings.Emailer.Host)
	if settings.DbConnStr != `` {
		settings.Log(`persisting to Snowflake`)
		settings.Db, err = sql.Open(`snowflake`, settings.DbConnStr)
	}
	return settings, err
}

func (s *Server) HandleHomepage(w http.ResponseWriter, _ *http.Request) {
	fullPath := path.Join(s.ServeFolder, `index.html`)
	content, err := os.ReadFile(fullPath)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	_, _ = w.Write(content)
}

func (s *Server) HandleFile(w http.ResponseWriter, r *http.Request) {
	fullPath := path.Join(s.ServeFolder, r.URL.Path)
	content, err := os.ReadFile(fullPath)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	_, _ = w.Write(content)
	w.WriteHeader(200)
}

func (s *Server) HandleCalculateApi(w http.ResponseWriter, r *http.Request) {
	var params CalcPayload
	var langValue lang.Lang
	var err error
	if s.Db != nil {
		defer func() {
			insertErr := s.insertApiRequest(params, langValue, err)
			if insertErr != nil {
				s.Log(`error writing to Snowflake: %v`, insertErr.Error())
			}
		}()
	}

	langValue = getLangFromRequest(r)
	params, err = ValidateCalcPayload(r.Body)
	if err != nil {
		writeError(w, err)
		return
	}
	err = s.emailPlan(params, langValue)
	if err != nil {
		writeError(w, err)
		return
	}
	writeSuccess(w)
}

func (s *Server) emailPlan(params CalcPayload, langValue lang.Lang) error {
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
	_, err := e.Attach(strings.NewReader(ics), `plan.ics`, `text/calendar`)
	if err != nil {
		return err
	}
	err = e.Send(s.Emailer.Address, s.Emailer.Auth)
	if err != nil {
		return err
	}
	return nil
}

const insert = `INSERT INTO API_USAGE (EXECUTED,DEPARTURE_OFFSET,ARRIVAL_OFFSET,ARRIVAL_TIME,LANG,DEPARTURE_LOC,ARRIVAL_LOC,WAKE,BREAKFAST,LUNCH,DINNER,SLEEP,ENVIRONMENT,ERRORS) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)`

func (s *Server) insertApiRequest(params CalcPayload, langValue lang.Lang, handleErr error) error {
	now := time.Now()
	handleErrStr := ``
	if handleErr != nil {
		handleErrStr = handleErr.Error()
	}
	_, err := s.Db.Exec(insert, now, params.DepartureOffset, params.ArrivalOffset, params.Arrival, langValue.String(), ``, ``, params.Wake, params.Breakfast, params.Lunch, params.Dinner, params.Sleep, s.env, handleErrStr)
	return err
}

func writeError(w http.ResponseWriter, err error) {
	_, _ = w.Write([]byte(err.Error()))
	w.WriteHeader(500)
}

func writeSuccess(w http.ResponseWriter) {
	w.WriteHeader(200)
}

func getLangFromRequest(r *http.Request) lang.Lang {
	langStr := r.Header.Get(http.CanonicalHeaderKey(`accept-language`))
	return lang.ParseLang(langStr)
}
