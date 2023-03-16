package main

import "testing"

func TestAddNumbers(t *testing.T) {
	result := addNumbers(1, 2)
	if result != 3 {
		t.Errorf("addNumbers(1, 2) = %d; want 3", result)
	}
}
