package timegiverserver

import "time"

type NoCaffeine struct {
	start time.Time
	end   time.Time
}

func (s NoCaffeine) ToIcs() string {
	return ``
}

type CaffeineOk struct {
	start time.Time
	end   time.Time
}

func (s CaffeineOk) ToIcs() string {
	return ``
}

type Caffeine3C struct {
	start time.Time
	end   time.Time
}

func (s Caffeine3C) ToIcs() string {
	return ``
}

type LightBreakfast struct {
	start time.Time
}

func (s LightBreakfast) ToIcs() string {
	return ``
}

type LightLunch struct {
	start time.Time
}

func (s LightLunch) ToIcs() string {
	return ``
}

type LightDinner struct {
	start time.Time
}

func (s LightDinner) ToIcs() string {
	return ``
}

type HeavyBreakfast struct {
	start time.Time
}

func (s HeavyBreakfast) ToIcs() string {
	return ``
}

type HeavyLunch struct {
	start time.Time
}

func (s HeavyLunch) ToIcs() string {
	return ``
}

type HeavyDinner struct {
	start time.Time
}

func (s HeavyDinner) ToIcs() string {
	return ``
}

type Sleep struct {
	start time.Time
	end   time.Time
}

func (s Sleep) ToIcs() string {
	return ``
}

type SetWatch struct {
	at time.Time
}

func (s SetWatch) ToIcs() string {
	return ``
}

type Arrive struct {
	at time.Time
}

func (s Arrive) ToIcs() string {
	return ``
}
