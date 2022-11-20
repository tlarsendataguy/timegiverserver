package steps

import (
	"strings"
	"testing"
	"time"
	"timegiverserver/calculator/steps/localization"
	"timegiverserver/lang"
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

func getIcsString(step Step, lang lang.Lang) string {
	builder := &strings.Builder{}
	step.BuildIcs(builder, lang)
	return builder.String()
}

func TestNoCaffeineToIcs(t *testing.T) {
	ics := getIcsString(NoCaffeine{Start: start, End: end}, lang.EN)
	expected := "BEGIN:VEVENT\r\nUID:20220102T030000ZNoCaffeine@timegiver.app\r\nDTSTAMP:20220102T030000Z\r\nDTSTART:20220102T030000Z\r\nDTEND:20220102T043000Z\r\nSUMMARY:No caffeine\r\nDESCRIPTION:Refrain from consuming caffeine.  In addition to coffee and tea\r\n , caffeine may also be present in soft drinks and chocolate.\r\nBEGIN:VALARM\r\nTRIGGER;RELATED=START:PT0M\r\nACTION:DISPLAY\r\nDESCRIPTION:No caffeine\r\nEND:VALARM\r\nCATEGORIES:TimeGiver\r\nEND:VEVENT\r\n"
	if ics != expected {
		t.Fatalf("expected\n\n%v\n\nbut got\n\n%v", expected, ics)
	}
	t.Log(ics)
}

func TestCaffeineOkToIcs(t *testing.T) {
	ics := getIcsString(CaffeineOk{Start: start, End: end}, lang.HI)
	if !strings.Contains(ics, `CaffeineOk`) {
		t.Fatalf(`ics did not have the expected step ID`)
	}
	if !strings.Contains(ics, localization.CaffeineOk[lang.HI].Title) {
		t.Fatalf(`ics does not contain expected localized title`)
	}
	t.Log(ics)
}

func TestCaffeine3CToIcs(t *testing.T) {
	ics := getIcsString(Caffeine3C{Start: start, End: end}, lang.PT)
	if !strings.Contains(ics, `Caffeine3C`) {
		t.Fatalf(`ics did not have the expected step ID`)
	}
	if !strings.Contains(ics, localization.Caffeine3C[lang.PT].Title) {
		t.Fatalf(`ics does not contain expected localized title`)
	}
	t.Log(ics)
}

func TestCaffeine2CToIcs(t *testing.T) {
	ics := getIcsString(Caffeine2C{Start: start, End: end}, lang.DE)
	if !strings.Contains(ics, `Caffeine2C`) {
		t.Fatalf(`ics did not have the expected step ID`)
	}
	if !strings.Contains(ics, localization.Caffeine2C[lang.DE].Title) {
		t.Fatalf(`ics does not contain expected localized title`)
	}
	t.Log(ics)
}

func TestLightBreakfastToIcs(t *testing.T) {
	ics := getIcsString(LightBreakfast{At: start}, lang.ES)
	if !strings.Contains(ics, `LightBreakfast`) {
		t.Fatalf(`ics did not have the expected step ID`)
	}
	if !strings.Contains(ics, localization.LightBreakfast[lang.ES].Title) {
		t.Fatalf(`ics does not contain expected localized title`)
	}
	t.Log(ics)
}

func TestLightLunchToIcs(t *testing.T) {
	ics := getIcsString(LightLunch{At: start}, lang.FR)
	if !strings.Contains(ics, `LightLunch`) {
		t.Fatalf(`ics did not have the expected step ID`)
	}
	if !strings.Contains(ics, localization.LightLunch[lang.FR].Title) {
		t.Fatalf(`ics does not contain expected localized title`)
	}
	t.Log(ics)
}

func TestLightDinnerToIcs(t *testing.T) {
	ics := getIcsString(LightDinner{At: start}, lang.JA)
	if !strings.Contains(ics, `LightDinner`) {
		t.Fatalf(`ics did not have the expected step ID`)
	}
	if !strings.Contains(ics, localization.LightDinner[lang.JA].Title) {
		t.Fatalf(`ics does not contain expected localized title`)
	}
	t.Log(ics)
}

func TestLightDinnerOptionalToIcs(t *testing.T) {
	ics := getIcsString(LightDinnerOptional{At: start}, lang.RU)
	if !strings.Contains(ics, `LightDinnerOptional`) {
		t.Fatalf(`ics did not have the expected step ID`)
	}
	if !strings.Contains(ics, localization.LightDinnerOptional[lang.RU].Title) {
		t.Fatalf(`ics does not contain expected localized title`)
	}
	t.Log(ics)
}

func TestHeavyBreakfastToIcs(t *testing.T) {
	ics := getIcsString(HeavyBreakfast{At: start}, lang.ES)
	if !strings.Contains(ics, `HeavyBreakfast`) {
		t.Fatalf(`ics did not have the expected step ID`)
	}
	if !strings.Contains(ics, localization.HeavyBreakfast[lang.ES].Title) {
		t.Fatalf(`ics does not contain expected localized title`)
	}
	t.Log(ics)
}

func TestHeavyLunchToIcs(t *testing.T) {
	ics := getIcsString(HeavyLunch{At: start}, lang.FR)
	if !strings.Contains(ics, `HeavyLunch`) {
		t.Fatalf(`ics did not have the expected step ID`)
	}
	if !strings.Contains(ics, localization.HeavyLunch[lang.FR].Title) {
		t.Fatalf(`ics does not contain expected localized title`)
	}
	t.Log(ics)
}

func TestHeavyDinnerToIcs(t *testing.T) {
	ics := getIcsString(HeavyDinner{At: start}, lang.JA)
	if !strings.Contains(ics, `HeavyDinner`) {
		t.Fatalf(`ics did not have the expected step ID`)
	}
	if !strings.Contains(ics, localization.HeavyDinner[lang.JA].Title) {
		t.Fatalf(`ics does not contain expected localized title`)
	}
	t.Log(ics)
}

func TestNoSnackToIcs(t *testing.T) {
	ics := getIcsString(NoSnack{Start: start, End: end}, lang.ZH)
	if !strings.Contains(ics, `NoSnack`) {
		t.Fatalf(`ics did not have the expected step ID`)
	}
	if !strings.Contains(ics, localization.NoSnack[lang.ZH].Title) {
		t.Fatalf(`ics does not contain expected localized title`)
	}
	t.Log(ics)
}

func TestSleepToIcs(t *testing.T) {
	ics := getIcsString(Sleep{Start: start, End: end}, lang.EN)
	if !strings.Contains(ics, `Sleep`) {
		t.Fatalf(`ics did not have the expected step ID`)
	}
	if !strings.Contains(ics, localization.Sleep[lang.EN].Title) {
		t.Fatalf(`ics does not contain expected localized title`)
	}
	t.Log(ics)
}

func TestNoNapToIcs(t *testing.T) {
	ics := getIcsString(NoNap{Start: start, End: end}, lang.EN)
	if !strings.Contains(ics, `NoNap`) {
		t.Fatalf(`ics did not have the expected step ID`)
	}
	if !strings.Contains(ics, localization.NoNap[lang.EN].Title) {
		t.Fatalf(`ics does not contain expected localized title`)
	}
	t.Log(ics)
}

func TestSetWatchToIcs(t *testing.T) {
	ics := getIcsString(SetWatch{At: start}, lang.EN)
	if !strings.Contains(ics, `SetWatch`) {
		t.Fatalf(`ics did not have the expected step ID`)
	}
	if !strings.Contains(ics, localization.SetWatch[lang.EN].Title) {
		t.Fatalf(`ics does not contain expected localized title`)
	}
	t.Log(ics)
}

func TestArriveToIcs(t *testing.T) {
	ics := getIcsString(Arrive{At: start}, lang.EN)
	if !strings.Contains(ics, `Arrive`) {
		t.Fatalf(`ics did not have the expected step ID`)
	}
	if !strings.Contains(ics, localization.Arrive[lang.EN].Title) {
		t.Fatalf(`ics does not contain expected localized title`)
	}
	t.Log(ics)
}
