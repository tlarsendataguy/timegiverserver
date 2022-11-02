package steps

import (
	"strings"
	"time"
	"timegiverserver/steps/localization"
)

func FmtDt(value time.Time) string {
	return value.Format(`20060102T150400Z`)
}

func Wrap(header, content string) string {
	builder := strings.Builder{}
	toWrite := append([]byte(header), []byte(content)...)
	size := len(toWrite)
	position := 0
	remainder := size

	for remainder > 0 {
		amount := min(remainder, 75)
		builder.Write(toWrite[position : position+amount])
		if remainder > 75 {
			builder.Write([]byte("\r\n "))
		}
		position += amount
		remainder -= amount
	}

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
	builder := strings.Builder{}
	builder.Write([]byte("BEGIN:VEVENT\r\n"))
	builder.Write([]byte("UID:"))
	builder.Write([]byte(FmtDt(s.Start)))
	builder.Write([]byte("NoCaffeine@timegiver.app\r\n"))
	builder.Write([]byte("DTSTAMP:"))
	builder.Write([]byte(FmtDt(s.Start)))
	builder.Write([]byte("\r\nDTSTART:"))
	builder.Write([]byte(FmtDt(s.Start)))
	builder.Write([]byte("\r\nDTEND:"))
	builder.Write([]byte(FmtDt(s.End)))
	builder.Write([]byte("\r\n"))
	builder.Write([]byte(Wrap("SUMMARY:", localization.NoCaffeine[lang].Title)))
	builder.Write([]byte("\r\n"))
	builder.Write([]byte(Wrap("DESCRIPTION:", localization.NoCaffeine[lang].Description)))
	builder.Write([]byte("\r\nCATEGORIES:TimeGiver\r\n"))
	builder.Write([]byte("END:VEVENT\r\n"))
	return builder.String()
}

type CaffeineOk struct {
	Start time.Time
	End   time.Time
}

func (s CaffeineOk) ToIcs(lang localization.Lang) string {
	return ``
}

type Caffeine3C struct {
	Start time.Time
	End   time.Time
}

func (s Caffeine3C) ToIcs(lang localization.Lang) string {
	return ``
}

type Caffeine2C struct {
	Start time.Time
	End   time.Time
}

func (s Caffeine2C) ToIcs(lang localization.Lang) string {
	return ``
}

type LightBreakfast struct {
	At time.Time
}

func (s LightBreakfast) ToIcs(lang localization.Lang) string {
	return ``
}

type LightLunch struct {
	At time.Time
}

func (s LightLunch) ToIcs(lang localization.Lang) string {
	return ``
}

type LightDinner struct {
	At time.Time
}

func (s LightDinner) ToIcs(lang localization.Lang) string {
	return ``
}

type LightDinnerOptional struct {
	At time.Time
}

func (s LightDinnerOptional) ToIcs(lang localization.Lang) string {
	return ``
}

type HeavyBreakfast struct {
	At time.Time
}

func (s HeavyBreakfast) ToIcs(lang localization.Lang) string {
	return ``
}

type HeavyLunch struct {
	At time.Time
}

func (s HeavyLunch) ToIcs(lang localization.Lang) string {
	return ``
}

type HeavyDinner struct {
	At time.Time
}

func (s HeavyDinner) ToIcs(lang localization.Lang) string {
	return ``
}

type NoSnack struct {
	Start time.Time
	End   time.Time
}

func (s NoSnack) ToIcs(lang localization.Lang) string {
	return ``
}

type Sleep struct {
	Start time.Time
	End   time.Time
}

func (s Sleep) ToIcs(lang localization.Lang) string {
	return ``
}

type NoNap struct {
	Start time.Time
	End   time.Time
}

func (s NoNap) ToIcs(lang localization.Lang) string {
	return ``
}

type SetWatch struct {
	At time.Time
}

func (s SetWatch) ToIcs(lang localization.Lang) string {
	return ``
}

type Arrive struct {
	At time.Time
}

func (s Arrive) ToIcs(lang localization.Lang) string {
	return ``
}
