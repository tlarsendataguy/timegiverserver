package timegiverserver

import "time"

const t0700 = time.Hour * 7
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
		Sleep{start: c.departureLess1At(c.sleep()), end: c.arrivalAt(c.wake())},
		Sleep{start: c.arrivalAt(c.sleep()), end: c.arrivalPlus1At(c.wake())},

		//Events
		SetWatch{at: c.arrivalAt(c.breakfast())},
		c.arrivalStep(),
	}
}

func East34(c *Calculator) []Step {
	pre1Dinner := min(c.dinner(), t1800)
	pre1Sleep := max(pre1Dinner+time.Hour, t1800)
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
	pre1Sleep := max(pre1Dinner+time.Hour, t1800)
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
		NoSnack{start: c.departureLess3At(c.dinner() + time.Hour), end: c.departureLess3At(c.sleep())},
		HeavyBreakfast{start: c.departureLess2At(c.breakfast())},
		HeavyLunch{start: c.departureLess2At(c.lunch())},
		HeavyDinner{start: c.departureLess2At(c.dinner())},
		NoSnack{start: c.departureLess2At(c.dinner() + time.Hour), end: c.departureLess3At(c.sleep())},
		LightBreakfast{start: c.departureLess1At(c.breakfast())},
		LightLunch{start: c.departureLess1At(c.lunch())},
		LightDinner{start: c.departureLess1At(pre1Dinner)},
		HeavyBreakfast{start: c.arrivalAt(c.breakfast())},
		HeavyLunch{start: c.arrivalAt(c.lunch())},
		HeavyDinner{start: c.arrivalAt(c.dinner())},
		NoSnack{start: c.arrivalAt(c.dinner() + time.Hour), end: c.arrivalAt(c.sleep())},

		//Sleep
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
	pre1Sleep := max(pre1Dinner+time.Hour, t1700)
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
		Caffeine2C{start: c.arrivalAt(c.breakfast()), end: c.arrivalAt(c.breakfast() + time.Hour)},
		NoCaffeine{start: c.arrivalAt(c.breakfast() + time.Hour), end: c.arrivalAt(c.sleep())},

		//Meals
		HeavyBreakfast{start: c.departureLess4At(c.breakfast())},
		HeavyLunch{start: c.departureLess4At(c.lunch())},
		HeavyDinner{start: c.departureLess4At(c.dinner())},
		LightBreakfast{start: c.departureLess3At(c.breakfast())},
		LightLunch{start: c.departureLess3At(c.lunch())},
		LightDinner{start: c.departureLess3At(c.dinner())},
		NoSnack{start: c.departureLess3At(c.dinner() + time.Hour), end: c.departureLess3At(c.sleep())},
		HeavyBreakfast{start: c.departureLess2At(c.breakfast())},
		HeavyLunch{start: c.departureLess2At(c.lunch())},
		HeavyDinner{start: c.departureLess2At(c.dinner())},
		NoSnack{start: c.departureLess2At(c.dinner() + time.Hour), end: c.departureLess3At(c.sleep())},
		LightBreakfast{start: c.departureLess1At(c.breakfast())},
		LightLunch{start: c.departureLess1At(c.lunch())},
		LightDinnerOptional{start: c.departureLess1At(pre1Dinner)},
		HeavyBreakfast{start: c.arrivalAt(c.breakfast())},
		HeavyLunch{start: c.arrivalAt(c.lunch())},
		HeavyDinner{start: c.arrivalAt(c.dinner())},
		NoSnack{start: c.arrivalAt(c.dinner() + time.Hour), end: c.arrivalAt(c.sleep())},

		//Sleep
		Sleep{start: c.departureLess1At(pre1Sleep), end: c.arrivalAt(c.wake())},
		NoNap{start: c.arrivalAt(c.wake()), end: c.arrivalAt(arrivalSleep)},
		Sleep{start: c.arrivalAt(arrivalSleep), end: c.arrivalPlus1At(c.wake())},

		//Events
		SetWatch{at: c.departureLess1At(t1800)},
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
