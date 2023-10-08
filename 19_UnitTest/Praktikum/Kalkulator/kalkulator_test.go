package Kalkulator

import "testing"

func TestAddition(t *testing.T) {
	result := Addition(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Addition result was incorrect, got: %d, expected: %d.", result, expected)
	}
}

func TestSubstract(t *testing.T) {
	result := Substract(7, 3)
	expected := 4
	if result != expected {
		t.Errorf("Substract result was incorrect, got: %d, expected: %d.", result, expected)
	}
}

func TestMult(t *testing.T) {
	result := Mult(2, 3)
	expected := 6
	if result != expected {
		t.Errorf("Mult result was incorrect, got: %d, expected: %d.", result, expected)
	}
}
func TestDiv(t *testing.T) {
	result := Div(10, 5)
	if result != 2 {
		t.Errorf("Div(20, 5) = %d; want 4", result)
	}

	// testing division by zero
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Div(10, 0) did not panic")
		}
	}()
	Div(10, 0)
}