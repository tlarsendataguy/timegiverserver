package base_date

import (
	"testing"
	"time"
)

var arrival = date(2017, 1, 15, 12)

func TestEastboundBaseDatesPositiveOffset(t *testing.T) {
	result := CalcBase(arrival, 3, 5)
	checkDate(t, `2017-01-14 21:00`, result.Departure)
	checkDate(t, `2017-01-14 19:00`, result.Arrival)
}

func TestEastboundBaseDatesNegativeOffset(t *testing.T) {
	result := CalcBase(arrival, -5, -3)
	checkDate(t, `2017-01-15 05:00`, result.Departure)
	checkDate(t, `2017-01-15 03:00`, result.Arrival)
}

func TestEastboundBaseDatesAcrossDateline(t *testing.T) {
	result := CalcBase(arrival, 10, -10)
	checkDate(t, `2017-01-15 14:00`, result.Departure)
	checkDate(t, `2017-01-15 10:00`, result.Arrival)
}

func TestWestboundBaseDatesPositiveOffset(t *testing.T) {
	result := CalcBase(arrival, 5, 3)
	checkDate(t, `2017-01-14 19:00`, result.Departure)
	checkDate(t, `2017-01-14 21:00`, result.Arrival)
}

func TestWestboundBaseDatesNegativeOffset(t *testing.T) {
	result := CalcBase(arrival, -3, -5)
	checkDate(t, `2017-01-15 03:00`, result.Departure)
	checkDate(t, `2017-01-15 05:00`, result.Arrival)
}

func TestWestboundBaseDatesAcrossDateline(t *testing.T) {
	result := CalcBase(arrival, -10, 10)
	checkDate(t, `2017-01-14 10:00`, result.Departure)
	checkDate(t, `2017-01-14 14:00`, result.Arrival)
}

func date(year, month, day, hour int) time.Time {
	return time.Date(year, time.Month(month), day, hour, 0, 0, 0, time.UTC)
}

func checkDate(t *testing.T, expected string, actual time.Time) {
	expectedDate, _ := time.Parse(`2006-01-02 15:04`, expected)
	if expectedDate.Equal(actual) {
		return
	}
	t.Fatalf(`expected %v but got %v`, expectedDate, actual)
}
