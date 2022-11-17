package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/jordan-wright/email"
	_ "github.com/snowflakedb/gosnowflake"
	"io"
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
	w.Header().Add("Content-Type", `text/html`)
	_, _ = w.Write(content)
}

func (s *Server) HandleFile(w http.ResponseWriter, r *http.Request) {
	fullPath := path.Join(s.ServeFolder, r.URL.Path)
	content, err := os.ReadFile(fullPath)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	mimeType := mime.TypeByExtension(filepath.Ext(fullPath))
	w.Header().Add("Content-Type", mimeType)
	_, _ = w.Write(content)
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
	ics := s.generateIcs(params, langValue)
	if to := params.Email; to != `` {
		err = s.emailPlan(ics, to)
		if err != nil {
			writeError(w, err)
			return
		}
	}
	w.Header().Add("Content-Type", "text/calendar")
	_, _ = w.Write([]byte(ics))
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
	w.WriteHeader(500)
	_, _ = w.Write([]byte(err.Error()))
}

func getLangFromRequest(r *http.Request) lang.Lang {
	langStr := r.Header.Get(http.CanonicalHeaderKey(`accept-language`))
	return lang.ParseLang(langStr)
}
