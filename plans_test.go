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
	if err := checkLightBreakfast(plan[5], `2017-01-14 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightLunch(plan[6], `2017-01-14 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightDinner(plan[7], `2017-01-14 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyBreakfast(plan[8], `2017-01-15 05:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyLunch(plan[9], `2017-01-15 10:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyDinner(plan[10], `2017-01-15 15:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan[11], `2017-01-13 22:00`, `2017-01-14 06:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan[12], `2017-01-14 22:00`, `2017-01-15 04:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan[13], `2017-01-15 20:00`, `2017-01-16 04:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSetWatch(plan[14], `2017-01-15 05:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkArrive(plan[15], `2017-01-15 10:00`); err != nil {
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

func checkLightBreakfast(item interface{}, at string) error {
	step, ok := item.(LightBreakfast)
	if !ok {
		return fmt.Errorf(`expected LightBreakfast but got %T`, item)
	}
	return checkDateAt(at, step.at)
}

func checkLightLunch(item interface{}, at string) error {
	step, ok := item.(LightLunch)
	if !ok {
		return fmt.Errorf(`expected LightLunch but got %T`, item)
	}
	return checkDateAt(at, step.at)
}

func checkLightDinner(item interface{}, at string) error {
	step, ok := item.(LightDinner)
	if !ok {
		return fmt.Errorf(`expected LightDinner but got %T`, item)
	}
	return checkDateAt(at, step.at)
}

func checkLightDinnerOptional(item interface{}, at string) error {
	step, ok := item.(LightDinnerOptional)
	if !ok {
		return fmt.Errorf(`expected LightDinnerOptionall but got %T`, item)
	}
	return checkDateAt(at, step.at)
}

func checkHeavyBreakfast(item interface{}, at string) error {
	step, ok := item.(HeavyBreakfast)
	if !ok {
		return fmt.Errorf(`expected HeavyBreakfast but got %T`, item)
	}
	return checkDateAt(at, step.at)
}

func checkHeavyLunch(item interface{}, at string) error {
	step, ok := item.(HeavyLunch)
	if !ok {
		return fmt.Errorf(`expected HeavyLunch but got %T`, item)
	}
	return checkDateAt(at, step.at)
}

func checkHeavyDinner(item interface{}, at string) error {
	step, ok := item.(HeavyDinner)
	if !ok {
		return fmt.Errorf(`expected HeavyDinner but got %T`, item)
	}
	return checkDateAt(at, step.at)
}

func checkSleep(item interface{}, start, end string) error {
	step, ok := item.(Sleep)
	if !ok {
		return fmt.Errorf(`expected Sleep but got %T`, item)
	}
	return checkDateRange(start, end, step.start, step.end)
}

func checkNoNap(item interface{}, start, end string) error {
	step, ok := item.(NoNap)
	if !ok {
		return fmt.Errorf(`expected NoNap but got %T`, item)
	}
	return checkDateRange(start, end, step.start, step.end)
}

func checkSetWatch(item interface{}, at string) error {
	step, ok := item.(SetWatch)
	if !ok {
		return fmt.Errorf(`expected SetWatch but got %T`, item)
	}
	return checkDateAt(at, step.at)
}

func checkArrive(item interface{}, at string) error {
	step, ok := item.(Arrive)
	if !ok {
		return fmt.Errorf(`expected Arrive but got %T`, item)
	}
	return checkDateAt(at, step.at)
}

func checkDateRange(expectedStart, expectedEnd string, actualStart, actualEnd time.Time) error {
	if !at(actualStart, expectedStart) || !at(actualEnd, expectedEnd) {
		return fmt.Errorf(`expected from %v to %v but got %v to %v`, expectedStart, expectedEnd, actualStart, actualEnd)
	}
	return nil
}

func checkDateAt(expectedAt string, actualAt time.Time) error {
	if !at(actualAt, expectedAt) {
		return fmt.Errorf(`expected %v but got %v`, expectedAt, actualAt)
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
