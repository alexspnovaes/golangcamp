package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testFilterAgeByRange(t *testing.T) {

	ages := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	var ageMin int = 0
	var ageMax int = 10

	var expectedCount int = 11

	want := filterByAgeRange(ageMin, ageMax, ages)

	// assert equality
	assert.Equal(t, want[0], ageMin, "they should be equal")
	assert.Equal(t, want[len(want)-1], ageMax, "they should be equal")
	assert.Equal(t, len(want), expectedCount, "they should be equal")

	assert.NotEmpty(t, want, "it should not empty")
	assert.NotNil(t, want, "it shout not be null")

}
