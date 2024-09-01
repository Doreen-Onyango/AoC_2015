package main

import "testing"

func TestPosition(t *testing.T) {
	got := Position("^>v<")
	want := 4

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestRoboSanta(t *testing.T) {
	got := RoboSanta("^v^v^v^v^v")
	want := 11

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
