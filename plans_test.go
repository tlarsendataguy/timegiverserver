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

func TestEast34(t *testing.T) {
	plan := generatePlan(0, 4)
	if err := checkNoCaffeine(plan[0], `2017-01-11 06:00`, `2017-01-11 15:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan[1], `2017-01-11 15:00`, `2017-01-11 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoCaffeine(plan[2], `2017-01-11 16:30`, `2017-01-12 15:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan[3], `2017-01-12 15:00`, `2017-01-12 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoCaffeine(plan[4], `2017-01-12 16:30`, `2017-01-13 15:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan[5], `2017-01-13 15:00`, `2017-01-13 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoCaffeine(plan[6], `2017-01-13 16:30`, `2017-01-14 18:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeine3C(plan[7], `2017-01-14 18:00`, `2017-01-14 19:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoCaffeine(plan[8], `2017-01-14 19:00`, `2017-01-15 20:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyBreakfast(plan[9], `2017-01-13 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyLunch(plan[10], `2017-01-13 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyDinner(plan[11], `2017-01-13 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightBreakfast(plan[12], `2017-01-14 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightLunch(plan[13], `2017-01-14 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightDinner(plan[14], `2017-01-14 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyBreakfast(plan[15], `2017-01-15 03:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyLunch(plan[16], `2017-01-15 08:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyDinner(plan[17], `2017-01-15 13:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan[18], `2017-01-11 22:00`, `2017-01-12 06:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan[19], `2017-01-12 22:00`, `2017-01-13 06:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan[20], `2017-01-13 22:00`, `2017-01-14 06:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan[21], `2017-01-14 18:00`, `2017-01-15 02:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoNap(plan[22], `2017-01-15 02:00`, `2017-01-15 18:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan[23], `2017-01-15 18:00`, `2017-01-16 02:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSetWatch(plan[24], `2017-01-14 18:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkArrive(plan[25], `2017-01-15 08:00`); err != nil {
		t.Fatal(err.Error())
	}
}

func TestEast56(t *testing.T) {
	plan := generatePlan(0, 6)
	if err := checkNoCaffeine(plan[0], `2017-01-11 06:00`, `2017-01-11 15:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan[1], `2017-01-11 15:00`, `2017-01-11 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoCaffeine(plan[2], `2017-01-11 16:30`, `2017-01-12 15:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan[3], `2017-01-12 15:00`, `2017-01-12 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoCaffeine(plan[4], `2017-01-12 16:30`, `2017-01-13 15:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan[5], `2017-01-13 15:00`, `2017-01-13 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoCaffeine(plan[6], `2017-01-13 16:30`, `2017-01-14 18:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeine3C(plan[7], `2017-01-14 18:00`, `2017-01-14 19:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoCaffeine(plan[8], `2017-01-14 19:00`, `2017-01-15 18:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyBreakfast(plan[9], `2017-01-11 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyLunch(plan[10], `2017-01-11 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyDinner(plan[11], `2017-01-11 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightBreakfast(plan[12], `2017-01-12 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightLunch(plan[13], `2017-01-12 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightDinner(plan[14], `2017-01-12 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoSnack(plan[15], `2017-01-12 18:00`, `2017-01-12 22:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyBreakfast(plan[16], `2017-01-13 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyLunch(plan[17], `2017-01-13 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyDinner(plan[18], `2017-01-13 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoSnack(plan[19], `2017-01-13 18:00`, `2017-01-13 22:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightBreakfast(plan[20], `2017-01-14 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightLunch(plan[21], `2017-01-14 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightDinner(plan[22], `2017-01-14 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyBreakfast(plan[23], `2017-01-15 01:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyLunch(plan[24], `2017-01-15 06:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyDinner(plan[25], `2017-01-15 11:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoSnack(plan[26], `2017-01-15 12:00`, `2017-01-15 16:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan[27], `2017-01-11 22:00`, `2017-01-12 06:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan[28], `2017-01-12 22:00`, `2017-01-13 06:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan[29], `2017-01-13 22:00`, `2017-01-14 06:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan[30], `2017-01-14 18:00`, `2017-01-15 00:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoNap(plan[31], `2017-01-15 00:00`, `2017-01-15 16:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan[32], `2017-01-15 16:00`, `2017-01-16 00:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSetWatch(plan[33], `2017-01-14 18:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkArrive(plan[34], `2017-01-15 06:00`); err != nil {
		t.Fatal(err.Error())
	}
}

func TestEast78(t *testing.T) {
	plan := generatePlan(0, 8)
	if err := checkNoCaffeine(plan[0], `2017-01-11 06:00`, `2017-01-11 15:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan[1], `2017-01-11 15:00`, `2017-01-11 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoCaffeine(plan[2], `2017-01-11 16:30`, `2017-01-12 15:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan[3], `2017-01-12 15:00`, `2017-01-12 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoCaffeine(plan[4], `2017-01-12 16:30`, `2017-01-13 15:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan[5], `2017-01-13 15:00`, `2017-01-13 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoCaffeine(plan[6], `2017-01-13 16:30`, `2017-01-14 18:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeine2C(plan[7], `2017-01-14 18:00`, `2017-01-14 19:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoCaffeine(plan[8], `2017-01-14 19:00`, `2017-01-14 23:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeine2C(plan[9], `2017-01-14 23:00`, `2017-01-15 00:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoCaffeine(plan[10], `2017-01-15 00:00`, `2017-01-15 14:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyBreakfast(plan[11], `2017-01-11 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyLunch(plan[12], `2017-01-11 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyDinner(plan[13], `2017-01-11 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightBreakfast(plan[14], `2017-01-12 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightLunch(plan[15], `2017-01-12 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightDinner(plan[16], `2017-01-12 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoSnack(plan[17], `2017-01-12 18:00`, `2017-01-12 22:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyBreakfast(plan[18], `2017-01-13 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyLunch(plan[19], `2017-01-13 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyDinner(plan[20], `2017-01-13 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoSnack(plan[21], `2017-01-13 18:00`, `2017-01-13 22:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightBreakfast(plan[22], `2017-01-14 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightLunch(plan[23], `2017-01-14 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightDinnerOptional(plan[24], `2017-01-14 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyBreakfast(plan[25], `2017-01-14 23:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyLunch(plan[26], `2017-01-15 04:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyDinner(plan[27], `2017-01-15 09:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoSnack(plan[28], `2017-01-15 10:00`, `2017-01-15 14:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan[29], `2017-01-11 22:00`, `2017-01-12 06:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan[30], `2017-01-12 22:00`, `2017-01-13 06:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan[31], `2017-01-13 22:00`, `2017-01-14 06:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan[32], `2017-01-14 18:00`, `2017-01-14 22:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoNap(plan[33], `2017-01-14 22:00`, `2017-01-15 14:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan[34], `2017-01-15 14:00`, `2017-01-15 22:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSetWatch(plan[35], `2017-01-14 18:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkArrive(plan[36], `2017-01-15 04:00`); err != nil {
		t.Fatal(err.Error())
	}
}

func TestEast910(t *testing.T) {
	plan := generatePlan(0, 10)
	if err := checkNoCaffeine(plan[0], `2017-01-11 06:00`, `2017-01-11 15:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan[1], `2017-01-11 15:00`, `2017-01-11 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoCaffeine(plan[2], `2017-01-11 16:30`, `2017-01-12 15:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan[3], `2017-01-12 15:00`, `2017-01-12 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoCaffeine(plan[4], `2017-01-12 16:30`, `2017-01-13 15:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan[5], `2017-01-13 15:00`, `2017-01-13 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoCaffeine(plan[6], `2017-01-13 16:30`, `2017-01-14 21:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeine3C(plan[7], `2017-01-14 21:00`, `2017-01-14 22:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoCaffeine(plan[8], `2017-01-14 22:00`, `2017-01-15 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyBreakfast(plan[9], `2017-01-11 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyLunch(plan[10], `2017-01-11 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyDinner(plan[11], `2017-01-11 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightBreakfast(plan[12], `2017-01-12 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightLunch(plan[13], `2017-01-12 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightDinner(plan[14], `2017-01-12 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoSnack(plan[15], `2017-01-12 18:00`, `2017-01-12 22:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyBreakfast(plan[16], `2017-01-13 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyLunch(plan[17], `2017-01-13 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyDinner(plan[18], `2017-01-13 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoSnack(plan[19], `2017-01-13 18:00`, `2017-01-13 22:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightBreakfast(plan[20], `2017-01-14 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightLunch(plan[21], `2017-01-14 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyBreakfast(plan[22], `2017-01-14 21:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyLunch(plan[23], `2017-01-15 02:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyDinner(plan[24], `2017-01-15 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoSnack(plan[25], `2017-01-15 08:00`, `2017-01-15 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan[26], `2017-01-11 22:00`, `2017-01-12 06:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan[27], `2017-01-12 22:00`, `2017-01-13 06:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan[28], `2017-01-13 22:00`, `2017-01-14 06:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan[29], `2017-01-14 13:00`, `2017-01-14 20:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoNap(plan[30], `2017-01-14 20:00`, `2017-01-15 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan[31], `2017-01-15 12:00`, `2017-01-15 20:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSetWatch(plan[32], `2017-01-14 13:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkArrive(plan[33], `2017-01-15 02:00`); err != nil {
		t.Fatal(err.Error())
	}
}

func TestEast11(t *testing.T) {
	plan := generatePlan(0, 11)
	if err := checkNoCaffeine(plan[0], `2017-01-11 06:00`, `2017-01-11 15:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan[1], `2017-01-11 15:00`, `2017-01-11 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoCaffeine(plan[2], `2017-01-11 16:30`, `2017-01-12 15:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan[3], `2017-01-12 15:00`, `2017-01-12 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoCaffeine(plan[4], `2017-01-12 16:30`, `2017-01-13 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan[5], `2017-01-13 07:00`, `2017-01-13 11:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoCaffeine(plan[6], `2017-01-13 11:30`, `2017-01-14 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeine3C(plan[7], `2017-01-14 07:00`, `2017-01-14 11:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoCaffeine(plan[8], `2017-01-14 11:30`, `2017-01-15 11:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyBreakfast(plan[9], `2017-01-11 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyLunch(plan[10], `2017-01-11 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyDinner(plan[11], `2017-01-11 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightBreakfast(plan[12], `2017-01-12 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightLunch(plan[13], `2017-01-12 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightDinner(plan[14], `2017-01-12 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoSnack(plan[15], `2017-01-12 18:00`, `2017-01-12 22:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyBreakfast(plan[16], `2017-01-13 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyLunch(plan[17], `2017-01-13 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyDinner(plan[18], `2017-01-13 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoSnack(plan[19], `2017-01-13 18:00`, `2017-01-13 22:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightBreakfast(plan[20], `2017-01-14 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightLunch(plan[21], `2017-01-14 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyBreakfast(plan[22], `2017-01-14 20:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyLunch(plan[23], `2017-01-15 01:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyDinner(plan[24], `2017-01-15 06:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoSnack(plan[25], `2017-01-15 07:00`, `2017-01-15 11:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan[26], `2017-01-11 22:00`, `2017-01-12 06:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan[27], `2017-01-12 22:00`, `2017-01-13 06:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan[28], `2017-01-13 22:00`, `2017-01-14 06:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan[29], `2017-01-14 13:00`, `2017-01-14 19:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoNap(plan[30], `2017-01-14 19:00`, `2017-01-15 11:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan[31], `2017-01-15 11:00`, `2017-01-15 19:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSetWatch(plan[32], `2017-01-14 13:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkArrive(plan[33], `2017-01-15 01:00`); err != nil {
		t.Fatal(err.Error())
	}
}

func TestWest11(t *testing.T) {
	plan := generatePlan(0, -11)
	if err := checkNoCaffeine(plan[0], `2017-01-12 06:00`, `2017-01-12 15:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan[1], `2017-01-12 15:00`, `2017-01-12 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoCaffeine(plan[2], `2017-01-12 16:30`, `2017-01-13 15:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan[3], `2017-01-13 15:00`, `2017-01-13 16:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoCaffeine(plan[4], `2017-01-13 16:30`, `2017-01-14 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeineOk(plan[5], `2017-01-14 07:00`, `2017-01-14 11:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoCaffeine(plan[6], `2017-01-14 11:30`, `2017-01-15 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkCaffeine3C(plan[7], `2017-01-15 07:00`, `2017-01-15 11:30`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoCaffeine(plan[8], `2017-01-15 11:30`, `2017-01-16 09:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyBreakfast(plan[9], `2017-01-12 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyLunch(plan[10], `2017-01-12 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyDinner(plan[11], `2017-01-12 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightBreakfast(plan[12], `2017-01-13 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightLunch(plan[13], `2017-01-13 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightDinner(plan[14], `2017-01-13 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoSnack(plan[15], `2017-01-13 18:00`, `2017-01-13 22:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyBreakfast(plan[16], `2017-01-14 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyLunch(plan[17], `2017-01-14 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyDinner(plan[18], `2017-01-14 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoSnack(plan[19], `2017-01-14 18:00`, `2017-01-14 22:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightBreakfast(plan[20], `2017-01-15 07:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkLightLunch(plan[21], `2017-01-15 12:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyBreakfast(plan[22], `2017-01-15 18:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyLunch(plan[23], `2017-01-15 23:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkHeavyDinner(plan[24], `2017-01-16 04:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoSnack(plan[25], `2017-01-16 05:00`, `2017-01-16 09:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan[26], `2017-01-12 22:00`, `2017-01-13 06:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan[27], `2017-01-13 22:00`, `2017-01-14 06:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan[28], `2017-01-14 22:00`, `2017-01-15 06:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan[29], `2017-01-15 13:00`, `2017-01-15 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkNoNap(plan[30], `2017-01-15 17:00`, `2017-01-16 09:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSleep(plan[31], `2017-01-16 09:00`, `2017-01-16 17:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkSetWatch(plan[32], `2017-01-15 13:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err := checkArrive(plan[33], `2017-01-15 23:00`); err != nil {
		t.Fatal(err.Error())
	}
}

func TestWest12(t *testing.T) {
	plan := generatePlan(0, -2)
	err := checkNoCaffeine(plan[0], `2017-01-13 06:00`, `2017-01-13 15:00`)
	if err != nil {
		t.Fatal(err.Error())
	}
	err = checkCaffeineOk(plan[1], `2017-01-13 15:00`, `2017-01-13 16:30`)
	if err != nil {
		t.Fatal(err.Error())
	}
	err = checkNoCaffeine(plan[2], `2017-01-13 16:30`, `2017-01-14 07:00`)
	if err != nil {
		t.Fatal(err.Error())
	}
	err = checkCaffeineOk(plan[3], `2017-01-14 07:00`, `2017-01-14 11:00`)
	if err != nil {
		t.Fatal(err.Error())
	}
	err = checkNoCaffeine(plan[4], `2017-01-14 11:00`, `2017-01-15 06:00`)
	if err != nil {
		t.Fatal(err.Error())
	}
	err = checkCaffeine3C(plan[5], `2017-01-15 06:00`, `2017-01-15 11:00`)
	if err != nil {
		t.Fatal(err.Error())
	}
	err = checkNoCaffeine(plan[6], `2017-01-15 11:00`, `2017-01-16 00:00`)
	if err != nil {
		t.Fatal(err.Error())
	}
	err = checkLightBreakfast(plan[7], `2017-01-14 07:00`)
	if err != nil {
		t.Fatal(err.Error())
	}
	err = checkLightLunch(plan[8], `2017-01-14 12:00`)
	if err != nil {
		t.Fatal(err.Error())
	}
	err = checkLightDinner(plan[9], `2017-01-14 17:00`)
	if err != nil {
		t.Fatal(err.Error())
	}
	err = checkHeavyBreakfast(plan[10], `2017-01-15 09:00`)
	if err != nil {
		t.Fatal(err.Error())
	}
	err = checkHeavyLunch(plan[11], `2017-01-15 14:00`)
	if err != nil {
		t.Fatal(err.Error())
	}
	err = checkHeavyDinner(plan[12], `2017-01-15 19:00`)
	if err != nil {
		t.Fatal(err.Error())
	}
	err = checkSleep(plan[13], `2017-01-13 22:00`, `2017-01-14 06:00`)
	if err != nil {
		t.Fatal(err.Error())
	}
	err = checkSleep(plan[14], `2017-01-14 22:00`, `2017-01-15 06:00`)
	if err != nil {
		t.Fatal(err.Error())
	}
	err = checkSleep(plan[15], `2017-01-16 00:00`, `2017-01-16 08:00`)
	if err != nil {
		t.Fatal(err.Error())
	}
	if err = checkSetWatch(plan[16], `2017-01-15 09:00`); err != nil {
		t.Fatal(err.Error())
	}
	if err = checkArrive(plan[17], `2017-01-15 14:00`); err != nil {
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

func checkCaffeine2C(item interface{}, start, end string) error {
	step, ok := item.(Caffeine2C)
	if !ok {
		return fmt.Errorf(`expected Caffeine2C but got %T`, item)
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

func checkNoSnack(item interface{}, start, end string) error {
	step, ok := item.(NoSnack)
	if !ok {
		return fmt.Errorf(`expected NoSnack but got %T`, item)
	}
	return checkDateRange(start, end, step.start, step.end)
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
