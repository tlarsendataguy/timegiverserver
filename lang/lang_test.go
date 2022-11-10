package lang

import (
	"fmt"
	"testing"
)

func TestAmericanEnglish(t *testing.T) {
	err := checkParse(`en-US`, EN)
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
