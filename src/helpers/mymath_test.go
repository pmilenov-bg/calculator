package mymath

import (
	"testing"
)

func TestAdd(t *testing.T) {
	a := 6
	b := 3
	expected := 9

	result := Add(a, b)

	if result != expected {
		t.Errorf("Addition failed. Expected: %d, Got: %d", expected, result)
	}
}

func TestSubtract(t *testing.T) {
	a := 6
	b := 3
	expected := 3

	result := Subtract(a, b)

	if result != expected {
		t.Errorf("Subtraction failed. Expected: %d, Got: %d", expected, result)
	}
}

func TestMultiply(t *testing.T) {
	a := 6
	b := 3
	expected := 18

	result := Multiply(a, b)

	if result != expected {
		t.Errorf("Multiplication failed. Expected: %d, Got: %d", expected, result)
	}
}

func TestDivide(t *testing.T) {
	a := 6
	b := 3
	expected := 2

	result := Divide(a, b)

	if result != expected {
		t.Errorf("Division failed. Expected: %d, Got: %d", expected, result)
	}
}

func TestMain(t *testing.T) {
	// Perform any additional tests for the main function if necessary
}

func TestAll(t *testing.T) {
	// Run all the tests
	t.Run("TestAdd", TestAdd)
	t.Run("TestSubtract", TestSubtract)
	t.Run("TestMultiply", TestMultiply)
	t.Run("TestDivide", TestDivide)
	t.Run("TestMain", TestMain)
}
