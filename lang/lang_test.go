package lang

import "testing"

func TestAmericanEnglish(t *testing.T) {
	lang, err := ParseLang(`en-US`)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if lang != EN {
		t.Fatalf(`expected %v but got %v`, EN, lang)
	}
}
