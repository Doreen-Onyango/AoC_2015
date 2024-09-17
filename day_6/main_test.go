package main

import (
	"os"
	"testing"
)

func TestReadInstructions(t *testing.T) {
	testContent := "turn on 0,0 through 999,999\n" +
		"toggle 0,0 through 999,0\n" +
		"turn off 499,499 through 500,500\n"

	file, err := os.CreateTemp("", "testfile_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(file.Name())

	_, err = file.WriteString(testContent)
	if err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	file.Close()

	instructions := ReadInstructions(file.Name())

	expected := []string{
		"turn on 0,0 through 999,999",
		"toggle 0,0 through 999,0",
		"turn off 499,499 through 500,500",
	}

	if len(instructions) != len(expected) {
		t.Fatalf("Expected %d instructions, got %d", len(expected), len(instructions))
	}

	for i := range expected {
		if instructions[i] != expected[i] {
			t.Errorf("Instruction at index %d: expected %q, got %q", i, expected[i], instructions[i])
		}
	}
}
