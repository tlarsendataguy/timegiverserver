package timegiverserver

import (
	"time"
	"timegiverserver/localization"
)

type NoCaffeine struct {
	start time.Time
	end   time.Time
}

func (s NoCaffeine) ToIcs(lang localization.Lang) string {
	return localization.NoCaffeine[lang].Title
}

type CaffeineOk struct {
	start time.Time
	end   time.Time
}

func (s CaffeineOk) ToIcs(lang localization.Lang) string {
	return ``
}

type Caffeine3C struct {
	start time.Time
	end   time.Time
}

func (s Caffeine3C) ToIcs(lang localization.Lang) string {
	return ``
}

type Caffeine2C struct {
	start time.Time
	end   time.Time
}

func (s Caffeine2C) ToIcs(lang localization.Lang) string {
	return ``
}

type LightBreakfast struct {
	at time.Time
}

func (s LightBreakfast) ToIcs(lang localization.Lang) string {
	return ``
}

type LightLunch struct {
	at time.Time
}

func (s LightLunch) ToIcs(lang localization.Lang) string {
	return ``
}

type LightDinner struct {
	at time.Time
}

func (s LightDinner) ToIcs(lang localization.Lang) string {
	return ``
}

type LightDinnerOptional struct {
	at time.Time
}

func (s LightDinnerOptional) ToIcs(lang localization.Lang) string {
	return ``
}

type HeavyBreakfast struct {
	at time.Time
}

func (s HeavyBreakfast) ToIcs(lang localization.Lang) string {
	return ``
}

type HeavyLunch struct {
	at time.Time
}

func (s HeavyLunch) ToIcs(lang localization.Lang) string {
	return ``
}

type HeavyDinner struct {
	at time.Time
}

func (s HeavyDinner) ToIcs(lang localization.Lang) string {
	return ``
}

type NoSnack struct {
	start time.Time
	end   time.Time
}

func (s NoSnack) ToIcs(lang localization.Lang) string {
	return ``
}

type Sleep struct {
	start time.Time
	end   time.Time
}

func (s Sleep) ToIcs(lang localization.Lang) string {
	return ``
}

type NoNap struct {
	start time.Time
	end   time.Time
}

func (s NoNap) ToIcs(lang localization.Lang) string {
	return ``
}

type SetWatch struct {
	at time.Time
}

func (s SetWatch) ToIcs(lang localization.Lang) string {
	return ``
}

type Arrive struct {
	at time.Time
}

func (s Arrive) ToIcs(lang localization.Lang) string {
	return ``
}
