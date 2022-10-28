package timegiverserver

import (
	"time"
	"timegiverserver/base"
)

type Step interface {
	ToIcs() string
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

type DailyRoutine struct {
	Wake      time.Duration
	Breakfast time.Duration
	Lunch     time.Duration
	Dinner    time.Duration
	Sleep     time.Duration
}

type Inputs struct {
	Arrival         time.Time
	DepartureOffset float64
	ArrivalOffset   float64
	Routine         DailyRoutine
}

func InitializeCalculator(inputs Inputs) *Calculator {
	calc := &Calculator{Routine: inputs.Routine}
	calc.Dates = base.InitializeDates(inputs.Arrival, inputs.ArrivalOffset, inputs.DepartureOffset)

	return calc
}
