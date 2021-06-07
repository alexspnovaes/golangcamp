package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestIntMinMultipleTestCases(t *testing.T) {
	var testsCases = []struct {
		want, a, b int
	}{
		{0, 10, 0},
		{0, 0, 10},
		{-2, -2, 2},
		{-2, 2, -2},
	}

	for _, tt := range testsCases {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func TestIntMaxAssertsWithMinValue(t *testing.T) {

	var want int = 2
	var min int = 2
	var max int = 10

	got := IntMin(min, max)

	// assert equality
	assert.Equal(t, want, got, "they should be equal")

}

func TestIntMaxAssertsWithMaxValue(t *testing.T) {

	var want int = 2
	var min int = 2
	var max int = 2

	got := IntMax(min, max)

	// assert equality
	assert.Equal(t, want, got, "they should be equal")

	want = 3
	min = 1
	got = IntMax(min, max)

	assert.Equal(t, want, got, "they should be equal")

	// assert inequality
	assert.NotEqual(t, want, max, "they should not be equal")

}
