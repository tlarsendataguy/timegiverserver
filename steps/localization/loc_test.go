package localization

import "testing"

func TestLocCompleteness(t *testing.T) {
	expectedQty := 9

	if qty := len(NoCaffeine); qty != expectedQty {
		t.Fatalf(`required %v translations but got %v`, expectedQty, qty)
	}
	if qty := len(Caffeine2C); qty != expectedQty {
		t.Fatalf(`required %v translations but got %v`, expectedQty, qty)
	}
	if qty := len(Caffeine3C); qty != expectedQty {
		t.Fatalf(`required %v translations but got %v`, expectedQty, qty)
	}
	if qty := len(CaffeineOk); qty != expectedQty {
		t.Fatalf(`required %v translations but got %v`, expectedQty, qty)
	}
	if qty := len(LightBreakfast); qty != expectedQty {
		t.Fatalf(`required %v translations but got %v`, expectedQty, qty)
	}
	if qty := len(LightLunch); qty != expectedQty {
		t.Fatalf(`required %v translations but got %v`, expectedQty, qty)
	}
	if qty := len(LightDinner); qty != expectedQty {
		t.Fatalf(`required %v translations but got %v`, expectedQty, qty)
	}
	if qty := len(LightDinnerOptional); qty != expectedQty {
		t.Fatalf(`required %v translations but got %v`, expectedQty, qty)
	}
	if qty := len(HeavyBreakfast); qty != expectedQty {
		t.Fatalf(`required %v translations but got %v`, expectedQty, qty)
	}
	if qty := len(HeavyLunch); qty != expectedQty {
		t.Fatalf(`required %v translations but got %v`, expectedQty, qty)
	}
	if qty := len(HeavyDinner); qty != expectedQty {
		t.Fatalf(`required %v translations but got %v`, expectedQty, qty)
	}
	if qty := len(NoSnack); qty != expectedQty {
		t.Fatalf(`required %v translations but got %v`, expectedQty, qty)
	}
	if qty := len(SetWatch); qty != expectedQty {
		t.Fatalf(`required %v translations but got %v`, expectedQty, qty)
	}
	if qty := len(Sleep); qty != expectedQty {
		t.Fatalf(`required %v translations but got %v`, expectedQty, qty)
	}
	if qty := len(NoNap); qty != expectedQty {
		t.Fatalf(`required %v translations but got %v`, expectedQty, qty)
	}
	if qty := len(Arrive); qty != expectedQty {
		t.Fatalf(`required %v translations but got %v`, expectedQty, qty)
	}
}
