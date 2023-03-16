package main

import "testing"

func TestAddNumbers(t *testing.T) {
	result := addNumbers(1, 2)
	if result != 3 {
		t.Errorf("addNumbers(1, 2) = %d; want 3", result)
	}
}

func TestSubtractNumbers(t *testing.T) {
	result := subtractNumbers(1, 2)
	if result != -1 {
		t.Errorf("subtractNumbers(1, 2) = %d; want -1", result)
	}
}

func TestMultiplyNumbers(t *testing.T) {
	result := multiplyNumbers(1, 2)
	if result != 2 {
		t.Errorf("multiplyNumbers(1, 2) = %d; want 2", result)
	}
}
