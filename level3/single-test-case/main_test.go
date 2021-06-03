package main

import "testing"

func TestIntMin(t *testing.T) {
	var min = 1
	got := IntMin(min, 2)
	if got != min {
		t.Errorf("Failed, expected %v got %v", min, got)
	}
}

func TestIntMax(t *testing.T) {
	var min = 2
	var max = 10
	got := IntMin(max, min)
	if got != min {
		t.Errorf("Failed, expected %v got %v", min, got)
	}
}

func TestIntMinError(t *testing.T) {
	var min = 10
	var max = 2
	got := IntMin(min, max)
	if got != min {
		t.Errorf("Failed, expected %v got %v", min, got)
	}
}
