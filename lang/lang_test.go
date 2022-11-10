package lang

import (
	"fmt"
	"testing"
)

func TestLanguageOnly(t *testing.T) {
	err := checkParse(`en`, EN)
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestLanguageAndLocale(t *testing.T) {
	err := checkParse(`es-CO`, ES)
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestAnyLanguageDefaultsToEnglish(t *testing.T) {
	err := checkParse(`*`, EN)
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestPickFirstWeightedMatch(t *testing.T) {
	err := checkParse(`af-ZA,ga;q=0.9,hi;q=0.5,en;q=0.1`, HI)
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func checkParse(value string, expected Lang) error {
	actual := ParseLang(value)
	if actual != expected {
		return fmt.Errorf(`expected %v but got %v`, expected, actual)
	}
	return nil
}
