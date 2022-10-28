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
		NoCaffeine{start: c.Dates.DepartureLess2.Add(c.Routine.Wake), end: c.Dates.DepartureLess2.Add(t1500)},
		CaffeineOk{start: c.Dates.DepartureLess2.Add(t1500), end: c.Dates.DepartureLess2.Add(t1630)},
		NoCaffeine{start: c.Dates.DepartureLess2.Add(t1630), end: c.Dates.DepartureLess1.Add(t1800)},
		Caffeine3C{start: c.Dates.DepartureLess1.Add(t1800), end: c.Dates.DepartureLess1.Add(t1900)},
		NoCaffeine{start: c.Dates.DepartureLess1.Add(t1900), end: c.Dates.Arrival.Add(t2400)},

		//Meals
		LightBreakfast{start: c.Dates.DepartureLess1.Add(c.Routine.Breakfast)},
		LightLunch{start: c.Dates.DepartureLess1.Add(c.Routine.Lunch)},
		LightDinner{start: c.Dates.DepartureLess1.Add(c.Routine.Dinner)},
		HeavyBreakfast{start: c.Dates.Arrival.Add(c.Routine.Breakfast)},
		HeavyLunch{start: c.Dates.Arrival.Add(c.Routine.Lunch)},
		HeavyDinner{start: c.Dates.Arrival.Add(c.Routine.Dinner)},

		//Sleep
		Sleep{start: c.Dates.DepartureLess1.Add(c.Routine.Sleep), end: c.Dates.Arrival.Add(c.Routine.Wake)},
		Sleep{start: c.Dates.Arrival.Add(c.Routine.Sleep), end: c.Dates.ArrivalPlus1.Add(c.Routine.Wake)},

		//Events
		SetWatch{at: c.Dates.Arrival.Add(c.Routine.Breakfast)},
		Arrive{at: c.Dates.ArriveAt},
	}
}
