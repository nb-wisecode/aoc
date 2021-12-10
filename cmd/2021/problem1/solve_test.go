package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solvePart1(t *testing.T) {
	input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	want := 7
	actual := solvePart1(input)
	assert.Equal(t, want, actual)
}

func Test_solvePart2(t *testing.T) {
	input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	want := 5
	actual := solvePart2(input)
	assert.Equal(t, want, actual)
}
