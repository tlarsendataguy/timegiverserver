package timegiverserver

import (
	"time"
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
		NoCaffeine{start: c.departureLess2At(c.wake()), end: c.departureLess2At(t1500)},
		CaffeineOk{start: c.departureLess2At(t1500), end: c.departureLess2At(t1630)},
		NoCaffeine{start: c.departureLess2At(t1630), end: c.departureLess1At(t1800)},
		Caffeine3C{start: c.departureLess1At(t1800), end: c.departureLess1At(t1900)},
		NoCaffeine{start: c.departureLess1At(t1900), end: c.arrivalAt(t2400)},

		//Meals
		LightBreakfast{start: c.departureLess1At(c.breakfast())},
		LightLunch{start: c.departureLess1At(c.lunch())},
		LightDinner{start: c.departureLess1At(c.dinner())},
		HeavyBreakfast{start: c.arrivalAt(c.breakfast())},
		HeavyLunch{start: c.arrivalAt(c.lunch())},
		HeavyDinner{start: c.arrivalAt(c.dinner())},

		//Sleep
		Sleep{start: c.departureLess2At(c.sleep()), end: c.departureLess1At(c.wake())},
		Sleep{start: c.departureLess1At(c.sleep()), end: c.arrivalAt(c.wake())},
		Sleep{start: c.arrivalAt(c.sleep()), end: c.arrivalPlus1At(c.wake())},

		//Events
		SetWatch{at: c.arrivalAt(c.breakfast())},
		c.arrivalStep(),
	}
}

func East34(c *Calculator) []Step {
	pre1Dinner := min(c.dinner(), t1800)
	pre1Sleep := max(pre1Dinner+oneHour, t1800)
	arrivalSleep := min(c.sleep(), t2230)

	return []Step{
		//Caffeine
		NoCaffeine{start: c.departureLess4At(c.wake()), end: c.departureLess4At(t1500)},
		CaffeineOk{start: c.departureLess4At(t1500), end: c.departureLess4At(t1630)},
		NoCaffeine{start: c.departureLess4At(t1630), end: c.departureLess3At(t1500)},
		CaffeineOk{start: c.departureLess3At(t1500), end: c.departureLess3At(t1630)},
		NoCaffeine{start: c.departureLess3At(t1630), end: c.departureLess2At(t1500)},
		CaffeineOk{start: c.departureLess2At(t1500), end: c.departureLess2At(t1630)},
		NoCaffeine{start: c.departureLess2At(t1630), end: c.departureLess1At(t1800)},
		Caffeine3C{start: c.departureLess1At(t1800), end: c.departureLess1At(t1900)},
		NoCaffeine{start: c.departureLess1At(t1900), end: c.arrivalAt(t2400)},

		//Meals
		HeavyBreakfast{start: c.departureLess2At(c.breakfast())},
		HeavyLunch{start: c.departureLess2At(c.lunch())},
		HeavyDinner{start: c.departureLess2At(c.dinner())},
		LightBreakfast{start: c.departureLess1At(c.breakfast())},
		LightLunch{start: c.departureLess1At(c.lunch())},
		LightDinner{start: c.departureLess1At(pre1Dinner)},
		HeavyBreakfast{start: c.arrivalAt(c.breakfast())},
		HeavyLunch{start: c.arrivalAt(c.lunch())},
		HeavyDinner{start: c.arrivalAt(c.dinner())},

		//Sleep
		Sleep{start: c.departureLess4At(c.sleep()), end: c.departureLess3At(c.wake())},
		Sleep{start: c.departureLess3At(c.sleep()), end: c.departureLess2At(c.wake())},
		Sleep{start: c.departureLess2At(c.sleep()), end: c.departureLess1At(c.wake())},
		Sleep{start: c.departureLess1At(pre1Sleep), end: c.arrivalAt(c.wake())},
		NoNap{start: c.arrivalAt(c.wake()), end: c.arrivalAt(c.sleep())},
		Sleep{start: c.arrivalAt(arrivalSleep), end: c.arrivalPlus1At(c.wake())},

		//Events
		SetWatch{at: c.departureLess1At(t1800)},
		c.arrivalStep(),
	}
}

func East56(c *Calculator) []Step {
	pre1Dinner := min(c.dinner(), t1800)
	pre1Sleep := max(pre1Dinner+oneHour, t1800)
	arrivalSleep := min(c.sleep(), t2200)

	return []Step{
		//Caffeine
		NoCaffeine{start: c.departureLess4At(c.wake()), end: c.departureLess4At(t1500)},
		CaffeineOk{start: c.departureLess4At(t1500), end: c.departureLess4At(t1630)},
		NoCaffeine{start: c.departureLess4At(t1630), end: c.departureLess3At(t1500)},
		CaffeineOk{start: c.departureLess3At(t1500), end: c.departureLess3At(t1630)},
		NoCaffeine{start: c.departureLess3At(t1630), end: c.departureLess2At(t1500)},
		CaffeineOk{start: c.departureLess2At(t1500), end: c.departureLess2At(t1630)},
		NoCaffeine{start: c.departureLess2At(t1630), end: c.departureLess1At(t1800)},
		Caffeine3C{start: c.departureLess1At(t1800), end: c.departureLess1At(t1900)},
		NoCaffeine{start: c.departureLess1At(t1900), end: c.arrivalAt(t2400)},

		//Meals
		HeavyBreakfast{start: c.departureLess4At(c.breakfast())},
		HeavyLunch{start: c.departureLess4At(c.lunch())},
		HeavyDinner{start: c.departureLess4At(c.dinner())},
		LightBreakfast{start: c.departureLess3At(c.breakfast())},
		LightLunch{start: c.departureLess3At(c.lunch())},
		LightDinner{start: c.departureLess3At(c.dinner())},
		NoSnack{start: c.departureLess3At(c.dinner() + oneHour), end: c.departureLess3At(c.sleep())},
		HeavyBreakfast{start: c.departureLess2At(c.breakfast())},
		HeavyLunch{start: c.departureLess2At(c.lunch())},
		HeavyDinner{start: c.departureLess2At(c.dinner())},
		NoSnack{start: c.departureLess2At(c.dinner() + oneHour), end: c.departureLess2At(c.sleep())},
		LightBreakfast{start: c.departureLess1At(c.breakfast())},
		LightLunch{start: c.departureLess1At(c.lunch())},
		LightDinner{start: c.departureLess1At(pre1Dinner)},
		HeavyBreakfast{start: c.arrivalAt(c.breakfast())},
		HeavyLunch{start: c.arrivalAt(c.lunch())},
		HeavyDinner{start: c.arrivalAt(c.dinner())},
		NoSnack{start: c.arrivalAt(c.dinner() + oneHour), end: c.arrivalAt(c.sleep())},

		//Sleep
		Sleep{start: c.departureLess4At(c.sleep()), end: c.departureLess3At(c.wake())},
		Sleep{start: c.departureLess3At(c.sleep()), end: c.departureLess2At(c.wake())},
		Sleep{start: c.departureLess2At(c.sleep()), end: c.departureLess1At(c.wake())},
		Sleep{start: c.departureLess1At(pre1Sleep), end: c.arrivalAt(c.wake())},
		NoNap{start: c.arrivalAt(c.wake()), end: c.arrivalAt(arrivalSleep)},
		Sleep{start: c.arrivalAt(arrivalSleep), end: c.arrivalPlus1At(c.wake())},

		//Events
		SetWatch{at: c.departureLess1At(t1800)},
		c.arrivalStep(),
	}
}

func East78(c *Calculator) []Step {
	pre1Dinner := min(c.dinner(), t1700)
	pre1Sleep := max(pre1Dinner+oneHour, t1700)
	arrivalSleep := min(c.sleep(), t2200)

	return []Step{
		//Caffeine
		NoCaffeine{start: c.departureLess4At(c.wake()), end: c.departureLess4At(t1500)},
		CaffeineOk{start: c.departureLess4At(t1500), end: c.departureLess4At(t1630)},
		NoCaffeine{start: c.departureLess4At(t1630), end: c.departureLess3At(t1500)},
		CaffeineOk{start: c.departureLess3At(t1500), end: c.departureLess3At(t1630)},
		NoCaffeine{start: c.departureLess3At(t1630), end: c.departureLess2At(t1500)},
		CaffeineOk{start: c.departureLess2At(t1500), end: c.departureLess2At(t1630)},
		NoCaffeine{start: c.departureLess2At(t1630), end: c.departureLess1At(t1800)},
		Caffeine2C{start: c.departureLess1At(t1800), end: c.departureLess1At(t1900)},
		NoCaffeine{start: c.departureLess1At(t1900), end: c.arrivalAt(c.breakfast())},
		Caffeine2C{start: c.arrivalAt(c.breakfast()), end: c.arrivalAt(c.breakfast() + oneHour)},
		NoCaffeine{start: c.arrivalAt(c.breakfast() + oneHour), end: c.arrivalAt(c.sleep())},

		//Meals
		HeavyBreakfast{start: c.departureLess4At(c.breakfast())},
		HeavyLunch{start: c.departureLess4At(c.lunch())},
		HeavyDinner{start: c.departureLess4At(c.dinner())},
		LightBreakfast{start: c.departureLess3At(c.breakfast())},
		LightLunch{start: c.departureLess3At(c.lunch())},
		LightDinner{start: c.departureLess3At(c.dinner())},
		NoSnack{start: c.departureLess3At(c.dinner() + oneHour), end: c.departureLess3At(c.sleep())},
		HeavyBreakfast{start: c.departureLess2At(c.breakfast())},
		HeavyLunch{start: c.departureLess2At(c.lunch())},
		HeavyDinner{start: c.departureLess2At(c.dinner())},
		NoSnack{start: c.departureLess2At(c.dinner() + oneHour), end: c.departureLess2At(c.sleep())},
		LightBreakfast{start: c.departureLess1At(c.breakfast())},
		LightLunch{start: c.departureLess1At(c.lunch())},
		LightDinnerOptional{start: c.departureLess1At(pre1Dinner)},
		HeavyBreakfast{start: c.arrivalAt(c.breakfast())},
		HeavyLunch{start: c.arrivalAt(c.lunch())},
		HeavyDinner{start: c.arrivalAt(c.dinner())},
		NoSnack{start: c.arrivalAt(c.dinner() + oneHour), end: c.arrivalAt(c.sleep())},

		//Sleep
		Sleep{start: c.departureLess4At(c.sleep()), end: c.departureLess3At(c.wake())},
		Sleep{start: c.departureLess3At(c.sleep()), end: c.departureLess2At(c.wake())},
		Sleep{start: c.departureLess2At(c.sleep()), end: c.departureLess1At(c.wake())},
		Sleep{start: c.departureLess1At(pre1Sleep), end: c.arrivalAt(c.wake())},
		NoNap{start: c.arrivalAt(c.wake()), end: c.arrivalAt(arrivalSleep)},
		Sleep{start: c.arrivalAt(arrivalSleep), end: c.arrivalPlus1At(c.wake())},

		//Events
		SetWatch{at: c.departureLess1At(t1800)},
		c.arrivalStep(),
	}
}

func East910(c *Calculator) []Step {
	arrivalSleep := min(c.sleep(), t2200)
	arrivalDinner := min(c.dinner(), arrivalSleep-oneHour)

	return []Step{
		//Caffeine
		NoCaffeine{start: c.departureLess4At(c.wake()), end: c.departureLess4At(t1500)},
		CaffeineOk{start: c.departureLess4At(t1500), end: c.departureLess4At(t1630)},
		NoCaffeine{start: c.departureLess4At(t1630), end: c.departureLess3At(t1500)},
		CaffeineOk{start: c.departureLess3At(t1500), end: c.departureLess3At(t1630)},
		NoCaffeine{start: c.departureLess3At(t1630), end: c.departureLess2At(t1500)},
		CaffeineOk{start: c.departureLess2At(t1500), end: c.departureLess2At(t1630)},
		NoCaffeine{start: c.departureLess2At(t1630), end: c.arrivalAt(c.breakfast())},
		Caffeine3C{start: c.arrivalAt(c.breakfast()), end: c.arrivalAt(c.breakfast() + oneHour)},
		NoCaffeine{start: c.arrivalAt(c.breakfast() + oneHour), end: c.arrivalAt(c.sleep())},

		//Meals
		HeavyBreakfast{start: c.departureLess4At(c.breakfast())},
		HeavyLunch{start: c.departureLess4At(c.lunch())},
		HeavyDinner{start: c.departureLess4At(c.dinner())},
		LightBreakfast{start: c.departureLess3At(c.breakfast())},
		LightLunch{start: c.departureLess3At(c.lunch())},
		LightDinner{start: c.departureLess3At(c.dinner())},
		NoSnack{start: c.departureLess3At(c.dinner() + oneHour), end: c.departureLess3At(c.sleep())},
		HeavyBreakfast{start: c.departureLess2At(c.breakfast())},
		HeavyLunch{start: c.departureLess2At(c.lunch())},
		HeavyDinner{start: c.departureLess2At(c.dinner())},
		NoSnack{start: c.departureLess2At(c.dinner() + oneHour), end: c.departureLess2At(c.sleep())},
		LightBreakfast{start: c.departureLess1At(c.breakfast())},
		LightLunch{start: c.departureLess1At(c.lunch())},
		HeavyBreakfast{start: c.arrivalAt(c.breakfast())},
		HeavyLunch{start: c.arrivalAt(c.lunch())},
		HeavyDinner{start: c.arrivalAt(arrivalDinner)},
		NoSnack{start: c.arrivalAt(arrivalDinner + oneHour), end: c.arrivalAt(c.sleep())},

		//Sleep
		Sleep{start: c.departureLess4At(c.sleep()), end: c.departureLess3At(c.wake())},
		Sleep{start: c.departureLess3At(c.sleep()), end: c.departureLess2At(c.wake())},
		Sleep{start: c.departureLess2At(c.sleep()), end: c.departureLess1At(c.wake())},
		NoNap{start: c.arrivalAt(c.wake()), end: c.arrivalAt(c.sleep())},
		Sleep{start: c.arrivalAt(c.sleep()), end: c.arrivalPlus1At(c.wake())},

		//Events
		SetWatch{at: c.departureLess1At(c.lunch() + oneHour)},
		c.arrivalStep(),
	}
}

func West12(c *Calculator) []Step {
	return []Step{
		//Caffeine
		NoCaffeine{start: c.departureLess2At(c.wake()), end: c.departureLess2At(t1500)},
		CaffeineOk{start: c.departureLess2At(t1500), end: c.departureLess2At(t1630)},
		NoCaffeine{start: c.departureLess2At(t1630), end: c.departureLess1At(t0700)},
		CaffeineOk{start: c.departureLess1At(t0700), end: c.departureLess1At(t1100)},
		NoCaffeine{start: c.departureLess1At(t1100), end: c.departureAt(c.wake())},
		Caffeine3C{start: c.departureAt(c.wake()), end: c.departureAt(t1100)},
		NoCaffeine{start: c.departureAt(t1100), end: c.arrivalAt(c.sleep())},

		//Meals
		LightBreakfast{start: c.departureLess1At(c.breakfast())},
		LightLunch{start: c.departureLess1At(c.lunch())},
		LightDinner{start: c.departureLess1At(c.dinner())},
		HeavyBreakfast{start: c.arrivalAt(c.breakfast())},
		HeavyLunch{start: c.arrivalAt(c.lunch())},
		HeavyDinner{start: c.arrivalAt(c.dinner())},

		//Sleep
		Sleep{start: c.departureLess2At(c.sleep()), end: c.departureLess1At(c.wake())},
		Sleep{start: c.departureLess1At(c.sleep()), end: c.departureAt(c.wake())},
		Sleep{start: c.arrivalAt(c.sleep()), end: c.arrivalPlus1At(c.wake())},

		//Events
		SetWatch{at: c.arrivalAt(c.breakfast())},
		c.arrivalStep(),
	}
}

func West34(c *Calculator) []Step {
	return []Step{
		//Caffeine
		NoCaffeine{start: c.departureLess3At(c.wake()), end: c.departureLess3At(t1500)},
		CaffeineOk{start: c.departureLess3At(t1500), end: c.departureLess3At(t1630)},
		NoCaffeine{start: c.departureLess3At(t1630), end: c.departureLess2At(t1500)},
		CaffeineOk{start: c.departureLess2At(t1500), end: c.departureLess2At(t1630)},
		NoCaffeine{start: c.departureLess2At(t1630), end: c.departureLess1At(t0700)},
		CaffeineOk{start: c.departureLess1At(t0700), end: c.departureLess1At(t1100)},
		NoCaffeine{start: c.departureLess1At(t1100), end: c.departureAt(c.wake())},
		Caffeine3C{start: c.departureAt(c.wake()), end: c.departureAt(t1100)},
		NoCaffeine{start: c.departureAt(t1100), end: c.arrivalAt(c.sleep())},

		//Meals
		LightBreakfast{start: c.departureLess1At(c.breakfast())},
		LightLunch{start: c.departureLess1At(c.lunch())},
		LightDinner{start: c.departureLess1At(c.dinner())},
		HeavyBreakfast{start: c.arrivalAt(c.breakfast())},
		HeavyLunch{start: c.arrivalAt(c.lunch())},
		HeavyDinner{start: c.arrivalAt(c.dinner())},

		//Sleep
		Sleep{start: c.departureLess3At(c.sleep()), end: c.departureLess2At(c.wake())},
		Sleep{start: c.departureLess2At(c.sleep()), end: c.departureLess1At(c.wake())},
		Sleep{start: c.departureLess1At(c.sleep()), end: c.departureAt(c.wake())},
		Sleep{start: c.arrivalAt(c.sleep()), end: c.arrivalPlus1At(c.wake())},

		//Events
		SetWatch{at: c.arrivalAt(c.breakfast())},
		c.arrivalStep(),
	}
}

func West56(c *Calculator) []Step {
	departureLess1Wakeup := min(c.sleep()+twoHours, t1000)

	return []Step{
		//Caffeine
		NoCaffeine{start: c.departureLess4At(c.wake()), end: c.departureLess4At(t1500)},
		CaffeineOk{start: c.departureLess4At(t1500), end: c.departureLess4At(t1630)},
		NoCaffeine{start: c.departureLess4At(t1630), end: c.departureLess3At(t1500)},
		CaffeineOk{start: c.departureLess3At(t1500), end: c.departureLess3At(t1630)},
		NoCaffeine{start: c.departureLess3At(t1630), end: c.departureLess2At(t0700)},
		CaffeineOk{start: c.departureLess2At(t0700), end: c.departureLess2At(t1130)},
		NoCaffeine{start: c.departureLess2At(t1130), end: c.departureLess1At(departureLess1Wakeup)},
		Caffeine3C{start: c.departureLess1At(departureLess1Wakeup), end: c.departureLess1At(t1100)},
		NoCaffeine{start: c.departureLess1At(t1100), end: c.arrivalAt(c.sleep())},

		//Meals
		HeavyBreakfast{start: c.departureLess4At(c.breakfast())},
		HeavyLunch{start: c.departureLess4At(c.lunch())},
		HeavyDinner{start: c.departureLess4At(c.dinner())},
		LightBreakfast{start: c.departureLess3At(c.breakfast())},
		LightLunch{start: c.departureLess3At(c.lunch())},
		LightDinner{start: c.departureLess3At(c.dinner())},
		NoSnack{start: c.departureLess3At(c.dinner() + oneHour), end: c.departureLess3At(c.sleep())},
		HeavyBreakfast{start: c.departureLess2At(c.breakfast())},
		HeavyLunch{start: c.departureLess2At(c.lunch())},
		HeavyDinner{start: c.departureLess2At(c.dinner())},
		LightBreakfast{start: c.departureLess1At(c.breakfast() + twoHours)},
		LightLunch{start: c.departureLess1At(c.lunch() + twoHours)},
		LightDinner{start: c.departureLess1At(c.dinner() + twoHours)},
		HeavyBreakfast{start: c.arrivalAt(c.breakfast())},
		HeavyLunch{start: c.arrivalAt(c.lunch())},
		HeavyDinner{start: c.arrivalAt(c.dinner())},

		//Sleep
		Sleep{start: c.departureLess4At(c.sleep()), end: c.departureLess3At(c.wake())},
		Sleep{start: c.departureLess3At(c.sleep()), end: c.departureLess2At(c.wake())},
		Sleep{start: c.departureLess2At(c.sleep()), end: c.departureLess1At(departureLess1Wakeup)},
		Sleep{start: c.arrivalLess1At(c.sleep()), end: c.arrivalAt(c.wake())},
		Sleep{start: c.arrivalAt(c.sleep()), end: c.arrivalPlus1At(c.wake())},

		//Events
		SetWatch{at: c.departureLess1At(c.dinner() + threeHours)},
		c.arrivalStep(),
	}
}

func West78(c *Calculator) []Step {
	departureWake := c.wake() + threeHours
	departureBreakfast := min(c.breakfast()+threeHours, c.wake()+fiveHours)

	return []Step{
		//Caffeine
		NoCaffeine{start: c.departureLess3At(c.wake()), end: c.departureLess3At(t1500)},
		CaffeineOk{start: c.departureLess3At(t1500), end: c.departureLess3At(t1630)},
		NoCaffeine{start: c.departureLess3At(t1630), end: c.departureLess2At(t1500)},
		CaffeineOk{start: c.departureLess2At(t1500), end: c.departureLess2At(t1630)},
		NoCaffeine{start: c.departureLess2At(t1630), end: c.departureLess1At(t1500)},
		CaffeineOk{start: c.departureLess1At(t1500), end: c.departureLess1At(t1630)},
		NoCaffeine{start: c.departureLess1At(t1630), end: c.departureAt(departureWake)},
		Caffeine3C{start: c.departureAt(departureWake), end: c.departureAt(t1130)},
		NoCaffeine{start: c.departureAt(t1130), end: c.arrivalAt(c.sleep())},

		//Meals
		HeavyBreakfast{start: c.departureLess3At(c.breakfast())},
		HeavyLunch{start: c.departureLess3At(c.lunch())},
		HeavyDinner{start: c.departureLess3At(c.dinner())},
		LightBreakfast{start: c.departureLess2At(c.breakfast())},
		LightLunch{start: c.departureLess2At(c.lunch())},
		LightDinner{start: c.departureLess2At(c.dinner())},
		NoSnack{start: c.departureLess2At(c.dinner() + oneHour), end: c.departureLess2At(c.sleep())},
		HeavyBreakfast{start: c.departureLess1At(c.breakfast())},
		HeavyLunch{start: c.departureLess1At(c.lunch())},
		HeavyDinner{start: c.departureLess1At(c.dinner())},
		LightBreakfast{start: c.departureAt(departureBreakfast)},
		HeavyBreakfast{start: c.arrivalAt(c.breakfast())},
		HeavyLunch{start: c.arrivalAt(c.lunch())},
		HeavyDinner{start: c.arrivalAt(c.dinner())},

		//Sleep
		Sleep{start: c.departureLess3At(c.sleep()), end: c.departureLess2At(c.wake())},
		Sleep{start: c.departureLess2At(c.sleep()), end: c.departureLess1At(c.wake())},
		Sleep{start: c.departureLess1At(c.sleep()), end: c.departureAt(departureWake)},
		Sleep{start: c.departureAt(departureBreakfast + oneHour), end: c.arrivalAt(c.breakfast())},
		NoNap{start: c.arrivalAt(c.breakfast()), end: c.arrivalAt(c.sleep())},
		Sleep{start: c.arrivalAt(c.sleep()), end: c.arrivalPlus1At(c.wake())},

		//Events
		SetWatch{at: c.departureAt(departureBreakfast + oneHour)},
		c.arrivalStep(),
	}
}

func West910(c *Calculator) []Step {
	departureWake := c.wake() + threeHours
	departureBreakfast := c.breakfast() + threeHours
	departureLunch := min(c.lunch()+threeHours, c.breakfast()+sevenHours)

	return []Step{
		//Caffeine
		NoCaffeine{start: c.departureLess3At(c.wake()), end: c.departureLess3At(t1500)},
		CaffeineOk{start: c.departureLess3At(t1500), end: c.departureLess3At(t1630)},
		NoCaffeine{start: c.departureLess3At(t1630), end: c.departureLess2At(t1500)},
		CaffeineOk{start: c.departureLess2At(t1500), end: c.departureLess2At(t1630)},
		NoCaffeine{start: c.departureLess2At(t1630), end: c.departureLess1At(t1500)},
		CaffeineOk{start: c.departureLess1At(t1500), end: c.departureLess1At(t1630)},
		NoCaffeine{start: c.departureLess1At(t1630), end: c.departureAt(departureWake)},
		Caffeine3C{start: c.departureAt(departureWake), end: c.departureAt(departureWake + oneHour)},
		NoCaffeine{start: c.departureAt(departureWake + oneHour), end: c.arrivalAt(c.sleep())},

		//Meals
		HeavyBreakfast{start: c.departureLess3At(c.breakfast())},
		HeavyLunch{start: c.departureLess3At(c.lunch())},
		HeavyDinner{start: c.departureLess3At(c.dinner())},
		LightBreakfast{start: c.departureLess2At(c.breakfast())},
		LightLunch{start: c.departureLess2At(c.lunch())},
		LightDinner{start: c.departureLess2At(c.dinner())},
		NoSnack{start: c.departureLess2At(c.dinner() + oneHour), end: c.departureLess2At(c.sleep())},
		HeavyBreakfast{start: c.departureLess1At(c.breakfast())},
		HeavyLunch{start: c.departureLess1At(c.lunch())},
		HeavyDinner{start: c.departureLess1At(c.dinner())},
		LightBreakfast{start: c.departureAt(departureBreakfast)},
		LightLunch{start: c.departureAt(departureLunch)},
		HeavyBreakfast{start: c.arrivalAt(c.breakfast())},
		HeavyLunch{start: c.arrivalAt(c.lunch())},
		HeavyDinner{start: c.arrivalAt(c.dinner())},

		//Sleep
		Sleep{start: c.departureLess3At(c.sleep()), end: c.departureLess2At(c.wake())},
		Sleep{start: c.departureLess2At(c.sleep()), end: c.departureLess1At(c.wake())},
		Sleep{start: c.departureLess1At(c.sleep()), end: c.departureAt(departureWake)},
		Sleep{start: c.departureAt(departureLunch + oneHour), end: c.arrivalAt(c.breakfast())},
		NoNap{start: c.arrivalAt(c.breakfast()), end: c.arrivalAt(c.sleep())},
		Sleep{start: c.arrivalAt(c.sleep()), end: c.arrivalPlus1At(c.wake())},

		//Events
		SetWatch{at: c.departureAt(departureLunch + oneHour)},
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
