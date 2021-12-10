package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_loadSubCommands(t *testing.T) {
	want := []SubCommand{
		{
			Type:   "forward",
			Amount: 5,
		},
		{
			Type:   "down",
			Amount: 5,
		},
		{
			Type:   "forward",
			Amount: 8,
		},
		{
			Type:   "up",
			Amount: 3,
		},
		{
			Type:   "down",
			Amount: 8,
		},
		{
			Type:   "forward",
			Amount: 2,
		},
	}
	actual, err := loadSubCommands("test_input.txt")
	if err != nil {
		t.Errorf("error running loadSubCommands func: %v", err)
	}
	assert.Equal(t, want, actual)
}

func Test_solvePart1(t *testing.T) {
	want := 150
	input := []SubCommand{
		{
			Type:   "forward",
			Amount: 5,
		},
		{
			Type:   "down",
			Amount: 5,
		},
		{
			Type:   "forward",
			Amount: 8,
		},
		{
			Type:   "up",
			Amount: 3,
		},
		{
			Type:   "down",
			Amount: 8,
		},
		{
			Type:   "forward",
			Amount: 2,
		},
	}
	actual := solvePart1(input)
	assert.Equal(t, want, actual)
}

func Test_solvePart2(t *testing.T) {
	want := 900
	input := []SubCommand{
		{
			Type:   "forward",
			Amount: 5,
		},
		{
			Type:   "down",
			Amount: 5,
		},
		{
			Type:   "forward",
			Amount: 8,
		},
		{
			Type:   "up",
			Amount: 3,
		},
		{
			Type:   "down",
			Amount: 8,
		},
		{
			Type:   "forward",
			Amount: 2,
		},
	}
	actual := solvePart2(input)
	assert.Equal(t, want, actual)
}
