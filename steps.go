package timegiverserver

import "time"

type NoCaffeine struct {
	start time.Time
	end   time.Time
}

func (s NoCaffeine) ToIcs(lang string) string {
	return ``
}

type CaffeineOk struct {
	start time.Time
	end   time.Time
}

func (s CaffeineOk) ToIcs(lang string) string {
	return ``
}

type Caffeine3C struct {
	start time.Time
	end   time.Time
}

func (s Caffeine3C) ToIcs(lang string) string {
	return ``
}

type Caffeine2C struct {
	start time.Time
	end   time.Time
}

func (s Caffeine2C) ToIcs(lang string) string {
	return ``
}

type LightBreakfast struct {
	start time.Time
}

func (s LightBreakfast) ToIcs(lang string) string {
	return ``
}

type LightLunch struct {
	start time.Time
}

func (s LightLunch) ToIcs(lang string) string {
	return ``
}

type LightDinner struct {
	start time.Time
}

func (s LightDinner) ToIcs(lang string) string {
	return ``
}

type LightDinnerOptional struct {
	start time.Time
}

func (s LightDinnerOptional) ToIcs(lang string) string {
	return ``
}

type HeavyBreakfast struct {
	start time.Time
}

func (s HeavyBreakfast) ToIcs(lang string) string {
	return ``
}

type HeavyLunch struct {
	start time.Time
}

func (s HeavyLunch) ToIcs(lang string) string {
	return ``
}

type HeavyDinner struct {
	start time.Time
}

func (s HeavyDinner) ToIcs(lang string) string {
	return ``
}

type NoSnack struct {
	start time.Time
	end   time.Time
}

func (s NoSnack) ToIcs(lang string) string {
	return ``
}

type Sleep struct {
	start time.Time
	end   time.Time
}

func (s Sleep) ToIcs(lang string) string {
	return ``
}

type NoNap struct {
	start time.Time
	end   time.Time
}

func (s NoNap) ToIcs(lang string) string {
	return ``
}

type SetWatch struct {
	at time.Time
}

func (s SetWatch) ToIcs(lang string) string {
	return ``
}

type Arrive struct {
	at time.Time
}

func (s Arrive) ToIcs(lang string) string {
	return ``
}
