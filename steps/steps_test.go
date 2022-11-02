package steps

import (
	"testing"
	"time"
	"timegiverserver/steps/localization"
)

var start = time.Date(2022, 1, 2, 3, 0, 12, 0, time.UTC)
var end = time.Date(2022, 1, 2, 4, 30, 0, 0, time.UTC)

func TestDateTimeFormatting(t *testing.T) {
	actual := FmtDt(start)
	expected := `20220102T030000Z`
	if actual != expected {
		t.Fatalf(`expected %v but got %v`, expected, actual)
	}
}

func TestWrapShort(t *testing.T) {
	actual := Wrap(`SUMMARY:`, `hello world`)
	expected := `SUMMARY:hello world`
	if actual != expected {
		t.Fatalf(`expected '%v' but got '%v'`, expected, actual)
	}
}

func TestWrap2Lines(t *testing.T) {
	actual := Wrap(`SUMMARY:`, `this line plus header is more than 75 bytes and must be wrapped onto 2 lines`)
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
