package main_test

import (
	"os"
	"testing"

	"github.com/chigley/advent2018"
	day1 "github.com/chigley/advent2018/cmd/day1"
	"github.com/stretchr/testify/assert"
)

var part1Tests = []struct {
	in  []int
	out int
}{
	{[]int{+1, -2, +3, +1}, 3},
	{[]int{+1, +1, +1}, 3},
	{[]int{+1, +1, -2}, 0},
	{[]int{-1, -2, -3}, -6},
}

var part2Tests = []struct {
	in  []int
	out int
}{
	{[]int{+1, -1}, 0},
	{[]int{+3, +3, +4, -2, -4}, 10},
	{[]int{-6, +3, +8, +5, -6}, 5},
	{[]int{+7, +7, -2, -7, -4}, 14},
}

func TestDay1(t *testing.T) {
	inputFile, err := os.Open("testdata/input")
	if err != nil {
		t.Fatal(err)
	}
	defer inputFile.Close()

	input, err := advent2018.ReadInts(inputFile)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 582, day1.Part1(input))
	assert.Equal(t, 488, day1.Part2(input))
}

func TestPart1(t *testing.T) {
	for _, tt := range part1Tests {
		assert.Equal(t, tt.out, day1.Part1(tt.in))
	}
}

func TestPart2(t *testing.T) {
	for _, tt := range part2Tests {
		assert.Equal(t, tt.out, day1.Part2(tt.in))
	}
}
