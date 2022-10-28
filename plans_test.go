package timegiverserver

import (
	"fmt"
	"testing"
	"time"
)

var arrival = time.Date(2017, 1, 15, 12, 0, 0, 0, time.UTC)
var routine = DailyRoutine{Wake: time.Hour * 6, Breakfast: time.Hour * 7, Lunch: time.Hour * 12, Dinner: time.Hour * 17, Sleep: time.Hour * 22}

func TestEast12(t *testing.T) {
	plan := generatePlan(0, 2)
	if err := checkNoCaffeine(plan[0], `2017-01-13 06:00`, `2017-01-13 15:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan[1], `2017-01-13 15:00`, `2017-01-13 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoCaffeine(plan[2], `2017-01-13 16:30`, `2017-01-14 18:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeine3C(plan[3], `2017-01-14 18:00`, `2017-01-14 19:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoCaffeine(plan[4], `2017-01-14 19:00`, `2017-01-15 22:00`); err != nil {
		t.Fatal(err.Error())
	}
}

func checkNoCaffeine(item interface{}, start, end string) error {
	step, ok := item.(NoCaffeine)
	if !ok {
		return fmt.Errorf(`expected NoCaffeine but got %T`, item)
	}
	return checkDateRange(start, end, step.start, step.end)
}

func checkCaffeineOk(item interface{}, start, end string) error {
	step, ok := item.(CaffeineOk)
	if !ok {
		return fmt.Errorf(`expected CaffeineOk but got %T`, item)
	}
	return checkDateRange(start, end, step.start, step.end)
}

func checkCaffeine3C(item interface{}, start, end string) error {
	step, ok := item.(Caffeine3C)
	if !ok {
		return fmt.Errorf(`expected Caffeine3C but got %T`, item)
	}
	return checkDateRange(start, end, step.start, step.end)
}

func checkDateRange(expectedStart, expectedEnd string, actualStart, actualEnd time.Time) error {
	if !at(actualStart, expectedStart) || !at(actualEnd, expectedEnd) {
		return fmt.Errorf(`expected from %v to %v but got %v to %v`, expectedStart, expectedEnd, actualStart, actualEnd)
	}
	return nil
}

func generatePlan(departureOffset, arrivalOffset float64) []Step {
	calc := InitializeCalculator(Inputs{
		Language:        "EN",
		Arrival:         arrival,
		DepartureOffset: departureOffset,
		ArrivalOffset:   arrivalOffset,
		Routine:         routine,
	})
	return calc.Plan()
}

func at(actual time.Time, expected string) bool {
	expectedDate, _ := time.Parse(`2006-01-02 15:04`, expected)
	return expectedDate.Equal(actual)
}
