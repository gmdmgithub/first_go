package main

import (
	"testing"
)

// TestMain - Testing main func
func TestCalculate(t *testing.T) {
	if Calculate(8) != 32 {
		t.Error("Expected 32 not get!!")
	}
}

func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 8},
		{-1, -4},
		{0, 1},
		{-2, -8},
		{1, 4},
	}

	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}
