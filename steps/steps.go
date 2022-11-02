package steps

import (
	"strings"
	"time"
	"timegiverserver/steps/localization"
)

func fmtDt(s *strings.Builder, value time.Time) {
	s.Write([]byte(value.Format(`20060102T150400Z`)))
}

func wrap(s *strings.Builder, header, content string) {
	toWrite := append([]byte(header), []byte(content)...)
	size := len(toWrite)
	position := 0
	remainder := size

	for remainder > 0 {
		amount := min(remainder, 75)
		s.Write(toWrite[position : position+amount])
		if remainder > 75 {
			s.Write([]byte("\r\n "))
		}
		position += amount
		remainder -= amount
	}
}

func beginEvent(s *strings.Builder) {
	s.Write([]byte("BEGIN:VEVENT\r\n"))
}

func writeUid(s *strings.Builder, start time.Time, stepId string) {
	s.Write([]byte("UID:"))
	fmtDt(s, start)
	s.Write([]byte(stepId))
	s.Write([]byte("@timegiver.app\r\n"))
}

func writeDates(s *strings.Builder, start, end time.Time) {
	s.Write([]byte("DTSTAMP:"))
	fmtDt(s, start)
	s.Write([]byte("\r\nDTSTART:"))
	fmtDt(s, start)
	s.Write([]byte("\r\nDTEND:"))
	fmtDt(s, end)
	s.Write([]byte("\r\n"))
}

func writeTexts(s *strings.Builder, texts localization.StepText) {
	wrap(s, "SUMMARY:", texts.Title)
	s.Write([]byte("\r\n"))
	wrap(s, "DESCRIPTION:", texts.Description)
	s.Write([]byte("\r\n"))
}

func endEvent(s *strings.Builder) {
	s.Write([]byte("CATEGORIES:TimeGiver\r\n"))
	s.Write([]byte("END:VEVENT\r\n"))
}

func toIcs(stepId string, texts localization.StepText, start, end time.Time) string {
	builder := &strings.Builder{}
	beginEvent(builder)
	writeUid(builder, start, stepId)
	writeDates(builder, start, end)
	writeTexts(builder, texts)
	endEvent(builder)
	return builder.String()
}

func min(left, right int) int {
	if left < right {
		return left
	}
	return right
}

type NoCaffeine struct {
	Start time.Time
	End   time.Time
}

func (s NoCaffeine) ToIcs(lang localization.Lang) string {
	return toIcs(`NoCaffeine`, localization.NoCaffeine[lang], s.Start, s.End)
}

type CaffeineOk struct {
	Start time.Time
	End   time.Time
}

func (s CaffeineOk) ToIcs(lang localization.Lang) string {
	return toIcs(`CaffeineOk`, localization.CaffeineOk[lang], s.Start, s.End)
}

type Caffeine3C struct {
	Start time.Time
	End   time.Time
}

func (s Caffeine3C) ToIcs(lang localization.Lang) string {
	return toIcs(`Caffeine3C`, localization.Caffeine3C[lang], s.Start, s.End)
}

type Caffeine2C struct {
	Start time.Time
	End   time.Time
}

func (s Caffeine2C) ToIcs(lang localization.Lang) string {
	return toIcs(`Caffeine2C`, localization.Caffeine2C[lang], s.Start, s.End)
}

type LightBreakfast struct {
	At time.Time
}

func (s LightBreakfast) ToIcs(lang localization.Lang) string {
	return toIcs(`LightBreakfast`, localization.LightBreakfast[lang], s.At, s.At.Add(time.Hour))
}

type LightLunch struct {
	At time.Time
}

func (s LightLunch) ToIcs(lang localization.Lang) string {
	return toIcs(`LightLunch`, localization.LightLunch[lang], s.At, s.At.Add(time.Hour))
}

type LightDinner struct {
	At time.Time
}

func (s LightDinner) ToIcs(lang localization.Lang) string {
	return toIcs(`LightDinner`, localization.LightDinner[lang], s.At, s.At.Add(time.Hour))
}

type LightDinnerOptional struct {
	At time.Time
}

func (s LightDinnerOptional) ToIcs(lang localization.Lang) string {
	return toIcs(`LightDinnerOptional`, localization.LightDinnerOptional[lang], s.At, s.At.Add(time.Hour))
}

type HeavyBreakfast struct {
	At time.Time
}

func (s HeavyBreakfast) ToIcs(lang localization.Lang) string {
	return toIcs(`HeavyBreakfast`, localization.HeavyBreakfast[lang], s.At, s.At.Add(time.Hour))
}

type HeavyLunch struct {
	At time.Time
}

func (s HeavyLunch) ToIcs(lang localization.Lang) string {
	return toIcs(`HeavyLunch`, localization.HeavyLunch[lang], s.At, s.At.Add(time.Hour))
}

type HeavyDinner struct {
	At time.Time
}

func (s HeavyDinner) ToIcs(lang localization.Lang) string {
	return toIcs(`HeavyDinner`, localization.HeavyDinner[lang], s.At, s.At.Add(time.Hour))
}

type NoSnack struct {
	Start time.Time
	End   time.Time
}

func (s NoSnack) ToIcs(lang localization.Lang) string {
	return toIcs(`NoSnack`, localization.NoSnack[lang], s.Start, s.End)
}

type Sleep struct {
	Start time.Time
	End   time.Time
}

func (s Sleep) ToIcs(lang localization.Lang) string {
	return toIcs(`Sleep`, localization.Sleep[lang], s.Start, s.End)
}

type NoNap struct {
	Start time.Time
	End   time.Time
}

func (s NoNap) ToIcs(lang localization.Lang) string {
	return toIcs(`NoNap`, localization.NoNap[lang], s.Start, s.End)
}

type SetWatch struct {
	At time.Time
}

func (s SetWatch) ToIcs(lang localization.Lang) string {
	return toIcs(`SetWatch`, localization.SetWatch[lang], s.At, s.At)
}

type Arrive struct {
	At time.Time
}

func (s Arrive) ToIcs(lang localization.Lang) string {
	return toIcs(`Arrive`, localization.Arrive[lang], s.At, s.At)
}
