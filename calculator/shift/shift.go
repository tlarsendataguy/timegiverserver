package shift

import "math"

func CalcTimezoneShift(startingOffset, endingOffset float64) int {
	shift := int(math.Trunc(endingOffset - startingOffset))
	if shift < -12 {
		return shift + 24
	}
	if shift > 12 {
		return shift - 24
	}
	return shift
}
