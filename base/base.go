package base

import "time"

type Dates struct {
	Departure      time.Time
	DepartureLess1 time.Time
	DepartureLess2 time.Time
	DepartureLess3 time.Time
	DepartureLess4 time.Time
	Arrival        time.Time
	ArrivalLess1   time.Time
	ArrivalPlus1   time.Time
	ArriveAt       time.Time
}

func InitializeDates(arrival time.Time, departureOffset, arrivalOffset float64) Dates {
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

	arriveAt := arrivalBase.Add((time.Hour * time.Duration(arrival.Hour())) + (time.Minute * time.Duration(arrival.Minute())))

	return Dates{
		Departure:      departureBase,
		DepartureLess1: addDays(departureBase, -1),
		DepartureLess2: addDays(departureBase, -2),
		DepartureLess3: addDays(departureBase, -3),
		DepartureLess4: addDays(departureBase, -4),
		Arrival:        arrivalBase,
		ArrivalLess1:   addDays(arrivalBase, -1),
		ArrivalPlus1:   addDays(arrivalBase, 1),
		ArriveAt:       arriveAt,
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
