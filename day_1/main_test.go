package main

import "testing"

func TestFloorDirection(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"", 0},
		{"(", 1},
		{")", -1},
		{"(()(", 2},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := FloorDirection(tt.input)
			if result != tt.expected {
				t.Errorf("FloorDirection(%q) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestFloorPosition(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"()())()", 5},
		{"((((())))))", 11},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := FloorPosition(tt.input)
			if result != tt.expected {
				t.Errorf("FloorPosition(%q) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}
