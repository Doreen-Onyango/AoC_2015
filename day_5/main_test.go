package main

import (
	"testing"
)

func TestStrStatus(t *testing.T) {
	a := []string{"ugknbfddgicrmopn", "jchzalrnumimnmhp", "haegwjzuvuyypxyu", "dvszwmarrgswjxmb", "ugknbfddgicrmopn"}
	got := NiceString(a)
	want := 2
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestThreVow(t *testing.T) {
	got := ThreVow("ugknbfddgicrmopn")
	want := true
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestContLet(t *testing.T) {
	got := ContLet("ugknbfddgicrmopn")
	want := true
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestNoStr(t *testing.T) {
	got := ContLet("ugknbfcdgicrmopn")
	want := false
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestNiceStr(t *testing.T) {
	x := []string{"qjhvhtzxzqqjkmpb", "qjhvhtzxzqqjkmpb", "qjhvhtzxzqqjkmpb", "uurcxstgmygtbstg"}
	got := NiceStr(x)
	want := 3
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestAlign(t *testing.T) {
	got := Align("qjhvhtzxzqqjkmpb")
	want := true
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestPaire(t *testing.T) {
	got := Paire("qjhvhtzxzqqjkmpb")
	want := true
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
