package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_loadInput(t *testing.T) {
	want := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}
	actual, err := loadInput("test_input.txt")
	if err != nil {
		t.Errorf("error running loadInput func: %v", err)
	}
	assert.Equal(t, want, actual)
}

func Test_solvePart1(t *testing.T) {
	input := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}
	want := 198
	actual, err := solvePart1(input)
	if err != nil {
		t.Errorf("error running solvePart1 func: %v", err)
	}
	assert.Equal(t, want, actual)
}
