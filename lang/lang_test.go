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

func checkParse(value string, expected Lang) error {
	actual, err := ParseLang(value)
	if err != nil {
		return fmt.Errorf(`expected no error but got: %v`, err.Error())
	}
	if actual != expected {
		return fmt.Errorf(`expected %v but got %v`, expected, actual)
	}
	return nil
}
