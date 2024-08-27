package main

import (
	"testing"
)

func Test_Apartment(t *testing.T) {
	tests := []struct {
		instruction string
		expected    int
	}{
		{"((", 2},
		{")))", -3},
		{"(())", 0},
	}

	for _, test := range tests {
		result := Apartment(test.instruction)
		if result != test.expected {
			t.Errorf("Apartment(%q); want %d", result, test.expected)
		}
	}
}
