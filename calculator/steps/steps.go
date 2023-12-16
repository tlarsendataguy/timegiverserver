package steps

import (
	"fmt"
	"strings"
	"time"
	"timegiverserver/calculator/steps/localization"
	"timegiverserver/lang"
)

const url = `https://www.timegiver.app`

type Step interface {
	BuildIcs(builder *strings.Builder, lang lang.Lang)
}

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
	wrap(s, "DESCRIPTION:", fmt.Sprintf(`%v\n\nGet another plan at %v`, texts.Description, url))
	s.Write([]byte("\r\n"))
}

func writeUrl(s *strings.Builder) {
	s.Write([]byte(fmt.Sprintf("URL:%v\r\n", url)))
}

func writeAlarm(s *strings.Builder, texts localization.StepText) {
	s.Write([]byte("BEGIN:VALARM\r\n"))
	s.Write([]byte("TRIGGER;RELATED=START:PT0M\r\n"))
	s.Write([]byte("ACTION:DISPLAY\r\n"))
	wrap(s, "DESCRIPTION:", texts.Title)
	s.Write([]byte("\r\nEND:VALARM\r\n"))
}

func endEvent(s *strings.Builder) {
	s.Write([]byte("CATEGORIES:TimeGiver\r\n"))
	s.Write([]byte("END:VEVENT\r\n"))
}

func buildIcs(builder *strings.Builder, stepId string, texts localization.StepText, start, end time.Time) {
	beginEvent(builder)
	writeUid(builder, start, stepId)
	writeDates(builder, start, end)
	writeTexts(builder, texts)
	writeUrl(builder)
	writeAlarm(builder, texts)
	endEvent(builder)
}

func min(left, right int) int {
	if left < right {
		return left
	}
	return right
}

type NoCaffeine struct {
	At time.Time
}

func (s NoCaffeine) BuildIcs(builder *strings.Builder, lang lang.Lang) {
	buildIcs(builder, `NoCaffeine`, localization.NoCaffeine[lang], s.At, s.At)
}

type CaffeineOk struct {
	Start time.Time
	End   time.Time
}

func (s CaffeineOk) BuildIcs(builder *strings.Builder, lang lang.Lang) {
	buildIcs(builder, `CaffeineOk`, localization.CaffeineOk[lang], s.Start, s.End)
}

type Caffeine3C struct {
	Start time.Time
	End   time.Time
}

func (s Caffeine3C) BuildIcs(builder *strings.Builder, lang lang.Lang) {
	buildIcs(builder, `Caffeine3C`, localization.Caffeine3C[lang], s.Start, s.End)
}

type Caffeine2C struct {
	Start time.Time
	End   time.Time
}

func (s Caffeine2C) BuildIcs(builder *strings.Builder, lang lang.Lang) {
	buildIcs(builder, `Caffeine2C`, localization.Caffeine2C[lang], s.Start, s.End)
}

type LightBreakfast struct {
	At time.Time
}

func (s LightBreakfast) BuildIcs(builder *strings.Builder, lang lang.Lang) {
	buildIcs(builder, `LightBreakfast`, localization.LightBreakfast[lang], s.At, s.At)
}

type LightLunch struct {
	At time.Time
}

func (s LightLunch) BuildIcs(builder *strings.Builder, lang lang.Lang) {
	buildIcs(builder, `LightLunch`, localization.LightLunch[lang], s.At, s.At)
}

type LightDinner struct {
	At time.Time
}

func (s LightDinner) BuildIcs(builder *strings.Builder, lang lang.Lang) {
	buildIcs(builder, `LightDinner`, localization.LightDinner[lang], s.At, s.At)
}

type LightDinnerOptional struct {
	At time.Time
}

func (s LightDinnerOptional) BuildIcs(builder *strings.Builder, lang lang.Lang) {
	buildIcs(builder, `LightDinnerOptional`, localization.LightDinnerOptional[lang], s.At, s.At)
}

type HeavyBreakfast struct {
	At time.Time
}

func (s HeavyBreakfast) BuildIcs(builder *strings.Builder, lang lang.Lang) {
	buildIcs(builder, `HeavyBreakfast`, localization.HeavyBreakfast[lang], s.At, s.At)
}

type HeavyLunch struct {
	At time.Time
}

func (s HeavyLunch) BuildIcs(builder *strings.Builder, lang lang.Lang) {
	buildIcs(builder, `HeavyLunch`, localization.HeavyLunch[lang], s.At, s.At)
}

type HeavyDinner struct {
	At time.Time
}

func (s HeavyDinner) BuildIcs(builder *strings.Builder, lang lang.Lang) {
	buildIcs(builder, `HeavyDinner`, localization.HeavyDinner[lang], s.At, s.At)
}

type Sleep struct {
	Start time.Time
	End   time.Time
}

func (s Sleep) BuildIcs(builder *strings.Builder, lang lang.Lang) {
	buildIcs(builder, `Sleep`, localization.Sleep[lang], s.Start, s.End)
}

type SetWatch struct {
	At time.Time
}

func (s SetWatch) BuildIcs(builder *strings.Builder, lang lang.Lang) {
	buildIcs(builder, `SetWatch`, localization.SetWatch[lang], s.At, s.At)
}
