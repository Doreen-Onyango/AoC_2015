package main

import "testing"

func TestLowN(t *testing.T) {
	got := LowNum("abcdef")
	want := 609043
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestLowNumber(t *testing.T) {
	got := LowNumber("abcdef")
	want := 6742839
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
