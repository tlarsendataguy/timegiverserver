package localization

import (
	"testing"
)

func TestStopCaffeine(t *testing.T) {
	value := StopCaffeine[JA]
	if value.Title != `カフェインなし` {
		t.Fatalf(`expected 'カフェインなし' but got '%v'`, value.Title)
	}
}

func TestCaffeineOk(t *testing.T) {
	value := CaffeineOk[DE]
	if value.Title != `Koffein Ok` {
		t.Fatalf(`expected 'Koffein Ok' but got '%v'`, value.Title)
	}
}

func TestCaffeine12(t *testing.T) {
	value := Caffeine12[EN]
	if value.Title != `Drink a caffeinated beverage` {
		t.Fatalf(`expected 'Drink a caffeinated beverage' but got '%v'`, value.Title)
	}
	t.Log(value.Description)
}
