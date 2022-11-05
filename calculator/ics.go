package calculator

import (
	"strings"
	"timegiverserver/calculator/steps"
	"timegiverserver/calculator/steps/localization"
)

func BuildIcsFile(plan []steps.Step, lang localization.Lang) string {
	builder := &strings.Builder{}
	beginCalendar(builder)
	for _, step := range plan {
		step.BuildIcs(builder, lang)
	}
	endCalendar(builder)
	return builder.String()
}

func beginCalendar(builder *strings.Builder) {
	builder.Write([]byte("BEGIN:VCALENDAR\r\n"))
	builder.Write([]byte("VERSION:2.0\r\n"))
	builder.Write([]byte("PRODID://tlarsendataguy//timegiver\r\n"))
}

func endCalendar(builder *strings.Builder) {
	builder.Write([]byte("END:VCALENDAR\r\n"))
}
