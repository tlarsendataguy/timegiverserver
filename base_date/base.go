package base_date

import "time"

type BaseDates struct {
	Departure time.Time
	Arrival   time.Time
}

func CalcBase(arrival time.Time, departureOffset, arrivalOffset float64) BaseDates {
	arrivalBase := toUtc(startOfDay(arrival), arrivalOffset)
	departureBase := toUtc(startOfDay(arrival), departureOffset)

	departureBase = addDays(departureBase, 4)
	for true {
		delta := departureBase.Sub(arrivalBase).Hours()
		if delta <= 12 {
			break
		}
		departureBase = addDays(departureBase, -1)
	}

	return BaseDates{
		Departure: departureBase,
		Arrival:   arrivalBase,
	}
}

func startOfDay(value time.Time) time.Time {
	return time.Date(value.Year(), value.Month(), value.Day(), 0, 0, 0, 0, time.UTC)
}

func addDays(value time.Time, days time.Duration) time.Time {
	return value.Add(time.Hour * 24 * days)
}

func toUtc(value time.Time, offset float64) time.Time {
	duration := int(float64(time.Hour) * -offset)
	return value.Add(time.Duration(duration))
}
