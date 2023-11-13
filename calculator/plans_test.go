package calculator

import (
	"fmt"
	"os"
	"path"
	"testing"
	"time"
	"timegiverserver/calculator/steps"
)

var arrival = time.Date(2017, 1, 15, 12, 0, 0, 0, time.UTC)
var routine = DailyRoutine{Wake: time.Hour * 6, Breakfast: time.Hour * 7, Lunch: time.Hour * 12, Dinner: time.Hour * 17, Sleep: time.Hour * 22}

func TestEast12(t *testing.T) {
	plan := generatePlan(0, 2)
	if err := checkNoCaffeine(plan.Caffeine[0], `2017-01-13 06:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan.Caffeine[1], `2017-01-13 15:00`, `2017-01-13 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeine3C(plan.Caffeine[2], `2017-01-14 18:00`, `2017-01-14 19:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightBreakfast(plan.Meals[0], `2017-01-14 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightLunch(plan.Meals[1], `2017-01-14 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightDinner(plan.Meals[2], `2017-01-14 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyBreakfast(plan.Meals[3], `2017-01-15 05:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyLunch(plan.Meals[4], `2017-01-15 10:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyDinner(plan.Meals[5], `2017-01-15 15:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan.Sleep[0], `2017-01-14 22:00`, `2017-01-15 04:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan.Sleep[1], `2017-01-15 20:00`, `2017-01-16 04:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSetWatch(plan.Events[0], `2017-01-15 05:00`); err != nil {
		t.Fatal(err.Error())
	}
}

func TestEast34(t *testing.T) {
	plan := generatePlan(0, 4)
	if err := checkNoCaffeine(plan.Caffeine[0], `2017-01-11 06:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan.Caffeine[1], `2017-01-11 15:00`, `2017-01-11 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan.Caffeine[2], `2017-01-12 15:00`, `2017-01-12 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan.Caffeine[3], `2017-01-13 15:00`, `2017-01-13 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeine3C(plan.Caffeine[4], `2017-01-14 18:00`, `2017-01-14 19:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyBreakfast(plan.Meals[0], `2017-01-13 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyLunch(plan.Meals[1], `2017-01-13 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyDinner(plan.Meals[2], `2017-01-13 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightBreakfast(plan.Meals[3], `2017-01-14 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightLunch(plan.Meals[4], `2017-01-14 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightDinner(plan.Meals[5], `2017-01-14 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyBreakfast(plan.Meals[6], `2017-01-15 03:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyLunch(plan.Meals[7], `2017-01-15 08:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyDinner(plan.Meals[8], `2017-01-15 13:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan.Sleep[0], `2017-01-14 18:00`, `2017-01-15 02:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan.Sleep[1], `2017-01-15 18:00`, `2017-01-16 02:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSetWatch(plan.Events[0], `2017-01-14 18:00`); err != nil {
		t.Fatal(err.Error())
	}
}

func TestEast56(t *testing.T) {
	plan := generatePlan(0, 6)
	if err := checkNoCaffeine(plan.Caffeine[0], `2017-01-11 06:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan.Caffeine[1], `2017-01-11 15:00`, `2017-01-11 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan.Caffeine[2], `2017-01-12 15:00`, `2017-01-12 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan.Caffeine[3], `2017-01-13 15:00`, `2017-01-13 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeine3C(plan.Caffeine[4], `2017-01-14 18:00`, `2017-01-14 19:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyBreakfast(plan.Meals[0], `2017-01-11 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyLunch(plan.Meals[1], `2017-01-11 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyDinner(plan.Meals[2], `2017-01-11 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightBreakfast(plan.Meals[3], `2017-01-12 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightLunch(plan.Meals[4], `2017-01-12 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightDinner(plan.Meals[5], `2017-01-12 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyBreakfast(plan.Meals[6], `2017-01-13 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyLunch(plan.Meals[7], `2017-01-13 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyDinner(plan.Meals[8], `2017-01-13 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightBreakfast(plan.Meals[9], `2017-01-14 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightLunch(plan.Meals[10], `2017-01-14 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightDinner(plan.Meals[11], `2017-01-14 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyBreakfast(plan.Meals[12], `2017-01-15 01:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyLunch(plan.Meals[13], `2017-01-15 06:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyDinner(plan.Meals[14], `2017-01-15 11:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan.Sleep[0], `2017-01-14 18:00`, `2017-01-15 00:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan.Sleep[1], `2017-01-15 16:00`, `2017-01-16 00:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSetWatch(plan.Events[0], `2017-01-14 18:00`); err != nil {
		t.Fatal(err.Error())
	}
}

func TestEast78(t *testing.T) {
	plan := generatePlan(0, 8)
	if err := checkNoCaffeine(plan.Caffeine[0], `2017-01-11 06:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan.Caffeine[1], `2017-01-11 15:00`, `2017-01-11 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan.Caffeine[2], `2017-01-12 15:00`, `2017-01-12 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan.Caffeine[3], `2017-01-13 15:00`, `2017-01-13 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeine2C(plan.Caffeine[4], `2017-01-14 18:00`, `2017-01-14 19:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeine2C(plan.Caffeine[5], `2017-01-14 23:00`, `2017-01-15 00:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyBreakfast(plan.Meals[0], `2017-01-11 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyLunch(plan.Meals[1], `2017-01-11 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyDinner(plan.Meals[2], `2017-01-11 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightBreakfast(plan.Meals[3], `2017-01-12 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightLunch(plan.Meals[4], `2017-01-12 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightDinner(plan.Meals[5], `2017-01-12 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyBreakfast(plan.Meals[6], `2017-01-13 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyLunch(plan.Meals[7], `2017-01-13 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyDinner(plan.Meals[8], `2017-01-13 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightBreakfast(plan.Meals[9], `2017-01-14 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightLunch(plan.Meals[10], `2017-01-14 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightDinnerOptional(plan.Meals[11], `2017-01-14 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyBreakfast(plan.Meals[12], `2017-01-14 23:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyLunch(plan.Meals[13], `2017-01-15 04:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyDinner(plan.Meals[14], `2017-01-15 09:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan.Sleep[0], `2017-01-14 18:00`, `2017-01-14 22:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan.Sleep[1], `2017-01-15 14:00`, `2017-01-15 22:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSetWatch(plan.Events[0], `2017-01-14 18:00`); err != nil {
		t.Fatal(err.Error())
	}
}

func TestEast910(t *testing.T) {
	plan := generatePlan(0, 10)
	if err := checkNoCaffeine(plan.Caffeine[0], `2017-01-11 06:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan.Caffeine[1], `2017-01-11 15:00`, `2017-01-11 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan.Caffeine[2], `2017-01-12 15:00`, `2017-01-12 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan.Caffeine[3], `2017-01-13 15:00`, `2017-01-13 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeine3C(plan.Caffeine[4], `2017-01-14 21:00`, `2017-01-14 22:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyBreakfast(plan.Meals[0], `2017-01-11 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyLunch(plan.Meals[1], `2017-01-11 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyDinner(plan.Meals[2], `2017-01-11 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightBreakfast(plan.Meals[3], `2017-01-12 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightLunch(plan.Meals[4], `2017-01-12 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightDinner(plan.Meals[5], `2017-01-12 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyBreakfast(plan.Meals[6], `2017-01-13 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyLunch(plan.Meals[7], `2017-01-13 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyDinner(plan.Meals[8], `2017-01-13 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightBreakfast(plan.Meals[9], `2017-01-14 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightLunch(plan.Meals[10], `2017-01-14 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyBreakfast(plan.Meals[11], `2017-01-14 21:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyLunch(plan.Meals[12], `2017-01-15 02:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyDinner(plan.Meals[13], `2017-01-15 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan.Sleep[0], `2017-01-14 13:00`, `2017-01-14 20:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan.Sleep[1], `2017-01-15 12:00`, `2017-01-15 20:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSetWatch(plan.Events[0], `2017-01-14 13:00`); err != nil {
		t.Fatal(err.Error())
	}
}

func TestEast11(t *testing.T) {
	plan := generatePlan(0, 11)
	if err := checkNoCaffeine(plan.Caffeine[0], `2017-01-11 06:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan.Caffeine[1], `2017-01-11 15:00`, `2017-01-11 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan.Caffeine[2], `2017-01-12 15:00`, `2017-01-12 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan.Caffeine[3], `2017-01-13 07:00`, `2017-01-13 11:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeine3C(plan.Caffeine[4], `2017-01-14 07:00`, `2017-01-14 11:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyBreakfast(plan.Meals[0], `2017-01-11 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyLunch(plan.Meals[1], `2017-01-11 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyDinner(plan.Meals[2], `2017-01-11 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightBreakfast(plan.Meals[3], `2017-01-12 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightLunch(plan.Meals[4], `2017-01-12 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightDinner(plan.Meals[5], `2017-01-12 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyBreakfast(plan.Meals[6], `2017-01-13 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyLunch(plan.Meals[7], `2017-01-13 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyDinner(plan.Meals[8], `2017-01-13 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightBreakfast(plan.Meals[9], `2017-01-14 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightLunch(plan.Meals[10], `2017-01-14 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyBreakfast(plan.Meals[11], `2017-01-14 20:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyLunch(plan.Meals[12], `2017-01-15 01:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyDinner(plan.Meals[13], `2017-01-15 06:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan.Sleep[0], `2017-01-14 13:00`, `2017-01-14 19:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan.Sleep[1], `2017-01-15 11:00`, `2017-01-15 19:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSetWatch(plan.Events[0], `2017-01-14 13:00`); err != nil {
		t.Fatal(err.Error())
	}

}

func TestWest11(t *testing.T) {
	plan := generatePlan(0, -11)
	if err := checkNoCaffeine(plan.Caffeine[0], `2017-01-12 06:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan.Caffeine[1], `2017-01-12 15:00`, `2017-01-12 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan.Caffeine[2], `2017-01-13 15:00`, `2017-01-13 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan.Caffeine[3], `2017-01-14 07:00`, `2017-01-14 11:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeine3C(plan.Caffeine[4], `2017-01-15 07:00`, `2017-01-15 11:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyBreakfast(plan.Meals[0], `2017-01-12 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyLunch(plan.Meals[1], `2017-01-12 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyDinner(plan.Meals[2], `2017-01-12 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightBreakfast(plan.Meals[3], `2017-01-13 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightLunch(plan.Meals[4], `2017-01-13 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightDinner(plan.Meals[5], `2017-01-13 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyBreakfast(plan.Meals[6], `2017-01-14 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyLunch(plan.Meals[7], `2017-01-14 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyDinner(plan.Meals[8], `2017-01-14 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightBreakfast(plan.Meals[9], `2017-01-15 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightLunch(plan.Meals[10], `2017-01-15 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyBreakfast(plan.Meals[11], `2017-01-15 18:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyLunch(plan.Meals[12], `2017-01-15 23:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyDinner(plan.Meals[13], `2017-01-16 04:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan.Sleep[0], `2017-01-15 13:00`, `2017-01-15 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan.Sleep[1], `2017-01-16 09:00`, `2017-01-16 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSetWatch(plan.Events[0], `2017-01-15 13:00`); err != nil {
		t.Fatal(err.Error())
	}
}

func TestWest12(t *testing.T) {
	plan := generatePlan(0, -2)
	err := checkNoCaffeine(plan.Caffeine[0], `2017-01-13 06:00`)
	if err != nil {
		t.Fatal(err.Error())
	}
	err = checkCaffeineOk(plan.Caffeine[1], `2017-01-13 15:00`, `2017-01-13 16:30`)
	if err != nil {
		t.Fatal(err.Error())
	}
	err = checkCaffeineOk(plan.Caffeine[2], `2017-01-14 07:00`, `2017-01-14 11:00`)
	if err != nil {
		t.Fatal(err.Error())
	}
	err = checkCaffeine3C(plan.Caffeine[3], `2017-01-15 06:00`, `2017-01-15 11:00`)
	if err != nil {
		t.Fatal(err.Error())
	}
	err = checkLightBreakfast(plan.Meals[0], `2017-01-14 07:00`)
	if err != nil {
		t.Fatal(err.Error())
	}
	err = checkLightLunch(plan.Meals[1], `2017-01-14 12:00`)
	if err != nil {
		t.Fatal(err.Error())
	}
	err = checkLightDinner(plan.Meals[2], `2017-01-14 17:00`)
	if err != nil {
		t.Fatal(err.Error())
	}
	err = checkHeavyBreakfast(plan.Meals[3], `2017-01-15 09:00`)
	if err != nil {
		t.Fatal(err.Error())
	}
	err = checkHeavyLunch(plan.Meals[4], `2017-01-15 14:00`)
	if err != nil {
		t.Fatal(err.Error())
	}
	err = checkHeavyDinner(plan.Meals[5], `2017-01-15 19:00`)
	if err != nil {
		t.Fatal(err.Error())
	}
	err = checkSleep(plan.Sleep[0], `2017-01-16 00:00`, `2017-01-16 08:00`)
	if err != nil {
		t.Fatal(err.Error())
	}
	if err = checkSetWatch(plan.Events[0], `2017-01-15 09:00`); err != nil {
		t.Fatal(err.Error())
	}
}

func TestWest34(t *testing.T) {
	plan := generatePlan(0, -4)
	if err := checkNoCaffeine(plan.Caffeine[0], `2017-01-12 06:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan.Caffeine[1], `2017-01-12 15:00`, `2017-01-12 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan.Caffeine[2], `2017-01-13 15:00`, `2017-01-13 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan.Caffeine[3], `2017-01-14 07:00`, `2017-01-14 11:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeine3C(plan.Caffeine[4], `2017-01-15 06:00`, `2017-01-15 11:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightBreakfast(plan.Meals[0], `2017-01-14 07:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkLightLunch(plan.Meals[1], `2017-01-14 12:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkLightDinner(plan.Meals[2], `2017-01-14 17:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkHeavyBreakfast(plan.Meals[3], `2017-01-15 11:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkHeavyLunch(plan.Meals[4], `2017-01-15 16:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkHeavyDinner(plan.Meals[5], `2017-01-15 21:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkSleep(plan.Sleep[0], `2017-01-16 02:00`, `2017-01-16 10:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkSetWatch(plan.Events[0], `2017-01-15 11:00`); err != nil {
		t.Fatal(err.Error())
	}
}

func TestWest56(t *testing.T) {
	plan := generatePlan(0, -6)
	if err := checkNoCaffeine(plan.Caffeine[0], `2017-01-11 06:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan.Caffeine[1], `2017-01-11 15:00`, `2017-01-11 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan.Caffeine[2], `2017-01-12 15:00`, `2017-01-12 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan.Caffeine[3], `2017-01-13 07:00`, `2017-01-13 11:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeine3C(plan.Caffeine[4], `2017-01-14 08:00`, `2017-01-14 11:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyBreakfast(plan.Meals[0], `2017-01-11 07:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkHeavyLunch(plan.Meals[1], `2017-01-11 12:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkHeavyDinner(plan.Meals[2], `2017-01-11 17:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkLightBreakfast(plan.Meals[3], `2017-01-12 07:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkLightLunch(plan.Meals[4], `2017-01-12 12:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkLightDinner(plan.Meals[5], `2017-01-12 17:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkHeavyBreakfast(plan.Meals[6], `2017-01-13 07:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkHeavyLunch(plan.Meals[7], `2017-01-13 12:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkHeavyDinner(plan.Meals[8], `2017-01-13 17:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkLightBreakfast(plan.Meals[9], `2017-01-14 09:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkLightLunch(plan.Meals[10], `2017-01-14 14:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkLightDinner(plan.Meals[11], `2017-01-14 19:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkHeavyBreakfast(plan.Meals[12], `2017-01-15 13:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkHeavyLunch(plan.Meals[13], `2017-01-15 18:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkHeavyDinner(plan.Meals[14], `2017-01-15 23:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkSleep(plan.Sleep[0], `2017-01-13 22:00`, `2017-01-14 08:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkSleep(plan.Sleep[1], `2017-01-15 04:00`, `2017-01-15 12:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkSleep(plan.Sleep[2], `2017-01-16 04:00`, `2017-01-16 12:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkSetWatch(plan.Events[0], `2017-01-14 20:00`); err != nil {
		t.Fatal(err.Error())
	}
}

func TestWest78(t *testing.T) {
	plan := generatePlan(0, -8)
	if err := checkNoCaffeine(plan.Caffeine[0], `2017-01-12 06:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan.Caffeine[1], `2017-01-12 15:00`, `2017-01-12 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan.Caffeine[2], `2017-01-13 15:00`, `2017-01-13 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan.Caffeine[3], `2017-01-14 15:00`, `2017-01-14 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeine3C(plan.Caffeine[4], `2017-01-15 09:00`, `2017-01-15 11:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyBreakfast(plan.Meals[0], `2017-01-12 07:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkHeavyLunch(plan.Meals[1], `2017-01-12 12:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkHeavyDinner(plan.Meals[2], `2017-01-12 17:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkLightBreakfast(plan.Meals[3], `2017-01-13 07:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkLightLunch(plan.Meals[4], `2017-01-13 12:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkLightDinner(plan.Meals[5], `2017-01-13 17:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkHeavyBreakfast(plan.Meals[6], `2017-01-14 07:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkHeavyLunch(plan.Meals[7], `2017-01-14 12:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkHeavyDinner(plan.Meals[8], `2017-01-14 17:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkLightBreakfast(plan.Meals[9], `2017-01-15 10:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkHeavyBreakfast(plan.Meals[10], `2017-01-15 15:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkHeavyLunch(plan.Meals[11], `2017-01-15 20:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkHeavyDinner(plan.Meals[12], `2017-01-16 01:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkSleep(plan.Sleep[0], `2017-01-14 22:00`, `2017-01-15 09:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkSleep(plan.Sleep[1], `2017-01-15 11:00`, `2017-01-15 15:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkSleep(plan.Sleep[2], `2017-01-16 06:00`, `2017-01-16 14:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkSetWatch(plan.Events[0], `2017-01-15 11:00`); err != nil {
		t.Fatal(err.Error())
	}
}

func TestWest910(t *testing.T) {
	plan := generatePlan(0, -10)
	if err := checkNoCaffeine(plan.Caffeine[0], `2017-01-12 06:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan.Caffeine[1], `2017-01-12 15:00`, `2017-01-12 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan.Caffeine[2], `2017-01-13 15:00`, `2017-01-13 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan.Caffeine[3], `2017-01-14 15:00`, `2017-01-14 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeine3C(plan.Caffeine[4], `2017-01-15 09:00`, `2017-01-15 10:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyBreakfast(plan.Meals[0], `2017-01-12 07:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkHeavyLunch(plan.Meals[1], `2017-01-12 12:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkHeavyDinner(plan.Meals[2], `2017-01-12 17:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkLightBreakfast(plan.Meals[3], `2017-01-13 07:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkLightLunch(plan.Meals[4], `2017-01-13 12:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkLightDinner(plan.Meals[5], `2017-01-13 17:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkHeavyBreakfast(plan.Meals[6], `2017-01-14 07:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkHeavyLunch(plan.Meals[7], `2017-01-14 12:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkHeavyDinner(plan.Meals[8], `2017-01-14 17:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkLightBreakfast(plan.Meals[9], `2017-01-15 10:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkLightLunch(plan.Meals[10], `2017-01-15 14:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkHeavyBreakfast(plan.Meals[11], `2017-01-15 17:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkHeavyLunch(plan.Meals[12], `2017-01-15 22:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkHeavyDinner(plan.Meals[13], `2017-01-16 03:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkSleep(plan.Sleep[0], `2017-01-14 22:00`, `2017-01-15 09:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkSleep(plan.Sleep[1], `2017-01-15 15:00`, `2017-01-15 17:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkSleep(plan.Sleep[2], `2017-01-16 08:00`, `2017-01-16 16:00`); err != nil {
		t.Fatalf(err.Error())
	}
	if err := checkSetWatch(plan.Events[0], `2017-01-15 15:00`); err != nil {
		t.Fatal(err.Error())
	}
}

func TestGenerateBatchIcs(t *testing.T) {
	folder := `./testing_batches`
	i := 1.0
	eastStart := -7.0
	westStart := 11.0
	for i < 14 {
		eastEnd := eastStart + i
		westEnd := westStart - i

		eastPlan := generatePlan(eastStart, eastEnd)
		westPlan := generatePlan(westStart, westEnd)

		eastIcs := BuildIcsFiles(eastPlan)
		westIcs := BuildIcsFiles(westPlan)

		err := os.WriteFile(path.Join(folder, fmt.Sprintf(`East %v Caffeine.ics`, i)), []byte(eastIcs.Caffeine), 0666)
		if err != nil {
			t.Fatalf(`expected no error but got: %v`, err.Error())
		}
		err = os.WriteFile(path.Join(folder, fmt.Sprintf(`East %v Meals.ics`, i)), []byte(eastIcs.Meals), 0666)
		if err != nil {
			t.Fatalf(`expected no error but got: %v`, err.Error())
		}
		err = os.WriteFile(path.Join(folder, fmt.Sprintf(`East %v Sleep.ics`, i)), []byte(eastIcs.Sleep), 0666)
		if err != nil {
			t.Fatalf(`expected no error but got: %v`, err.Error())
		}
		err = os.WriteFile(path.Join(folder, fmt.Sprintf(`East %v Events.ics`, i)), []byte(eastIcs.Events), 0666)
		if err != nil {
			t.Fatalf(`expected no error but got: %v`, err.Error())
		}
		err = os.WriteFile(path.Join(folder, fmt.Sprintf(`West %v Caffeine.ics`, i)), []byte(westIcs.Caffeine), 0666)
		if err != nil {
			t.Fatalf(`expected no error but got: %v`, err.Error())
		}
		err = os.WriteFile(path.Join(folder, fmt.Sprintf(`West %v Meals.ics`, i)), []byte(westIcs.Meals), 0666)
		if err != nil {
			t.Fatalf(`expected no error but got: %v`, err.Error())
		}
		err = os.WriteFile(path.Join(folder, fmt.Sprintf(`West %v Sleep.ics`, i)), []byte(westIcs.Sleep), 0666)
		if err != nil {
			t.Fatalf(`expected no error but got: %v`, err.Error())
		}
		err = os.WriteFile(path.Join(folder, fmt.Sprintf(`West %v Events.ics`, i)), []byte(westIcs.Events), 0666)
		if err != nil {
			t.Fatalf(`expected no error but got: %v`, err.Error())
		}
		i++
	}
}

func checkNoCaffeine(item interface{}, at string) error {
	step, ok := item.(steps.NoCaffeine)
	if !ok {
		return fmt.Errorf(`expected NoCaffeine but got %T`, item)
	}
	return checkDateAt(at, step.At)
}

func checkCaffeineOk(item interface{}, start, end string) error {
	step, ok := item.(steps.CaffeineOk)
	if !ok {
		return fmt.Errorf(`expected CaffeineOk but got %T`, item)
	}
	return checkDateRange(start, end, step.Start, step.End)
}

func checkCaffeine2C(item interface{}, start, end string) error {
	step, ok := item.(steps.Caffeine2C)
	if !ok {
		return fmt.Errorf(`expected Caffeine2C but got %T`, item)
	}
	return checkDateRange(start, end, step.Start, step.End)
}

func checkCaffeine3C(item interface{}, start, end string) error {
	step, ok := item.(steps.Caffeine3C)
	if !ok {
		return fmt.Errorf(`expected Caffeine3C but got %T`, item)
	}
	return checkDateRange(start, end, step.Start, step.End)
}

func checkLightBreakfast(item interface{}, at string) error {
	step, ok := item.(steps.LightBreakfast)
	if !ok {
		return fmt.Errorf(`expected LightBreakfast but got %T`, item)
	}
	return checkDateAt(at, step.At)
}

func checkLightLunch(item interface{}, at string) error {
	step, ok := item.(steps.LightLunch)
	if !ok {
		return fmt.Errorf(`expected LightLunch but got %T`, item)
	}
	return checkDateAt(at, step.At)
}

func checkLightDinner(item interface{}, at string) error {
	step, ok := item.(steps.LightDinner)
	if !ok {
		return fmt.Errorf(`expected LightDinner but got %T`, item)
	}
	return checkDateAt(at, step.At)
}

func checkLightDinnerOptional(item interface{}, at string) error {
	step, ok := item.(steps.LightDinnerOptional)
	if !ok {
		return fmt.Errorf(`expected LightDinnerOptionall but got %T`, item)
	}
	return checkDateAt(at, step.At)
}

func checkHeavyBreakfast(item interface{}, at string) error {
	step, ok := item.(steps.HeavyBreakfast)
	if !ok {
		return fmt.Errorf(`expected HeavyBreakfast but got %T`, item)
	}
	return checkDateAt(at, step.At)
}

func checkHeavyLunch(item interface{}, at string) error {
	step, ok := item.(steps.HeavyLunch)
	if !ok {
		return fmt.Errorf(`expected HeavyLunch but got %T`, item)
	}
	return checkDateAt(at, step.At)
}

func checkHeavyDinner(item interface{}, at string) error {
	step, ok := item.(steps.HeavyDinner)
	if !ok {
		return fmt.Errorf(`expected HeavyDinner but got %T`, item)
	}
	return checkDateAt(at, step.At)
}

func checkSleep(item interface{}, start, end string) error {
	step, ok := item.(steps.Sleep)
	if !ok {
		return fmt.Errorf(`expected Sleep but got %T`, item)
	}
	return checkDateRange(start, end, step.Start, step.End)
}

func checkSetWatch(item interface{}, at string) error {
	step, ok := item.(steps.SetWatch)
	if !ok {
		return fmt.Errorf(`expected SetWatch but got %T`, item)
	}
	return checkDateAt(at, step.At)
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

func generatePlan(departureOffset, arrivalOffset float64) PlanSteps {
	calc := InitializeCalculator(Inputs{
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
