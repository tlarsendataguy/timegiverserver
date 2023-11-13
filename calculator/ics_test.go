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
		steps.SetWatch{At: time.Date(2022, 1, 2, 3, 30, 0, 0, time.UTC)},
	}
	attachment := BuildIcsFile(plan, lang.EN)
	expected := "BEGIN:VCALENDAR\r\nVERSION:2.0\r\nPRODID://tlarsendataguy//timegiver\r\nX-WR-CALNAME: Timegiver\r\nNAME: Timegiver\r\nBEGIN:VEVENT\r\nUID:20220102T033000ZSetWatch@timegiver.app\r\nDTSTAMP:20220102T033000Z\r\nDTSTART:20220102T033000Z\r\nDTEND:20220102T033000Z\r\nSUMMARY:Set watch\r\nDESCRIPTION:Set your watch to destination time.  You are now making the swi\r\n tch to organize your day around the destination timezone.\r\nBEGIN:VALARM\r\nTRIGGER;RELATED=START:PT0M\r\nACTION:DISPLAY\r\nDESCRIPTION:Set watch\r\nEND:VALARM\r\nCATEGORIES:TimeGiver\r\nEND:VEVENT\r\nEND:VCALENDAR\r\n"
	if attachment != expected {
		t.Fatalf("expected\n\n%v\n\nbut got\n\n%v", expected, attachment)
	}
	_ = os.WriteFile(`test_ics.ics`, []byte(attachment), fs.ModePerm)
}
