package main

import "testing"

func TestStrStatus(t *testing.T) {
	got := StrStatus("ugknbfddgicrmopn")
	want := "nice"
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
