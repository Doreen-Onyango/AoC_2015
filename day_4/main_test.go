package main

import "testing"

func TestLowN(t *testing.T) {
	got := LowNum("abcdef")
	want := 609043
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
