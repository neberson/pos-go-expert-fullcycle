package tax

import "testing"

func TestTax(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculateTax(float64(amount))

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}
