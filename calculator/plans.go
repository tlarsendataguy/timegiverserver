package calculator

import (
	"time"
	"timegiverserver/calculator/steps"
)

const t0700 = time.Hour * 7
const t1000 = time.Hour * 10
const t1100 = time.Hour * 11
const t1130 = (time.Hour * 11) + (time.Minute * 30)
const t1500 = time.Hour * 15
const t1630 = (time.Hour * 16) + (time.Minute * 30)
const t1700 = time.Hour * 17
const t1800 = time.Hour * 18
const t1900 = time.Hour * 19
const t2200 = time.Hour * 22
const t2230 = (time.Hour * 22) + (time.Minute * 30)
const oneHour = time.Hour
const twoHours = time.Hour * 2
const threeHours = time.Hour * 3
const fiveHours = time.Hour * 5
const sevenHours = time.Hour * 7

func East12(c *Calculator) PlanSteps {
	return PlanSteps{
		Caffeine: []steps.Step{
			steps.NoCaffeine{At: c.departureLess2At(c.wake())},
			steps.CaffeineOk{Start: c.departureLess2At(t1500), End: c.departureLess2At(t1630)},
			steps.Caffeine3C{Start: c.departureLess1At(t1800), End: c.departureLess1At(t1900)},
		},
		Meals: []steps.Step{
			steps.LightBreakfast{At: c.departureLess1At(c.breakfast())},
			steps.LightLunch{At: c.departureLess1At(c.lunch())},
			steps.LightDinner{At: c.departureLess1At(c.dinner())},
			steps.HeavyBreakfast{At: c.arrivalAt(c.breakfast())},
			steps.HeavyLunch{At: c.arrivalAt(c.lunch())},
			steps.HeavyDinner{At: c.arrivalAt(c.dinner())},
		},
		Sleep: []steps.Step{
			steps.Sleep{Start: c.departureLess1At(c.sleep()), End: c.arrivalAt(c.wake())},
			steps.Sleep{Start: c.arrivalAt(c.sleep()), End: c.arrivalPlus1At(c.wake())},
		},
		Events: []steps.Step{
			steps.SetWatch{At: c.arrivalAt(c.breakfast())},
		},
	}
}

func East34(c *Calculator) PlanSteps {
	pre1Dinner := min(c.dinner(), t1800)
	pre1Sleep := max(pre1Dinner+oneHour, t1800)
	arrivalSleep := min(c.sleep(), t2230)

	return PlanSteps{
		Caffeine: []steps.Step{
			steps.NoCaffeine{At: c.departureLess4At(c.wake())},
			steps.CaffeineOk{Start: c.departureLess4At(t1500), End: c.departureLess4At(t1630)},
			steps.CaffeineOk{Start: c.departureLess3At(t1500), End: c.departureLess3At(t1630)},
			steps.CaffeineOk{Start: c.departureLess2At(t1500), End: c.departureLess2At(t1630)},
			steps.Caffeine3C{Start: c.departureLess1At(t1800), End: c.departureLess1At(t1900)},
		},
		Meals: []steps.Step{
			steps.HeavyBreakfast{At: c.departureLess2At(c.breakfast())},
			steps.HeavyLunch{At: c.departureLess2At(c.lunch())},
			steps.HeavyDinner{At: c.departureLess2At(c.dinner())},
			steps.LightBreakfast{At: c.departureLess1At(c.breakfast())},
			steps.LightLunch{At: c.departureLess1At(c.lunch())},
			steps.LightDinner{At: c.departureLess1At(pre1Dinner)},
			steps.HeavyBreakfast{At: c.arrivalAt(c.breakfast())},
			steps.HeavyLunch{At: c.arrivalAt(c.lunch())},
			steps.HeavyDinner{At: c.arrivalAt(c.dinner())},
		},
		Sleep: []steps.Step{
			steps.Sleep{Start: c.departureLess1At(pre1Sleep), End: c.arrivalAt(c.wake())},
			steps.Sleep{Start: c.arrivalAt(arrivalSleep), End: c.arrivalPlus1At(c.wake())},
		},
		Events: []steps.Step{
			steps.SetWatch{At: c.departureLess1At(t1800)},
		},
	}
}

func East56(c *Calculator) PlanSteps {
	pre1Dinner := min(c.dinner(), t1800)
	pre1Sleep := max(pre1Dinner+oneHour, t1800)
	arrivalSleep := min(c.sleep(), t2200)

	return PlanSteps{
		Caffeine: []steps.Step{
			steps.NoCaffeine{At: c.departureLess4At(c.wake())},
			steps.CaffeineOk{Start: c.departureLess4At(t1500), End: c.departureLess4At(t1630)},
			steps.CaffeineOk{Start: c.departureLess3At(t1500), End: c.departureLess3At(t1630)},
			steps.CaffeineOk{Start: c.departureLess2At(t1500), End: c.departureLess2At(t1630)},
			steps.Caffeine3C{Start: c.departureLess1At(t1800), End: c.departureLess1At(t1900)},
		},
		Meals: []steps.Step{
			steps.HeavyBreakfast{At: c.departureLess4At(c.breakfast())},
			steps.HeavyLunch{At: c.departureLess4At(c.lunch())},
			steps.HeavyDinner{At: c.departureLess4At(c.dinner())},
			steps.LightBreakfast{At: c.departureLess3At(c.breakfast())},
			steps.LightLunch{At: c.departureLess3At(c.lunch())},
			steps.LightDinner{At: c.departureLess3At(c.dinner())},
			steps.HeavyBreakfast{At: c.departureLess2At(c.breakfast())},
			steps.HeavyLunch{At: c.departureLess2At(c.lunch())},
			steps.HeavyDinner{At: c.departureLess2At(c.dinner())},
			steps.LightBreakfast{At: c.departureLess1At(c.breakfast())},
			steps.LightLunch{At: c.departureLess1At(c.lunch())},
			steps.LightDinner{At: c.departureLess1At(pre1Dinner)},
			steps.HeavyBreakfast{At: c.arrivalAt(c.breakfast())},
			steps.HeavyLunch{At: c.arrivalAt(c.lunch())},
			steps.HeavyDinner{At: c.arrivalAt(c.dinner())},
		},
		Sleep: []steps.Step{
			steps.Sleep{Start: c.departureLess1At(pre1Sleep), End: c.arrivalAt(c.wake())},
			steps.Sleep{Start: c.arrivalAt(arrivalSleep), End: c.arrivalPlus1At(c.wake())},
		},
		Events: []steps.Step{
			steps.SetWatch{At: c.departureLess1At(t1800)},
		},
	}
}

func East78(c *Calculator) PlanSteps {
	pre1Dinner := min(c.dinner(), t1700)
	pre1Sleep := max(pre1Dinner+oneHour, t1700)
	arrivalSleep := min(c.sleep(), t2200)

	return PlanSteps{
		Caffeine: []steps.Step{
			steps.NoCaffeine{At: c.departureLess4At(c.wake())},
			steps.CaffeineOk{Start: c.departureLess4At(t1500), End: c.departureLess4At(t1630)},
			steps.CaffeineOk{Start: c.departureLess3At(t1500), End: c.departureLess3At(t1630)},
			steps.CaffeineOk{Start: c.departureLess2At(t1500), End: c.departureLess2At(t1630)},
			steps.Caffeine2C{Start: c.departureLess1At(t1800), End: c.departureLess1At(t1900)},
			steps.Caffeine2C{Start: c.arrivalAt(c.breakfast()), End: c.arrivalAt(c.breakfast() + oneHour)},
		},
		Meals: []steps.Step{
			steps.HeavyBreakfast{At: c.departureLess4At(c.breakfast())},
			steps.HeavyLunch{At: c.departureLess4At(c.lunch())},
			steps.HeavyDinner{At: c.departureLess4At(c.dinner())},
			steps.LightBreakfast{At: c.departureLess3At(c.breakfast())},
			steps.LightLunch{At: c.departureLess3At(c.lunch())},
			steps.LightDinner{At: c.departureLess3At(c.dinner())},
			steps.HeavyBreakfast{At: c.departureLess2At(c.breakfast())},
			steps.HeavyLunch{At: c.departureLess2At(c.lunch())},
			steps.HeavyDinner{At: c.departureLess2At(c.dinner())},
			steps.LightBreakfast{At: c.departureLess1At(c.breakfast())},
			steps.LightLunch{At: c.departureLess1At(c.lunch())},
			steps.LightDinnerOptional{At: c.departureLess1At(pre1Dinner)},
			steps.HeavyBreakfast{At: c.arrivalAt(c.breakfast())},
			steps.HeavyLunch{At: c.arrivalAt(c.lunch())},
			steps.HeavyDinner{At: c.arrivalAt(c.dinner())},
		},
		Sleep: []steps.Step{
			steps.Sleep{Start: c.departureLess1At(pre1Sleep), End: c.arrivalAt(c.wake())},
			steps.Sleep{Start: c.arrivalAt(arrivalSleep), End: c.arrivalPlus1At(c.wake())},
		},
		Events: []steps.Step{
			steps.SetWatch{At: c.departureLess1At(t1800)},
		},
	}
}

func East910(c *Calculator) PlanSteps {
	arrivalSleep := min(c.sleep(), t2200)
	arrivalDinner := min(c.dinner(), arrivalSleep-oneHour)

	return PlanSteps{
		Caffeine: []steps.Step{
			steps.NoCaffeine{At: c.departureLess4At(c.wake())},
			steps.CaffeineOk{Start: c.departureLess4At(t1500), End: c.departureLess4At(t1630)},
			steps.CaffeineOk{Start: c.departureLess3At(t1500), End: c.departureLess3At(t1630)},
			steps.CaffeineOk{Start: c.departureLess2At(t1500), End: c.departureLess2At(t1630)},
			steps.Caffeine3C{Start: c.arrivalAt(c.breakfast()), End: c.arrivalAt(c.breakfast() + oneHour)},
		},
		Meals: []steps.Step{
			steps.HeavyBreakfast{At: c.departureLess4At(c.breakfast())},
			steps.HeavyLunch{At: c.departureLess4At(c.lunch())},
			steps.HeavyDinner{At: c.departureLess4At(c.dinner())},
			steps.LightBreakfast{At: c.departureLess3At(c.breakfast())},
			steps.LightLunch{At: c.departureLess3At(c.lunch())},
			steps.LightDinner{At: c.departureLess3At(c.dinner())},
			steps.HeavyBreakfast{At: c.departureLess2At(c.breakfast())},
			steps.HeavyLunch{At: c.departureLess2At(c.lunch())},
			steps.HeavyDinner{At: c.departureLess2At(c.dinner())},
			steps.LightBreakfast{At: c.departureLess1At(c.breakfast())},
			steps.LightLunch{At: c.departureLess1At(c.lunch())},
			steps.HeavyBreakfast{At: c.arrivalAt(c.breakfast())},
			steps.HeavyLunch{At: c.arrivalAt(c.lunch())},
			steps.HeavyDinner{At: c.arrivalAt(arrivalDinner)},
		},
		Sleep: []steps.Step{
			steps.Sleep{Start: c.departureLess1At(c.lunch() + oneHour), End: c.arrivalAt(c.wake())},
			steps.Sleep{Start: c.arrivalAt(c.sleep()), End: c.arrivalPlus1At(c.wake())},
		},
		Events: []steps.Step{
			steps.SetWatch{At: c.departureLess1At(c.lunch() + oneHour)},
		},
	}
}

func West12(c *Calculator) PlanSteps {
	return PlanSteps{
		Caffeine: []steps.Step{
			steps.NoCaffeine{At: c.departureLess2At(c.wake())},
			steps.CaffeineOk{Start: c.departureLess2At(t1500), End: c.departureLess2At(t1630)},
			steps.CaffeineOk{Start: c.departureLess1At(t0700), End: c.departureLess1At(t1100)},
			steps.Caffeine3C{Start: c.departureAt(c.wake()), End: c.departureAt(t1100)},
		},
		Meals: []steps.Step{
			steps.LightBreakfast{At: c.departureLess1At(c.breakfast())},
			steps.LightLunch{At: c.departureLess1At(c.lunch())},
			steps.LightDinner{At: c.departureLess1At(c.dinner())},
			steps.HeavyBreakfast{At: c.arrivalAt(c.breakfast())},
			steps.HeavyLunch{At: c.arrivalAt(c.lunch())},
			steps.HeavyDinner{At: c.arrivalAt(c.dinner())},
		},
		Sleep: []steps.Step{
			steps.Sleep{Start: c.arrivalAt(c.sleep()), End: c.arrivalPlus1At(c.wake())},
		},
		Events: []steps.Step{
			steps.SetWatch{At: c.arrivalAt(c.breakfast())},
		},
	}
}

func West34(c *Calculator) PlanSteps {
	return PlanSteps{
		Caffeine: []steps.Step{
			steps.NoCaffeine{At: c.departureLess3At(c.wake())},
			steps.CaffeineOk{Start: c.departureLess3At(t1500), End: c.departureLess3At(t1630)},
			steps.CaffeineOk{Start: c.departureLess2At(t1500), End: c.departureLess2At(t1630)},
			steps.CaffeineOk{Start: c.departureLess1At(t0700), End: c.departureLess1At(t1100)},
			steps.Caffeine3C{Start: c.departureAt(c.wake()), End: c.departureAt(t1100)},
		},
		Meals: []steps.Step{
			steps.LightBreakfast{At: c.departureLess1At(c.breakfast())},
			steps.LightLunch{At: c.departureLess1At(c.lunch())},
			steps.LightDinner{At: c.departureLess1At(c.dinner())},
			steps.HeavyBreakfast{At: c.arrivalAt(c.breakfast())},
			steps.HeavyLunch{At: c.arrivalAt(c.lunch())},
			steps.HeavyDinner{At: c.arrivalAt(c.dinner())},
		},
		Sleep: []steps.Step{
			steps.Sleep{Start: c.arrivalAt(c.sleep()), End: c.arrivalPlus1At(c.wake())},
		},
		Events: []steps.Step{
			steps.SetWatch{At: c.arrivalAt(c.breakfast())},
		},
	}
}

func West56(c *Calculator) PlanSteps {
	departureLess1Wakeup := min(c.wake()+twoHours, t1000)

	return PlanSteps{
		Caffeine: []steps.Step{
			steps.NoCaffeine{At: c.departureLess4At(c.wake())},
			steps.CaffeineOk{Start: c.departureLess4At(t1500), End: c.departureLess4At(t1630)},
			steps.CaffeineOk{Start: c.departureLess3At(t1500), End: c.departureLess3At(t1630)},
			steps.CaffeineOk{Start: c.departureLess2At(t0700), End: c.departureLess2At(t1130)},
			steps.Caffeine3C{Start: c.departureLess1At(departureLess1Wakeup), End: c.departureLess1At(t1100)},
		},
		Meals: []steps.Step{
			steps.HeavyBreakfast{At: c.departureLess4At(c.breakfast())},
			steps.HeavyLunch{At: c.departureLess4At(c.lunch())},
			steps.HeavyDinner{At: c.departureLess4At(c.dinner())},
			steps.LightBreakfast{At: c.departureLess3At(c.breakfast())},
			steps.LightLunch{At: c.departureLess3At(c.lunch())},
			steps.LightDinner{At: c.departureLess3At(c.dinner())},
			steps.HeavyBreakfast{At: c.departureLess2At(c.breakfast())},
			steps.HeavyLunch{At: c.departureLess2At(c.lunch())},
			steps.HeavyDinner{At: c.departureLess2At(c.dinner())},
			steps.LightBreakfast{At: c.departureLess1At(c.breakfast() + twoHours)},
			steps.LightLunch{At: c.departureLess1At(c.lunch() + twoHours)},
			steps.LightDinner{At: c.departureLess1At(c.dinner() + twoHours)},
			steps.HeavyBreakfast{At: c.arrivalAt(c.breakfast())},
			steps.HeavyLunch{At: c.arrivalAt(c.lunch())},
			steps.HeavyDinner{At: c.arrivalAt(c.dinner())},
		},
		Sleep: []steps.Step{
			steps.Sleep{Start: c.departureLess2At(c.sleep()), End: c.departureLess1At(departureLess1Wakeup)},
			steps.Sleep{Start: c.arrivalLess1At(c.sleep()), End: c.arrivalAt(c.wake())},
			steps.Sleep{Start: c.arrivalAt(c.sleep()), End: c.arrivalPlus1At(c.wake())},
		},
		Events: []steps.Step{
			steps.SetWatch{At: c.departureLess1At(c.dinner() + threeHours)},
		},
	}
}

func West78(c *Calculator) PlanSteps {
	departureWake := c.wake() + threeHours
	departureBreakfast := min(c.breakfast()+threeHours, c.wake()+fiveHours)

	return PlanSteps{
		Caffeine: []steps.Step{
			steps.NoCaffeine{At: c.departureLess3At(c.wake())},
			steps.CaffeineOk{Start: c.departureLess3At(t1500), End: c.departureLess3At(t1630)},
			steps.CaffeineOk{Start: c.departureLess2At(t1500), End: c.departureLess2At(t1630)},
			steps.CaffeineOk{Start: c.departureLess1At(t1500), End: c.departureLess1At(t1630)},
			steps.Caffeine3C{Start: c.departureAt(departureWake), End: c.departureAt(t1130)},
		},
		Meals: []steps.Step{
			steps.HeavyBreakfast{At: c.departureLess3At(c.breakfast())},
			steps.HeavyLunch{At: c.departureLess3At(c.lunch())},
			steps.HeavyDinner{At: c.departureLess3At(c.dinner())},
			steps.LightBreakfast{At: c.departureLess2At(c.breakfast())},
			steps.LightLunch{At: c.departureLess2At(c.lunch())},
			steps.LightDinner{At: c.departureLess2At(c.dinner())},
			steps.HeavyBreakfast{At: c.departureLess1At(c.breakfast())},
			steps.HeavyLunch{At: c.departureLess1At(c.lunch())},
			steps.HeavyDinner{At: c.departureLess1At(c.dinner())},
			steps.LightBreakfast{At: c.departureAt(departureBreakfast)},
			steps.HeavyBreakfast{At: c.arrivalAt(c.breakfast())},
			steps.HeavyLunch{At: c.arrivalAt(c.lunch())},
			steps.HeavyDinner{At: c.arrivalAt(c.dinner())},
		},
		Sleep: []steps.Step{
			steps.Sleep{Start: c.departureLess1At(c.sleep()), End: c.departureAt(departureWake)},
			steps.Sleep{Start: c.departureAt(departureBreakfast + oneHour), End: c.arrivalAt(c.breakfast())},
			steps.Sleep{Start: c.arrivalAt(c.sleep()), End: c.arrivalPlus1At(c.wake())},
		},
		Events: []steps.Step{
			steps.SetWatch{At: c.departureAt(departureBreakfast + oneHour)},
		},
	}
}

func West910(c *Calculator) PlanSteps {
	departureWake := c.wake() + threeHours
	departureBreakfast := c.breakfast() + threeHours
	departureLunch := min(c.lunch()+threeHours, c.breakfast()+sevenHours)

	return PlanSteps{
		Caffeine: []steps.Step{
			steps.NoCaffeine{At: c.departureLess3At(c.wake())},
			steps.CaffeineOk{Start: c.departureLess3At(t1500), End: c.departureLess3At(t1630)},
			steps.CaffeineOk{Start: c.departureLess2At(t1500), End: c.departureLess2At(t1630)},
			steps.CaffeineOk{Start: c.departureLess1At(t1500), End: c.departureLess1At(t1630)},
			steps.Caffeine3C{Start: c.departureAt(departureWake), End: c.departureAt(departureWake + oneHour)},
		},
		Meals: []steps.Step{
			steps.HeavyBreakfast{At: c.departureLess3At(c.breakfast())},
			steps.HeavyLunch{At: c.departureLess3At(c.lunch())},
			steps.HeavyDinner{At: c.departureLess3At(c.dinner())},
			steps.LightBreakfast{At: c.departureLess2At(c.breakfast())},
			steps.LightLunch{At: c.departureLess2At(c.lunch())},
			steps.LightDinner{At: c.departureLess2At(c.dinner())},
			steps.HeavyBreakfast{At: c.departureLess1At(c.breakfast())},
			steps.HeavyLunch{At: c.departureLess1At(c.lunch())},
			steps.HeavyDinner{At: c.departureLess1At(c.dinner())},
			steps.LightBreakfast{At: c.departureAt(departureBreakfast)},
			steps.LightLunch{At: c.departureAt(departureLunch)},
			steps.HeavyBreakfast{At: c.arrivalAt(c.breakfast())},
			steps.HeavyLunch{At: c.arrivalAt(c.lunch())},
			steps.HeavyDinner{At: c.arrivalAt(c.dinner())},
		},
		Sleep: []steps.Step{
			steps.Sleep{Start: c.departureLess1At(c.sleep()), End: c.departureAt(departureWake)},
			steps.Sleep{Start: c.departureAt(departureLunch + oneHour), End: c.arrivalAt(c.breakfast())},
			steps.Sleep{Start: c.arrivalAt(c.sleep()), End: c.arrivalPlus1At(c.wake())},
		},
		Events: []steps.Step{
			steps.SetWatch{At: c.departureAt(departureLunch + oneHour)},
		},
	}
}

func Both1112(c *Calculator) PlanSteps {
	arrivalSleep := min(c.sleep(), t2200)

	return PlanSteps{
		Caffeine: []steps.Step{
			steps.NoCaffeine{At: c.departureLess4At(c.wake())},
			steps.CaffeineOk{Start: c.departureLess4At(t1500), End: c.departureLess4At(t1630)},
			steps.CaffeineOk{Start: c.departureLess3At(t1500), End: c.departureLess3At(t1630)},
			steps.CaffeineOk{Start: c.departureLess2At(t0700), End: c.departureLess2At(t1130)},
			steps.Caffeine3C{Start: c.departureLess1At(t0700), End: c.departureLess1At(t1130)},
		},
		Meals: []steps.Step{
			steps.HeavyBreakfast{At: c.departureLess4At(c.breakfast())},
			steps.HeavyLunch{At: c.departureLess4At(c.lunch())},
			steps.HeavyDinner{At: c.departureLess4At(c.dinner())},
			steps.LightBreakfast{At: c.departureLess3At(c.breakfast())},
			steps.LightLunch{At: c.departureLess3At(c.lunch())},
			steps.LightDinner{At: c.departureLess3At(c.dinner())},
			steps.HeavyBreakfast{At: c.departureLess2At(c.breakfast())},
			steps.HeavyLunch{At: c.departureLess2At(c.lunch())},
			steps.HeavyDinner{At: c.departureLess2At(c.dinner())},
			steps.LightBreakfast{At: c.departureLess1At(c.breakfast())},
			steps.LightLunch{At: c.departureLess1At(c.lunch())},
			steps.HeavyBreakfast{At: c.arrivalAt(c.breakfast())},
			steps.HeavyLunch{At: c.arrivalAt(c.lunch())},
			steps.HeavyDinner{At: c.arrivalAt(c.dinner())},
		},
		Sleep: []steps.Step{
			steps.Sleep{Start: c.departureLess1At(c.lunch() + oneHour), End: c.arrivalAt(c.wake())},
			steps.Sleep{Start: c.arrivalAt(arrivalSleep), End: c.arrivalPlus1At(c.wake())},
		},
		Events: []steps.Step{
			steps.SetWatch{At: c.departureLess1At(c.lunch() + oneHour)},
		},
	}
}

func min(left, right time.Duration) time.Duration {
	if left < right {
		return left
	}
	return right
}

func max(left, right time.Duration) time.Duration {
	if left > right {
		return left
	}
	return right
}
