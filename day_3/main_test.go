package main

import "testing"

func TestPresents(t *testing.T) {
	got := Position("^>v<")
	want := 4

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
