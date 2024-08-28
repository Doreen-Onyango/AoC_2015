package main

import "testing"

func TestSurfaceArea(t *testing.T) {
	got := SurfaceArea("25x14x4")
	want := 1068

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestPerimeter(t *testing.T) {
	got := Perimeter("2x3x4")
	want := 34

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
