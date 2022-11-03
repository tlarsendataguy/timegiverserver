package calculator

import (
	"fmt"
	"time"
	"timegiverserver/calculator/base"
	"timegiverserver/calculator/shift"
	"timegiverserver/calculator/steps"
	"timegiverserver/calculator/steps/localization"
)

type Step interface {
	ToIcs(lang localization.Lang) string
}

type DailyRoutine struct {
	Wake      time.Duration
	Breakfast time.Duration
	Lunch     time.Duration
	Dinner    time.Duration
	Sleep     time.Duration
}

type Inputs struct {
	Language        string
	Arrival         time.Time
	DepartureOffset float64
	ArrivalOffset   float64
	Routine         DailyRoutine
}

type CalcPlan func(*Calculator) []Step

type Calculator struct {
	Dates    base.Dates
	Routine  DailyRoutine
	calcPlan CalcPlan
}

func (c *Calculator) Plan() []Step {
	return c.calcPlan(c)
}

func (c *Calculator) arrivalStep() Step {
	return steps.Arrive{At: c.Dates.ArriveAt}
}

func (c *Calculator) arrivalAt(t time.Duration) time.Time {
	return c.Dates.Arrival.Add(t)
}

func (c *Calculator) arrivalLess1At(t time.Duration) time.Time {
	return c.Dates.ArrivalLess1.Add(t)
}

func (c *Calculator) arrivalPlus1At(t time.Duration) time.Time {
	return c.Dates.ArrivalPlus1.Add(t)
}

func (c *Calculator) departureAt(t time.Duration) time.Time {
	return c.Dates.Departure.Add(t)
}

func (c *Calculator) departureLess1At(t time.Duration) time.Time {
	return c.Dates.DepartureLess1.Add(t)
}

func (c *Calculator) departureLess2At(t time.Duration) time.Time {
	return c.Dates.DepartureLess2.Add(t)
}

func (c *Calculator) departureLess3At(t time.Duration) time.Time {
	return c.Dates.DepartureLess3.Add(t)
}

func (c *Calculator) departureLess4At(t time.Duration) time.Time {
	return c.Dates.DepartureLess4.Add(t)
}

func (c *Calculator) wake() time.Duration {
	return c.Routine.Wake
}

func (c *Calculator) breakfast() time.Duration {
	return c.Routine.Breakfast
}

func (c *Calculator) lunch() time.Duration {
	return c.Routine.Lunch
}

func (c *Calculator) dinner() time.Duration {
	return c.Routine.Dinner
}

func (c *Calculator) sleep() time.Duration {
	return c.Routine.Sleep
}

func InitializeCalculator(inputs Inputs) *Calculator {
	calc := &Calculator{Routine: inputs.Routine}
	calc.Dates = base.InitializeDates(inputs.Arrival, inputs.DepartureOffset, inputs.ArrivalOffset)
	timezones := shift.CalcTimezoneShift(inputs.DepartureOffset, inputs.ArrivalOffset)
	switch timezones {
	case 1, 2:
		calc.calcPlan = East12
	case 3, 4:
		calc.calcPlan = East34
	case 5, 6:
		calc.calcPlan = East56
	case 7, 8:
		calc.calcPlan = East78
	case 9, 10:
		calc.calcPlan = East910
	case -1, -2:
		calc.calcPlan = West12
	case -3, -4:
		calc.calcPlan = West34
	case -5, -6:
		calc.calcPlan = West56
	case -7, -8:
		calc.calcPlan = West78
	case -9, -10:
		calc.calcPlan = West910
	case -11, -12, 11, 12:
		//If we are going westbound we need to adjust the pre-departure days
		//up by 1 for the plan to be correct
		if timezones < 0 && calc.departureAt(0).Before(calc.arrivalAt(0)) {
			calc.Dates.DepartureLess4 = addDay(calc.Dates.DepartureLess4)
			calc.Dates.DepartureLess3 = addDay(calc.Dates.DepartureLess3)
			calc.Dates.DepartureLess2 = addDay(calc.Dates.DepartureLess2)
			calc.Dates.DepartureLess1 = addDay(calc.Dates.DepartureLess1)
		}
		calc.calcPlan = Both1112
	default:
		panic(fmt.Sprintf(`invalid timezone shift %v`, timezones))
	}
	return calc
}

func addDay(value time.Time) time.Time {
	return value.Add(time.Hour * 24)
}
