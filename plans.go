package timegiverserver

import (
	"time"
	"timegiverserver/steps"
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
const t2400 = time.Hour * 24
const oneHour = time.Hour
const twoHours = time.Hour * 2
const threeHours = time.Hour * 3
const fiveHours = time.Hour * 5
const sevenHours = time.Hour * 7

func East12(c *Calculator) []Step {
	return []Step{
		//Caffeine
		steps.NoCaffeine{Start: c.departureLess2At(c.wake()), End: c.departureLess2At(t1500)},
		steps.CaffeineOk{Start: c.departureLess2At(t1500), End: c.departureLess2At(t1630)},
		steps.NoCaffeine{Start: c.departureLess2At(t1630), End: c.departureLess1At(t1800)},
		steps.Caffeine3C{Start: c.departureLess1At(t1800), End: c.departureLess1At(t1900)},
		steps.NoCaffeine{Start: c.departureLess1At(t1900), End: c.arrivalAt(t2400)},

		//Meals
		steps.LightBreakfast{At: c.departureLess1At(c.breakfast())},
		steps.LightLunch{At: c.departureLess1At(c.lunch())},
		steps.LightDinner{At: c.departureLess1At(c.dinner())},
		steps.HeavyBreakfast{At: c.arrivalAt(c.breakfast())},
		steps.HeavyLunch{At: c.arrivalAt(c.lunch())},
		steps.HeavyDinner{At: c.arrivalAt(c.dinner())},

		//Sleep
		steps.Sleep{Start: c.departureLess2At(c.sleep()), End: c.departureLess1At(c.wake())},
		steps.Sleep{Start: c.departureLess1At(c.sleep()), End: c.arrivalAt(c.wake())},
		steps.Sleep{Start: c.arrivalAt(c.sleep()), End: c.arrivalPlus1At(c.wake())},

		//Events
		steps.SetWatch{At: c.arrivalAt(c.breakfast())},
		c.arrivalStep(),
	}
}

func East34(c *Calculator) []Step {
	pre1Dinner := min(c.dinner(), t1800)
	pre1Sleep := max(pre1Dinner+oneHour, t1800)
	arrivalSleep := min(c.sleep(), t2230)

	return []Step{
		//Caffeine
		steps.NoCaffeine{Start: c.departureLess4At(c.wake()), End: c.departureLess4At(t1500)},
		steps.CaffeineOk{Start: c.departureLess4At(t1500), End: c.departureLess4At(t1630)},
		steps.NoCaffeine{Start: c.departureLess4At(t1630), End: c.departureLess3At(t1500)},
		steps.CaffeineOk{Start: c.departureLess3At(t1500), End: c.departureLess3At(t1630)},
		steps.NoCaffeine{Start: c.departureLess3At(t1630), End: c.departureLess2At(t1500)},
		steps.CaffeineOk{Start: c.departureLess2At(t1500), End: c.departureLess2At(t1630)},
		steps.NoCaffeine{Start: c.departureLess2At(t1630), End: c.departureLess1At(t1800)},
		steps.Caffeine3C{Start: c.departureLess1At(t1800), End: c.departureLess1At(t1900)},
		steps.NoCaffeine{Start: c.departureLess1At(t1900), End: c.arrivalAt(t2400)},

		//Meals
		steps.HeavyBreakfast{At: c.departureLess2At(c.breakfast())},
		steps.HeavyLunch{At: c.departureLess2At(c.lunch())},
		steps.HeavyDinner{At: c.departureLess2At(c.dinner())},
		steps.LightBreakfast{At: c.departureLess1At(c.breakfast())},
		steps.LightLunch{At: c.departureLess1At(c.lunch())},
		steps.LightDinner{At: c.departureLess1At(pre1Dinner)},
		steps.HeavyBreakfast{At: c.arrivalAt(c.breakfast())},
		steps.HeavyLunch{At: c.arrivalAt(c.lunch())},
		steps.HeavyDinner{At: c.arrivalAt(c.dinner())},

		//Sleep
		steps.Sleep{Start: c.departureLess4At(c.sleep()), End: c.departureLess3At(c.wake())},
		steps.Sleep{Start: c.departureLess3At(c.sleep()), End: c.departureLess2At(c.wake())},
		steps.Sleep{Start: c.departureLess2At(c.sleep()), End: c.departureLess1At(c.wake())},
		steps.Sleep{Start: c.departureLess1At(pre1Sleep), End: c.arrivalAt(c.wake())},
		steps.NoNap{Start: c.arrivalAt(c.wake()), End: c.arrivalAt(c.sleep())},
		steps.Sleep{Start: c.arrivalAt(arrivalSleep), End: c.arrivalPlus1At(c.wake())},

		//Events
		steps.SetWatch{At: c.departureLess1At(t1800)},
		c.arrivalStep(),
	}
}

func East56(c *Calculator) []Step {
	pre1Dinner := min(c.dinner(), t1800)
	pre1Sleep := max(pre1Dinner+oneHour, t1800)
	arrivalSleep := min(c.sleep(), t2200)

	return []Step{
		//Caffeine
		steps.NoCaffeine{Start: c.departureLess4At(c.wake()), End: c.departureLess4At(t1500)},
		steps.CaffeineOk{Start: c.departureLess4At(t1500), End: c.departureLess4At(t1630)},
		steps.NoCaffeine{Start: c.departureLess4At(t1630), End: c.departureLess3At(t1500)},
		steps.CaffeineOk{Start: c.departureLess3At(t1500), End: c.departureLess3At(t1630)},
		steps.NoCaffeine{Start: c.departureLess3At(t1630), End: c.departureLess2At(t1500)},
		steps.CaffeineOk{Start: c.departureLess2At(t1500), End: c.departureLess2At(t1630)},
		steps.NoCaffeine{Start: c.departureLess2At(t1630), End: c.departureLess1At(t1800)},
		steps.Caffeine3C{Start: c.departureLess1At(t1800), End: c.departureLess1At(t1900)},
		steps.NoCaffeine{Start: c.departureLess1At(t1900), End: c.arrivalAt(t2400)},

		//Meals
		steps.HeavyBreakfast{At: c.departureLess4At(c.breakfast())},
		steps.HeavyLunch{At: c.departureLess4At(c.lunch())},
		steps.HeavyDinner{At: c.departureLess4At(c.dinner())},
		steps.LightBreakfast{At: c.departureLess3At(c.breakfast())},
		steps.LightLunch{At: c.departureLess3At(c.lunch())},
		steps.LightDinner{At: c.departureLess3At(c.dinner())},
		steps.NoSnack{Start: c.departureLess3At(c.dinner() + oneHour), End: c.departureLess3At(c.sleep())},
		steps.HeavyBreakfast{At: c.departureLess2At(c.breakfast())},
		steps.HeavyLunch{At: c.departureLess2At(c.lunch())},
		steps.HeavyDinner{At: c.departureLess2At(c.dinner())},
		steps.NoSnack{Start: c.departureLess2At(c.dinner() + oneHour), End: c.departureLess2At(c.sleep())},
		steps.LightBreakfast{At: c.departureLess1At(c.breakfast())},
		steps.LightLunch{At: c.departureLess1At(c.lunch())},
		steps.LightDinner{At: c.departureLess1At(pre1Dinner)},
		steps.HeavyBreakfast{At: c.arrivalAt(c.breakfast())},
		steps.HeavyLunch{At: c.arrivalAt(c.lunch())},
		steps.HeavyDinner{At: c.arrivalAt(c.dinner())},
		steps.NoSnack{Start: c.arrivalAt(c.dinner() + oneHour), End: c.arrivalAt(c.sleep())},

		//Sleep
		steps.Sleep{Start: c.departureLess4At(c.sleep()), End: c.departureLess3At(c.wake())},
		steps.Sleep{Start: c.departureLess3At(c.sleep()), End: c.departureLess2At(c.wake())},
		steps.Sleep{Start: c.departureLess2At(c.sleep()), End: c.departureLess1At(c.wake())},
		steps.Sleep{Start: c.departureLess1At(pre1Sleep), End: c.arrivalAt(c.wake())},
		steps.NoNap{Start: c.arrivalAt(c.wake()), End: c.arrivalAt(arrivalSleep)},
		steps.Sleep{Start: c.arrivalAt(arrivalSleep), End: c.arrivalPlus1At(c.wake())},

		//Events
		steps.SetWatch{At: c.departureLess1At(t1800)},
		c.arrivalStep(),
	}
}

func East78(c *Calculator) []Step {
	pre1Dinner := min(c.dinner(), t1700)
	pre1Sleep := max(pre1Dinner+oneHour, t1700)
	arrivalSleep := min(c.sleep(), t2200)

	return []Step{
		//Caffeine
		steps.NoCaffeine{Start: c.departureLess4At(c.wake()), End: c.departureLess4At(t1500)},
		steps.CaffeineOk{Start: c.departureLess4At(t1500), End: c.departureLess4At(t1630)},
		steps.NoCaffeine{Start: c.departureLess4At(t1630), End: c.departureLess3At(t1500)},
		steps.CaffeineOk{Start: c.departureLess3At(t1500), End: c.departureLess3At(t1630)},
		steps.NoCaffeine{Start: c.departureLess3At(t1630), End: c.departureLess2At(t1500)},
		steps.CaffeineOk{Start: c.departureLess2At(t1500), End: c.departureLess2At(t1630)},
		steps.NoCaffeine{Start: c.departureLess2At(t1630), End: c.departureLess1At(t1800)},
		steps.Caffeine2C{Start: c.departureLess1At(t1800), End: c.departureLess1At(t1900)},
		steps.NoCaffeine{Start: c.departureLess1At(t1900), End: c.arrivalAt(c.breakfast())},
		steps.Caffeine2C{Start: c.arrivalAt(c.breakfast()), End: c.arrivalAt(c.breakfast() + oneHour)},
		steps.NoCaffeine{Start: c.arrivalAt(c.breakfast() + oneHour), End: c.arrivalAt(c.sleep())},

		//Meals
		steps.HeavyBreakfast{At: c.departureLess4At(c.breakfast())},
		steps.HeavyLunch{At: c.departureLess4At(c.lunch())},
		steps.HeavyDinner{At: c.departureLess4At(c.dinner())},
		steps.LightBreakfast{At: c.departureLess3At(c.breakfast())},
		steps.LightLunch{At: c.departureLess3At(c.lunch())},
		steps.LightDinner{At: c.departureLess3At(c.dinner())},
		steps.NoSnack{Start: c.departureLess3At(c.dinner() + oneHour), End: c.departureLess3At(c.sleep())},
		steps.HeavyBreakfast{At: c.departureLess2At(c.breakfast())},
		steps.HeavyLunch{At: c.departureLess2At(c.lunch())},
		steps.HeavyDinner{At: c.departureLess2At(c.dinner())},
		steps.NoSnack{Start: c.departureLess2At(c.dinner() + oneHour), End: c.departureLess2At(c.sleep())},
		steps.LightBreakfast{At: c.departureLess1At(c.breakfast())},
		steps.LightLunch{At: c.departureLess1At(c.lunch())},
		steps.LightDinnerOptional{At: c.departureLess1At(pre1Dinner)},
		steps.HeavyBreakfast{At: c.arrivalAt(c.breakfast())},
		steps.HeavyLunch{At: c.arrivalAt(c.lunch())},
		steps.HeavyDinner{At: c.arrivalAt(c.dinner())},
		steps.NoSnack{Start: c.arrivalAt(c.dinner() + oneHour), End: c.arrivalAt(c.sleep())},

		//Sleep
		steps.Sleep{Start: c.departureLess4At(c.sleep()), End: c.departureLess3At(c.wake())},
		steps.Sleep{Start: c.departureLess3At(c.sleep()), End: c.departureLess2At(c.wake())},
		steps.Sleep{Start: c.departureLess2At(c.sleep()), End: c.departureLess1At(c.wake())},
		steps.Sleep{Start: c.departureLess1At(pre1Sleep), End: c.arrivalAt(c.wake())},
		steps.NoNap{Start: c.arrivalAt(c.wake()), End: c.arrivalAt(arrivalSleep)},
		steps.Sleep{Start: c.arrivalAt(arrivalSleep), End: c.arrivalPlus1At(c.wake())},

		//Events
		steps.SetWatch{At: c.departureLess1At(t1800)},
		c.arrivalStep(),
	}
}

func East910(c *Calculator) []Step {
	arrivalSleep := min(c.sleep(), t2200)
	arrivalDinner := min(c.dinner(), arrivalSleep-oneHour)

	return []Step{
		//Caffeine
		steps.NoCaffeine{Start: c.departureLess4At(c.wake()), End: c.departureLess4At(t1500)},
		steps.CaffeineOk{Start: c.departureLess4At(t1500), End: c.departureLess4At(t1630)},
		steps.NoCaffeine{Start: c.departureLess4At(t1630), End: c.departureLess3At(t1500)},
		steps.CaffeineOk{Start: c.departureLess3At(t1500), End: c.departureLess3At(t1630)},
		steps.NoCaffeine{Start: c.departureLess3At(t1630), End: c.departureLess2At(t1500)},
		steps.CaffeineOk{Start: c.departureLess2At(t1500), End: c.departureLess2At(t1630)},
		steps.NoCaffeine{Start: c.departureLess2At(t1630), End: c.arrivalAt(c.breakfast())},
		steps.Caffeine3C{Start: c.arrivalAt(c.breakfast()), End: c.arrivalAt(c.breakfast() + oneHour)},
		steps.NoCaffeine{Start: c.arrivalAt(c.breakfast() + oneHour), End: c.arrivalAt(c.sleep())},

		//Meals
		steps.HeavyBreakfast{At: c.departureLess4At(c.breakfast())},
		steps.HeavyLunch{At: c.departureLess4At(c.lunch())},
		steps.HeavyDinner{At: c.departureLess4At(c.dinner())},
		steps.LightBreakfast{At: c.departureLess3At(c.breakfast())},
		steps.LightLunch{At: c.departureLess3At(c.lunch())},
		steps.LightDinner{At: c.departureLess3At(c.dinner())},
		steps.NoSnack{Start: c.departureLess3At(c.dinner() + oneHour), End: c.departureLess3At(c.sleep())},
		steps.HeavyBreakfast{At: c.departureLess2At(c.breakfast())},
		steps.HeavyLunch{At: c.departureLess2At(c.lunch())},
		steps.HeavyDinner{At: c.departureLess2At(c.dinner())},
		steps.NoSnack{Start: c.departureLess2At(c.dinner() + oneHour), End: c.departureLess2At(c.sleep())},
		steps.LightBreakfast{At: c.departureLess1At(c.breakfast())},
		steps.LightLunch{At: c.departureLess1At(c.lunch())},
		steps.HeavyBreakfast{At: c.arrivalAt(c.breakfast())},
		steps.HeavyLunch{At: c.arrivalAt(c.lunch())},
		steps.HeavyDinner{At: c.arrivalAt(arrivalDinner)},
		steps.NoSnack{Start: c.arrivalAt(arrivalDinner + oneHour), End: c.arrivalAt(c.sleep())},

		//Sleep
		steps.Sleep{Start: c.departureLess4At(c.sleep()), End: c.departureLess3At(c.wake())},
		steps.Sleep{Start: c.departureLess3At(c.sleep()), End: c.departureLess2At(c.wake())},
		steps.Sleep{Start: c.departureLess2At(c.sleep()), End: c.departureLess1At(c.wake())},
		steps.Sleep{Start: c.departureLess1At(c.lunch() + oneHour), End: c.arrivalAt(c.wake())},
		steps.NoNap{Start: c.arrivalAt(c.wake()), End: c.arrivalAt(c.sleep())},
		steps.Sleep{Start: c.arrivalAt(c.sleep()), End: c.arrivalPlus1At(c.wake())},

		//Events
		steps.SetWatch{At: c.departureLess1At(c.lunch() + oneHour)},
		c.arrivalStep(),
	}
}

func West12(c *Calculator) []Step {
	return []Step{
		//Caffeine
		steps.NoCaffeine{Start: c.departureLess2At(c.wake()), End: c.departureLess2At(t1500)},
		steps.CaffeineOk{Start: c.departureLess2At(t1500), End: c.departureLess2At(t1630)},
		steps.NoCaffeine{Start: c.departureLess2At(t1630), End: c.departureLess1At(t0700)},
		steps.CaffeineOk{Start: c.departureLess1At(t0700), End: c.departureLess1At(t1100)},
		steps.NoCaffeine{Start: c.departureLess1At(t1100), End: c.departureAt(c.wake())},
		steps.Caffeine3C{Start: c.departureAt(c.wake()), End: c.departureAt(t1100)},
		steps.NoCaffeine{Start: c.departureAt(t1100), End: c.arrivalAt(c.sleep())},

		//Meals
		steps.LightBreakfast{At: c.departureLess1At(c.breakfast())},
		steps.LightLunch{At: c.departureLess1At(c.lunch())},
		steps.LightDinner{At: c.departureLess1At(c.dinner())},
		steps.HeavyBreakfast{At: c.arrivalAt(c.breakfast())},
		steps.HeavyLunch{At: c.arrivalAt(c.lunch())},
		steps.HeavyDinner{At: c.arrivalAt(c.dinner())},

		//Sleep
		steps.Sleep{Start: c.departureLess2At(c.sleep()), End: c.departureLess1At(c.wake())},
		steps.Sleep{Start: c.departureLess1At(c.sleep()), End: c.departureAt(c.wake())},
		steps.Sleep{Start: c.arrivalAt(c.sleep()), End: c.arrivalPlus1At(c.wake())},

		//Events
		steps.SetWatch{At: c.arrivalAt(c.breakfast())},
		c.arrivalStep(),
	}
}

func West34(c *Calculator) []Step {
	return []Step{
		//Caffeine
		steps.NoCaffeine{Start: c.departureLess3At(c.wake()), End: c.departureLess3At(t1500)},
		steps.CaffeineOk{Start: c.departureLess3At(t1500), End: c.departureLess3At(t1630)},
		steps.NoCaffeine{Start: c.departureLess3At(t1630), End: c.departureLess2At(t1500)},
		steps.CaffeineOk{Start: c.departureLess2At(t1500), End: c.departureLess2At(t1630)},
		steps.NoCaffeine{Start: c.departureLess2At(t1630), End: c.departureLess1At(t0700)},
		steps.CaffeineOk{Start: c.departureLess1At(t0700), End: c.departureLess1At(t1100)},
		steps.NoCaffeine{Start: c.departureLess1At(t1100), End: c.departureAt(c.wake())},
		steps.Caffeine3C{Start: c.departureAt(c.wake()), End: c.departureAt(t1100)},
		steps.NoCaffeine{Start: c.departureAt(t1100), End: c.arrivalAt(c.sleep())},

		//Meals
		steps.LightBreakfast{At: c.departureLess1At(c.breakfast())},
		steps.LightLunch{At: c.departureLess1At(c.lunch())},
		steps.LightDinner{At: c.departureLess1At(c.dinner())},
		steps.HeavyBreakfast{At: c.arrivalAt(c.breakfast())},
		steps.HeavyLunch{At: c.arrivalAt(c.lunch())},
		steps.HeavyDinner{At: c.arrivalAt(c.dinner())},

		//Sleep
		steps.Sleep{Start: c.departureLess3At(c.sleep()), End: c.departureLess2At(c.wake())},
		steps.Sleep{Start: c.departureLess2At(c.sleep()), End: c.departureLess1At(c.wake())},
		steps.Sleep{Start: c.departureLess1At(c.sleep()), End: c.departureAt(c.wake())},
		steps.Sleep{Start: c.arrivalAt(c.sleep()), End: c.arrivalPlus1At(c.wake())},

		//Events
		steps.SetWatch{At: c.arrivalAt(c.breakfast())},
		c.arrivalStep(),
	}
}

func West56(c *Calculator) []Step {
	departureLess1Wakeup := min(c.wake()+twoHours, t1000)

	return []Step{
		//Caffeine
		steps.NoCaffeine{Start: c.departureLess4At(c.wake()), End: c.departureLess4At(t1500)},
		steps.CaffeineOk{Start: c.departureLess4At(t1500), End: c.departureLess4At(t1630)},
		steps.NoCaffeine{Start: c.departureLess4At(t1630), End: c.departureLess3At(t1500)},
		steps.CaffeineOk{Start: c.departureLess3At(t1500), End: c.departureLess3At(t1630)},
		steps.NoCaffeine{Start: c.departureLess3At(t1630), End: c.departureLess2At(t0700)},
		steps.CaffeineOk{Start: c.departureLess2At(t0700), End: c.departureLess2At(t1130)},
		steps.NoCaffeine{Start: c.departureLess2At(t1130), End: c.departureLess1At(departureLess1Wakeup)},
		steps.Caffeine3C{Start: c.departureLess1At(departureLess1Wakeup), End: c.departureLess1At(t1100)},
		steps.NoCaffeine{Start: c.departureLess1At(t1100), End: c.arrivalAt(c.sleep())},

		//Meals
		steps.HeavyBreakfast{At: c.departureLess4At(c.breakfast())},
		steps.HeavyLunch{At: c.departureLess4At(c.lunch())},
		steps.HeavyDinner{At: c.departureLess4At(c.dinner())},
		steps.LightBreakfast{At: c.departureLess3At(c.breakfast())},
		steps.LightLunch{At: c.departureLess3At(c.lunch())},
		steps.LightDinner{At: c.departureLess3At(c.dinner())},
		steps.NoSnack{Start: c.departureLess3At(c.dinner() + oneHour), End: c.departureLess3At(c.sleep())},
		steps.HeavyBreakfast{At: c.departureLess2At(c.breakfast())},
		steps.HeavyLunch{At: c.departureLess2At(c.lunch())},
		steps.HeavyDinner{At: c.departureLess2At(c.dinner())},
		steps.LightBreakfast{At: c.departureLess1At(c.breakfast() + twoHours)},
		steps.LightLunch{At: c.departureLess1At(c.lunch() + twoHours)},
		steps.LightDinner{At: c.departureLess1At(c.dinner() + twoHours)},
		steps.HeavyBreakfast{At: c.arrivalAt(c.breakfast())},
		steps.HeavyLunch{At: c.arrivalAt(c.lunch())},
		steps.HeavyDinner{At: c.arrivalAt(c.dinner())},

		//Sleep
		steps.Sleep{Start: c.departureLess4At(c.sleep()), End: c.departureLess3At(c.wake())},
		steps.Sleep{Start: c.departureLess3At(c.sleep()), End: c.departureLess2At(c.wake())},
		steps.Sleep{Start: c.departureLess2At(c.sleep()), End: c.departureLess1At(departureLess1Wakeup)},
		steps.Sleep{Start: c.arrivalLess1At(c.sleep()), End: c.arrivalAt(c.wake())},
		steps.Sleep{Start: c.arrivalAt(c.sleep()), End: c.arrivalPlus1At(c.wake())},

		//Events
		steps.SetWatch{At: c.departureLess1At(c.dinner() + threeHours)},
		c.arrivalStep(),
	}
}

func West78(c *Calculator) []Step {
	departureWake := c.wake() + threeHours
	departureBreakfast := min(c.breakfast()+threeHours, c.wake()+fiveHours)

	return []Step{
		//Caffeine
		steps.NoCaffeine{Start: c.departureLess3At(c.wake()), End: c.departureLess3At(t1500)},
		steps.CaffeineOk{Start: c.departureLess3At(t1500), End: c.departureLess3At(t1630)},
		steps.NoCaffeine{Start: c.departureLess3At(t1630), End: c.departureLess2At(t1500)},
		steps.CaffeineOk{Start: c.departureLess2At(t1500), End: c.departureLess2At(t1630)},
		steps.NoCaffeine{Start: c.departureLess2At(t1630), End: c.departureLess1At(t1500)},
		steps.CaffeineOk{Start: c.departureLess1At(t1500), End: c.departureLess1At(t1630)},
		steps.NoCaffeine{Start: c.departureLess1At(t1630), End: c.departureAt(departureWake)},
		steps.Caffeine3C{Start: c.departureAt(departureWake), End: c.departureAt(t1130)},
		steps.NoCaffeine{Start: c.departureAt(t1130), End: c.arrivalAt(c.sleep())},

		//Meals
		steps.HeavyBreakfast{At: c.departureLess3At(c.breakfast())},
		steps.HeavyLunch{At: c.departureLess3At(c.lunch())},
		steps.HeavyDinner{At: c.departureLess3At(c.dinner())},
		steps.LightBreakfast{At: c.departureLess2At(c.breakfast())},
		steps.LightLunch{At: c.departureLess2At(c.lunch())},
		steps.LightDinner{At: c.departureLess2At(c.dinner())},
		steps.NoSnack{Start: c.departureLess2At(c.dinner() + oneHour), End: c.departureLess2At(c.sleep())},
		steps.HeavyBreakfast{At: c.departureLess1At(c.breakfast())},
		steps.HeavyLunch{At: c.departureLess1At(c.lunch())},
		steps.HeavyDinner{At: c.departureLess1At(c.dinner())},
		steps.LightBreakfast{At: c.departureAt(departureBreakfast)},
		steps.HeavyBreakfast{At: c.arrivalAt(c.breakfast())},
		steps.HeavyLunch{At: c.arrivalAt(c.lunch())},
		steps.HeavyDinner{At: c.arrivalAt(c.dinner())},

		//Sleep
		steps.Sleep{Start: c.departureLess3At(c.sleep()), End: c.departureLess2At(c.wake())},
		steps.Sleep{Start: c.departureLess2At(c.sleep()), End: c.departureLess1At(c.wake())},
		steps.Sleep{Start: c.departureLess1At(c.sleep()), End: c.departureAt(departureWake)},
		steps.Sleep{Start: c.departureAt(departureBreakfast + oneHour), End: c.arrivalAt(c.breakfast())},
		steps.NoNap{Start: c.arrivalAt(c.breakfast()), End: c.arrivalAt(c.sleep())},
		steps.Sleep{Start: c.arrivalAt(c.sleep()), End: c.arrivalPlus1At(c.wake())},

		//Events
		steps.SetWatch{At: c.departureAt(departureBreakfast + oneHour)},
		c.arrivalStep(),
	}
}

func West910(c *Calculator) []Step {
	departureWake := c.wake() + threeHours
	departureBreakfast := c.breakfast() + threeHours
	departureLunch := min(c.lunch()+threeHours, c.breakfast()+sevenHours)

	return []Step{
		//Caffeine
		steps.NoCaffeine{Start: c.departureLess3At(c.wake()), End: c.departureLess3At(t1500)},
		steps.CaffeineOk{Start: c.departureLess3At(t1500), End: c.departureLess3At(t1630)},
		steps.NoCaffeine{Start: c.departureLess3At(t1630), End: c.departureLess2At(t1500)},
		steps.CaffeineOk{Start: c.departureLess2At(t1500), End: c.departureLess2At(t1630)},
		steps.NoCaffeine{Start: c.departureLess2At(t1630), End: c.departureLess1At(t1500)},
		steps.CaffeineOk{Start: c.departureLess1At(t1500), End: c.departureLess1At(t1630)},
		steps.NoCaffeine{Start: c.departureLess1At(t1630), End: c.departureAt(departureWake)},
		steps.Caffeine3C{Start: c.departureAt(departureWake), End: c.departureAt(departureWake + oneHour)},
		steps.NoCaffeine{Start: c.departureAt(departureWake + oneHour), End: c.arrivalAt(c.sleep())},

		//Meals
		steps.HeavyBreakfast{At: c.departureLess3At(c.breakfast())},
		steps.HeavyLunch{At: c.departureLess3At(c.lunch())},
		steps.HeavyDinner{At: c.departureLess3At(c.dinner())},
		steps.LightBreakfast{At: c.departureLess2At(c.breakfast())},
		steps.LightLunch{At: c.departureLess2At(c.lunch())},
		steps.LightDinner{At: c.departureLess2At(c.dinner())},
		steps.NoSnack{Start: c.departureLess2At(c.dinner() + oneHour), End: c.departureLess2At(c.sleep())},
		steps.HeavyBreakfast{At: c.departureLess1At(c.breakfast())},
		steps.HeavyLunch{At: c.departureLess1At(c.lunch())},
		steps.HeavyDinner{At: c.departureLess1At(c.dinner())},
		steps.LightBreakfast{At: c.departureAt(departureBreakfast)},
		steps.LightLunch{At: c.departureAt(departureLunch)},
		steps.HeavyBreakfast{At: c.arrivalAt(c.breakfast())},
		steps.HeavyLunch{At: c.arrivalAt(c.lunch())},
		steps.HeavyDinner{At: c.arrivalAt(c.dinner())},

		//Sleep
		steps.Sleep{Start: c.departureLess3At(c.sleep()), End: c.departureLess2At(c.wake())},
		steps.Sleep{Start: c.departureLess2At(c.sleep()), End: c.departureLess1At(c.wake())},
		steps.Sleep{Start: c.departureLess1At(c.sleep()), End: c.departureAt(departureWake)},
		steps.Sleep{Start: c.departureAt(departureLunch + oneHour), End: c.arrivalAt(c.breakfast())},
		steps.NoNap{Start: c.arrivalAt(c.breakfast()), End: c.arrivalAt(c.sleep())},
		steps.Sleep{Start: c.arrivalAt(c.sleep()), End: c.arrivalPlus1At(c.wake())},

		//Events
		steps.SetWatch{At: c.departureAt(departureLunch + oneHour)},
		c.arrivalStep(),
	}
}

func Both1112(c *Calculator) []Step {
	arrivalSleep := min(c.sleep(), t2200)

	return []Step{
		//Caffeine
		steps.NoCaffeine{Start: c.departureLess4At(c.wake()), End: c.departureLess4At(t1500)},
		steps.CaffeineOk{Start: c.departureLess4At(t1500), End: c.departureLess4At(t1630)},
		steps.NoCaffeine{Start: c.departureLess4At(t1630), End: c.departureLess3At(t1500)},
		steps.CaffeineOk{Start: c.departureLess3At(t1500), End: c.departureLess3At(t1630)},
		steps.NoCaffeine{Start: c.departureLess3At(t1630), End: c.departureLess2At(t0700)},
		steps.CaffeineOk{Start: c.departureLess2At(t0700), End: c.departureLess2At(t1130)},
		steps.NoCaffeine{Start: c.departureLess2At(t1130), End: c.departureLess1At(t0700)},
		steps.Caffeine3C{Start: c.departureLess1At(t0700), End: c.departureLess1At(t1130)},
		steps.NoCaffeine{Start: c.departureLess1At(t1130), End: c.arrivalAt(c.sleep())},

		//Meals
		steps.HeavyBreakfast{At: c.departureLess4At(c.breakfast())},
		steps.HeavyLunch{At: c.departureLess4At(c.lunch())},
		steps.HeavyDinner{At: c.departureLess4At(c.dinner())},
		steps.LightBreakfast{At: c.departureLess3At(c.breakfast())},
		steps.LightLunch{At: c.departureLess3At(c.lunch())},
		steps.LightDinner{At: c.departureLess3At(c.dinner())},
		steps.NoSnack{Start: c.departureLess3At(c.dinner() + oneHour), End: c.departureLess3At(c.sleep())},
		steps.HeavyBreakfast{At: c.departureLess2At(c.breakfast())},
		steps.HeavyLunch{At: c.departureLess2At(c.lunch())},
		steps.HeavyDinner{At: c.departureLess2At(c.dinner())},
		steps.NoSnack{Start: c.departureLess2At(c.dinner() + oneHour), End: c.departureLess2At(c.sleep())},
		steps.LightBreakfast{At: c.departureLess1At(c.breakfast())},
		steps.LightLunch{At: c.departureLess1At(c.lunch())},
		steps.HeavyBreakfast{At: c.arrivalAt(c.breakfast())},
		steps.HeavyLunch{At: c.arrivalAt(c.lunch())},
		steps.HeavyDinner{At: c.arrivalAt(c.dinner())},
		steps.NoSnack{Start: c.arrivalAt(c.dinner() + oneHour), End: c.arrivalAt(c.sleep())},

		//Sleep
		steps.Sleep{Start: c.departureLess4At(c.sleep()), End: c.departureLess3At(c.wake())},
		steps.Sleep{Start: c.departureLess3At(c.sleep()), End: c.departureLess2At(c.wake())},
		steps.Sleep{Start: c.departureLess2At(c.sleep()), End: c.departureLess1At(c.wake())},
		steps.Sleep{Start: c.departureLess1At(c.lunch() + oneHour), End: c.arrivalAt(c.wake())},
		steps.NoNap{Start: c.arrivalAt(c.wake()), End: c.arrivalAt(arrivalSleep)},
		steps.Sleep{c.arrivalAt(arrivalSleep), c.arrivalPlus1At(c.wake())},

		//Events
		steps.SetWatch{At: c.departureLess1At(c.lunch() + oneHour)},
		c.arrivalStep(),
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
