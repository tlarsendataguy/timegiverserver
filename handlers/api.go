package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"timegiverserver/handlers/kneeboard"
	"timegiverserver/lang"
)

func (s *Server) handleCalculateApi(w http.ResponseWriter, r *http.Request) {
	var params CalcPayload
	var langValue lang.Lang
	var err error
	if s.Db != nil {
		defer func() {
			insertErr := s.insertApiRequest(params, langValue, err)
			if insertErr != nil {
				log.Printf(`error writing to Snowflake: %v`, insertErr.Error())
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

func (s *Server) handleTimezoneApi(w http.ResponseWriter, r *http.Request) {
	var payload TimezoneRequestPayload
	d := json.NewDecoder(r.Body)
	err := d.Decode(&payload)
	if err != nil {
		writeErrorMsg(w, `JSON payload is invalid`)
		return
	}
	timestamp, err := time.Parse(`2006-01-02T15:04`, payload.Timestamp)
	if err != nil {
		writeErrorMsg(w, `Timestamp is not formatted correctly`)
		return
	}

	fromOffset, err := s.requestTimezone(payload.From, timestamp)
	if err != nil {
		writeErrorMsg(w, fmt.Sprintf(`error obtaining departure timezone: %v`, err.Error()))
		return
	}
	toOffset, err := s.requestTimezone(payload.To, timestamp)
	if err != nil {
		writeErrorMsg(w, fmt.Sprintf(`error obtaining arrival timezone: %v`, err.Error()))
		return
	}

	response := TimezoneResponsePayload{
		FromOffset: fromOffset.Offset(),
		ToOffset:   toOffset.Offset(),
	}
	responseBytes, err := json.Marshal(response)
	if err != nil {
		writeErrorMsg(w, `the server did not marshall the JSON correctly`)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	_, _ = w.Write(responseBytes)
}

func (s *Server) handleKneeboardApi(w http.ResponseWriter, r *http.Request) {
	rootPath := filepath.Join(s.ServeFolder, `kneeboard`)
	variables := r.URL.Query()
	from := variables.Get(`from`)
	to := variables.Get(`to`)
	kb := &kneeboard.Kneeboard{From: from, To: to}

	result, err := s.Db.Query(`SELECT * FROM NASR.PUBLIC.FREQUENCIES WHERE FACILITY IN (?, ?)`, from, to)
	if err != nil {
		writeError(w, err)
		return
	}
	err = kb.LoadFrequencies(result)
	if err != nil {
		writeError(w, err)
		return
	}

	result, err = s.Db.Query(`SELECT * FROM NASR.PUBLIC.RUNWAY_BOXES WHERE FACILITY in (?, ?)`, from, to)
	if err != nil {
		writeError(w, err)
		return
	}
	err = kb.LoadRunways(result)
	if err != nil {
		writeError(w, err)
		return
	}

	result, err = s.Db.Query(`SELECT "Location Identifier", "Official Facility Name", "Airport Elevation" FROM NASR.PUBLIC.APT WHERE "Location Identifier" IN (?, ?)`, from, to)
	if err != nil {
		writeError(w, err)
		return
	}
	err = kb.LoadInfo(result)
	if err != nil {
		writeError(w, err)
		return
	}

	var f *os.File
	if from != `` {
		f, err = os.OpenFile(filepath.Join(rootPath, from+`.png`), os.O_CREATE, os.ModePerm)
		if err != nil {
			writeError(w, err)
			return
		}
		err = kb.CreateRwyImage(from, f)
		if err != nil {
			writeError(w, err)
			return
		}
		_ = f.Close()
	}

	if to != `` {
		f, err = os.OpenFile(filepath.Join(rootPath, to+`.png`), os.O_CREATE, os.ModePerm)
		if err != nil {
			writeError(w, err)
			return
		}
		err = kb.CreateRwyImage(to, f)
		if err != nil {
			writeError(w, err)
			return
		}
		_ = f.Close()
	}

	_, _ = w.Write(kb.BuildHtml())
}
