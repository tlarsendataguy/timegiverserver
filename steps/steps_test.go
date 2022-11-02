package steps

import (
	"strings"
	"testing"
	"time"
	"timegiverserver/steps/localization"
)

var start = time.Date(2022, 1, 2, 3, 0, 12, 0, time.UTC)
var end = time.Date(2022, 1, 2, 4, 30, 0, 0, time.UTC)

func TestDateTimeFormatting(t *testing.T) {
	s := &strings.Builder{}
	fmtDt(s, start)
	actual := s.String()
	expected := `20220102T030000Z`
	if actual != expected {
		t.Fatalf(`expected %v but got %v`, expected, actual)
	}
}

func TestWrapShort(t *testing.T) {
	s := &strings.Builder{}
	wrap(s, `SUMMARY:`, `hello world`)
	actual := s.String()
	expected := `SUMMARY:hello world`
	if actual != expected {
		t.Fatalf(`expected '%v' but got '%v'`, expected, actual)
	}
}

func TestWrap2Lines(t *testing.T) {
	s := &strings.Builder{}
	wrap(s, `SUMMARY:`, `this line plus header is more than 75 bytes and must be wrapped onto 2 lines`)
	actual := s.String()
	expected := "SUMMARY:this line plus header is more than 75 bytes and must be wrapped ont\r\n o 2 lines"
	if actual != expected {
		t.Fatalf(`expected '%v' but got '%v'`, expected, actual)
	}
}

func TestNoCaffeineToIcs(t *testing.T) {
	ics := NoCaffeine{Start: start, End: end}.ToIcs(localization.EN)
	expected := "BEGIN:VEVENT\r\nUID:20220102T030000ZNoCaffeine@timegiver.app\r\nDTSTAMP:20220102T030000Z\r\nDTSTART:20220102T030000Z\r\nDTEND:20220102T043000Z\r\nSUMMARY:No caffeine\r\nDESCRIPTION:Refrain from consuming caffeine.  In addition to coffee and tea\r\n , caffeine may also be present in soft drinks and chocolate.\r\nCATEGORIES:TimeGiver\r\nEND:VEVENT\r\n"
	if ics != expected {
		t.Fatalf("expected\n\n%v\n\nbut got\n\n%v", expected, ics)
	}
	t.Log(ics)
}

func TestCaffeineOkToIcs(t *testing.T) {
	ics := CaffeineOk{Start: start, End: end}.ToIcs(localization.HI)
	if !strings.Contains(ics, `CaffeineOk`) {
		t.Fatalf(`ics did not have the expected step ID`)
	}
	if !strings.Contains(ics, localization.CaffeineOk[localization.HI].Title) {
		t.Fatalf(`ics does not contain expected localized title`)
	}
	t.Log(ics)
}

func TestCaffeine3CToIcs(t *testing.T) {
	ics := Caffeine3C{Start: start, End: end}.ToIcs(localization.PT)
	if !strings.Contains(ics, `Caffeine3C`) {
		t.Fatalf(`ics did not have the expected step ID`)
	}
	if !strings.Contains(ics, localization.Caffeine3C[localization.PT].Title) {
		t.Fatalf(`ics does not contain expected localized title`)
	}
	t.Log(ics)
}

func TestCaffeine2CToIcs(t *testing.T) {
	ics := Caffeine2C{Start: start, End: end}.ToIcs(localization.DE)
	if !strings.Contains(ics, `Caffeine2C`) {
		t.Fatalf(`ics did not have the expected step ID`)
	}
	if !strings.Contains(ics, localization.Caffeine2C[localization.DE].Title) {
		t.Fatalf(`ics does not contain expected localized title`)
	}
	t.Log(ics)
}

func TestLightBreakfastToIcs(t *testing.T) {
	ics := LightBreakfast{At: start}.ToIcs(localization.ES)
	if !strings.Contains(ics, `LightBreakfast`) {
		t.Fatalf(`ics did not have the expected step ID`)
	}
	if !strings.Contains(ics, localization.LightBreakfast[localization.ES].Title) {
		t.Fatalf(`ics does not contain expected localized title`)
	}
	t.Log(ics)
}

func TestLightLunchToIcs(t *testing.T) {
	ics := LightLunch{At: start}.ToIcs(localization.FR)
	if !strings.Contains(ics, `LightLunch`) {
		t.Fatalf(`ics did not have the expected step ID`)
	}
	if !strings.Contains(ics, localization.LightLunch[localization.FR].Title) {
		t.Fatalf(`ics does not contain expected localized title`)
	}
	t.Log(ics)
}

func TestLightDinnerToIcs(t *testing.T) {
	ics := LightDinner{At: start}.ToIcs(localization.JA)
	if !strings.Contains(ics, `LightDinner`) {
		t.Fatalf(`ics did not have the expected step ID`)
	}
	if !strings.Contains(ics, localization.LightDinner[localization.JA].Title) {
		t.Fatalf(`ics does not contain expected localized title`)
	}
	t.Log(ics)
}

func TestLightDinnerOptionalToIcs(t *testing.T) {
	ics := LightDinnerOptional{At: start}.ToIcs(localization.RU)
	if !strings.Contains(ics, `LightDinnerOptional`) {
		t.Fatalf(`ics did not have the expected step ID`)
	}
	if !strings.Contains(ics, localization.LightDinnerOptional[localization.RU].Title) {
		t.Fatalf(`ics does not contain expected localized title`)
	}
	t.Log(ics)
}

func TestHeavyBreakfastToIcs(t *testing.T) {
	ics := HeavyBreakfast{At: start}.ToIcs(localization.ES)
	if !strings.Contains(ics, `HeavyBreakfast`) {
		t.Fatalf(`ics did not have the expected step ID`)
	}
	if !strings.Contains(ics, localization.HeavyBreakfast[localization.ES].Title) {
		t.Fatalf(`ics does not contain expected localized title`)
	}
	t.Log(ics)
}

func TestHeavyLunchToIcs(t *testing.T) {
	ics := HeavyLunch{At: start}.ToIcs(localization.FR)
	if !strings.Contains(ics, `HeavyLunch`) {
		t.Fatalf(`ics did not have the expected step ID`)
	}
	if !strings.Contains(ics, localization.HeavyLunch[localization.FR].Title) {
		t.Fatalf(`ics does not contain expected localized title`)
	}
	t.Log(ics)
}

func TestHeavyDinnerToIcs(t *testing.T) {
	ics := HeavyDinner{At: start}.ToIcs(localization.JA)
	if !strings.Contains(ics, `HeavyDinner`) {
		t.Fatalf(`ics did not have the expected step ID`)
	}
	if !strings.Contains(ics, localization.HeavyDinner[localization.JA].Title) {
		t.Fatalf(`ics does not contain expected localized title`)
	}
	t.Log(ics)
}

func TestNoSnackToIcs(t *testing.T) {
	ics := NoSnack{Start: start, End: end}.ToIcs(localization.ZH)
	if !strings.Contains(ics, `NoSnack`) {
		t.Fatalf(`ics did not have the expected step ID`)
	}
	if !strings.Contains(ics, localization.NoSnack[localization.ZH].Title) {
		t.Fatalf(`ics does not contain expected localized title`)
	}
	t.Log(ics)
}

func TestSleepToIcs(t *testing.T) {
	ics := Sleep{Start: start, End: end}.ToIcs(localization.EN)
	if !strings.Contains(ics, `Sleep`) {
		t.Fatalf(`ics did not have the expected step ID`)
	}
	if !strings.Contains(ics, localization.Sleep[localization.EN].Title) {
		t.Fatalf(`ics does not contain expected localized title`)
	}
	t.Log(ics)
}

func TestNoNapToIcs(t *testing.T) {
	ics := NoNap{Start: start, End: end}.ToIcs(localization.EN)
	if !strings.Contains(ics, `NoNap`) {
		t.Fatalf(`ics did not have the expected step ID`)
	}
	if !strings.Contains(ics, localization.NoNap[localization.EN].Title) {
		t.Fatalf(`ics does not contain expected localized title`)
	}
	t.Log(ics)
}

func TestSetWatchToIcs(t *testing.T) {
	ics := SetWatch{At: start}.ToIcs(localization.EN)
	if !strings.Contains(ics, `SetWatch`) {
		t.Fatalf(`ics did not have the expected step ID`)
	}
	if !strings.Contains(ics, localization.SetWatch[localization.EN].Title) {
		t.Fatalf(`ics does not contain expected localized title`)
	}
	t.Log(ics)
}

func TestArriveToIcs(t *testing.T) {
	ics := Arrive{At: start}.ToIcs(localization.EN)
	if !strings.Contains(ics, `Arrive`) {
		t.Fatalf(`ics did not have the expected step ID`)
	}
	if !strings.Contains(ics, localization.Arrive[localization.EN].Title) {
		t.Fatalf(`ics does not contain expected localized title`)
	}
	t.Log(ics)
}
