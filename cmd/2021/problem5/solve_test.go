package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getLine(t *testing.T) {
	cases := []struct {
		name   string
		want   []Point
		input1 Point
		input2 Point
	}{
		{
			name: "1,1-1,3",
			want: []Point{
				{X: 1, Y: 1},
				{X: 1, Y: 2},
				{X: 1, Y: 3},
			},
			input1: Point{X: 1, Y: 1},
			input2: Point{X: 1, Y: 3},
		},
		{
			name: "9,7-7,7",
			want: []Point{
				{X: 7, Y: 7},
				{X: 8, Y: 7},
				{X: 9, Y: 7},
			},
			input1: Point{X: 9, Y: 7},
			input2: Point{X: 7, Y: 7},
		},
		{
			name: "1,3-3,3",
			want: []Point{
				{X: 1, Y: 1},
				{X: 2, Y: 2},
				{X: 3, Y: 3},
			},
			input1: Point{X: 1, Y: 3},
			input2: Point{X: 3, Y: 3},
		},
		{
			name: "9,7-7,9",
			want: []Point{
				{X: 7, Y: 9},
				{X: 8, Y: 8},
				{X: 9, Y: 7},
			},
			input1: Point{X: 9, Y: 7},
			input2: Point{X: 7, Y: 9},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := getLine(c.input1, c.input2)
			assert.Equal(t, c.want, actual)
		})
	}

}

func Test_diff(t *testing.T) {
	cases := []struct {
		name   string
		want   []int
		input1 int
		input2 int
	}{
		{
			name:   "1-3",
			want:   []int{1, 2, 3},
			input1: 1,
			input2: 3,
		},
		{
			name:   "9-7",
			want:   []int{7, 8, 9},
			input1: 9,
			input2: 7,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := diff(c.input1, c.input2)
			assert.Equal(t, c.want, actual)
		})
	}
}

func Test_loadInput(t *testing.T) {
	want1 := []Point{
		{X: 0, Y: 9},
		{X: 8, Y: 0},
		{X: 9, Y: 4},
		{X: 2, Y: 2},
		{X: 7, Y: 0},
		{X: 6, Y: 4},
		{X: 0, Y: 9},
		{X: 3, Y: 4},
		{X: 0, Y: 0},
		{X: 5, Y: 5},
	}
	want2 := []Point{
		{X: 5, Y: 9},
		{X: 0, Y: 8},
		{X: 3, Y: 4},
		{X: 2, Y: 1},
		{X: 7, Y: 4},
		{X: 2, Y: 0},
		{X: 2, Y: 9},
		{X: 1, Y: 4},
		{X: 8, Y: 8},
		{X: 8, Y: 2},
	}
	act1, act2, err := loadInput("test_input.txt")
	if err != nil {
		t.Fatalf("error running loadInput func: %v", err)
	}
	assert.Equal(t, want1, act1)
	assert.Equal(t, want2, act2)
}

func Test_solvePart1(t *testing.T) {
	p1s := []Point{
		{X: 0, Y: 9},
		{X: 8, Y: 0},
		{X: 9, Y: 4},
		{X: 2, Y: 2},
		{X: 7, Y: 0},
		{X: 6, Y: 4},
		{X: 0, Y: 9},
		{X: 3, Y: 4},
		{X: 0, Y: 0},
		{X: 5, Y: 5},
	}
	p2s := []Point{
		{X: 5, Y: 9},
		{X: 0, Y: 8},
		{X: 3, Y: 4},
		{X: 2, Y: 1},
		{X: 7, Y: 4},
		{X: 2, Y: 0},
		{X: 2, Y: 9},
		{X: 1, Y: 4},
		{X: 8, Y: 8},
		{X: 8, Y: 2},
	}
	actual := solvePart1(p1s, p2s)
	assert.Equal(t, 5, actual)
}

func Test_filterLines(t *testing.T) {
	p1s := []Point{
		{X: 0, Y: 9},
		{X: 8, Y: 0},
		{X: 9, Y: 4},
		{X: 2, Y: 2},
		{X: 7, Y: 0},
		{X: 6, Y: 4},
		{X: 0, Y: 9},
		{X: 3, Y: 4},
		{X: 0, Y: 0},
		{X: 5, Y: 5},
	}
	p2s := []Point{
		{X: 5, Y: 9},
		{X: 0, Y: 8},
		{X: 3, Y: 4},
		{X: 2, Y: 1},
		{X: 7, Y: 4},
		{X: 2, Y: 0},
		{X: 2, Y: 9},
		{X: 1, Y: 4},
		{X: 8, Y: 8},
		{X: 8, Y: 2},
	}
	want1 := []Point{
		{X: 0, Y: 9},
		{X: 9, Y: 4},
		{X: 2, Y: 2},
		{X: 7, Y: 0},
		{X: 0, Y: 9},
		{X: 3, Y: 4},
	}
	want2 := []Point{
		{X: 5, Y: 9},
		{X: 3, Y: 4},
		{X: 2, Y: 1},
		{X: 7, Y: 4},
		{X: 2, Y: 9},
		{X: 1, Y: 4},
	}
	act1, act2 := filterLines(p1s, p2s)
	assert.Equal(t, want1, act1)
	assert.Equal(t, want2, act2)
}
