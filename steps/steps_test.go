package steps

import (
	"testing"
	"time"
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
