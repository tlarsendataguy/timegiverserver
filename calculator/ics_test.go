package calculator

import (
	"io/fs"
	"os"
	"testing"
	"time"
	"timegiverserver/calculator/steps"
	"timegiverserver/lang"
)

func TestIcs(t *testing.T) {
	plan := []steps.Step{
		steps.Arrive{At: time.Date(2022, 1, 2, 3, 30, 0, 0, time.UTC)},
	}
	attachment := BuildIcsFile(plan, lang.EN)
	expected := "BEGIN:VCALENDAR\r\nVERSION:2.0\r\nPRODID://tlarsendataguy//timegiver\r\nBEGIN:VEVENT\r\nUID:20220102T033000ZArrive@timegiver.app\r\nDTSTAMP:20220102T033000Z\r\nDTSTART:20220102T033000Z\r\nDTEND:20220102T033000Z\r\nSUMMARY:Arrival time\r\nDESCRIPTION:If your travel arrangements are running on time, you should be \r\n arriving at your destination.\r\nCATEGORIES:TimeGiver\r\nEND:VEVENT\r\nEND:VCALENDAR\r\n"
	if attachment != expected {
		t.Fatalf("expected\n\n%v\n\nbut got\n\n%v", expected, attachment)
	}
	_ = os.WriteFile(`test_ics.ics`, []byte(attachment), fs.ModePerm)
}
