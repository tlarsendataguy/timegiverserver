package calculator

import (
	"strings"
	"timegiverserver/calculator/steps"
	"timegiverserver/lang"
)

type PlanIcs struct {
	Caffeine string
	Meals    string
	Sleep    string
	Events   string
}

func BuildIcsFiles(plan PlanSteps) PlanIcs {
	files := PlanIcs{}
	files.Caffeine = BuildIcsFile(plan.Caffeine, lang.EN)
	files.Meals = BuildIcsFile(plan.Meals, lang.EN)
	files.Sleep = BuildIcsFile(plan.Sleep, lang.EN)
	files.Events = BuildIcsFile(plan.Events, lang.EN)
	return files
}

func BuildIcsFile(plan []steps.Step, lang lang.Lang) string {
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
	builder.Write([]byte("X-WR-CALNAME: Timegiver\r\n"))
	builder.Write([]byte("NAME: Timegiver\r\n"))
}

func endCalendar(builder *strings.Builder) {
	builder.Write([]byte("END:VCALENDAR\r\n"))
}
