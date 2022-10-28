package shift

import "testing"

func TestEastboundNegativeToNegativeOffset(t *testing.T) {
	var result = CalcTimezoneShift(-6, -5)
	checkResult(t, 1, result)
}

func TestEastboundPositiveToPositiveOffset(t *testing.T) {
	var result = CalcTimezoneShift(5, 6)
	checkResult(t, 1, result)
}

func TestWestboundNegativeToNegativeOffset(t *testing.T) {
	var result = CalcTimezoneShift(-5, -6)
	checkResult(t, -1, result)
}

func TestWestboundPositiveToPositiveOffset(t *testing.T) {
	var result = CalcTimezoneShift(6, 5)
	checkResult(t, -1, result)
}

func TestEastboundNegativeToPositiveOffset(t *testing.T) {
	var result = CalcTimezoneShift(-6, 5)
	checkResult(t, 11, result)
}

func TestWestboundPositiveToNegativeOffset(t *testing.T) {
	var result = CalcTimezoneShift(5, -6)
	checkResult(t, -11, result)
}

func TestEastboundPositiveToNegativeOffset(t *testing.T) {
	var result = CalcTimezoneShift(10, -10)
	checkResult(t, 4, result)
}

func TestWestboundNegativeToPositiveOffset(t *testing.T) {
	var result = CalcTimezoneShift(-10, 10)
	checkResult(t, -4, result)
}

func TestWestboundSuperToPositiveOffset(t *testing.T) {
	var result = CalcTimezoneShift(13, 12)
	checkResult(t, -1, result)
}

func TestEastboundPositiveToSuperOffset(t *testing.T) {
	var result = CalcTimezoneShift(12, 13)
	checkResult(t, 1, result)
}

func TestZeroShiftInternationalDateLine(t *testing.T) {
	var result = CalcTimezoneShift(-12, 12)
	checkResult(t, 0, result)
}

func TestZeroShiftInternationalDateLinePlus1(t *testing.T) {
	var result = CalcTimezoneShift(13, -11)
	checkResult(t, 0, result)
}

func TestZeroShiftPlusSix(t *testing.T) {
	var result = CalcTimezoneShift(6, 6)
	checkResult(t, 0, result)
}

func TestZeroShiftNegativeSix(t *testing.T) {
	var result = CalcTimezoneShift(-6, -6)
	checkResult(t, 0, result)
}

func TestEastbound12(t *testing.T) {
	var result = CalcTimezoneShift(0, 12)
	checkResult(t, 12, result)
}

func TestWestbound12(t *testing.T) {
	var result = CalcTimezoneShift(0, -12)
	checkResult(t, -12, result)
}

func TestEastbound5Hour30Minute(t *testing.T) {
	var result = CalcTimezoneShift(0, 5.5)
	checkResult(t, 5, result)
}

func TestWestbound5Hour30Minute(t *testing.T) {
	var result = CalcTimezoneShift(0, -5.5)
	checkResult(t, -5, result)
}

func checkResult(t *testing.T, expected int, actual int) {
	if actual != expected {
		t.Fatalf(`expected %v but got %v`, expected, actual)
	}
}
