package main_test

import (
	"os"
	"testing"

	day3 "github.com/chigley/advent2018/cmd/day3"
	"github.com/stretchr/testify/assert"
)

func TestDay1(t *testing.T) {
	inputFile, err := os.Open("testdata/input")
	if err != nil {
		t.Fatal(err)
	}
	defer inputFile.Close()

	claims, err := day3.ReadClaims(inputFile)
	if err != nil {
		t.Fatal(err)
	}

	part1, err := day3.Part1(claims)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 104126, part1)
}

func TestPart1(t *testing.T) {
	input := []day3.Claim{
		{ID: 1, OffsetX: 1, OffsetY: 3, Width: 4, Height: 4},
		{ID: 2, OffsetX: 3, OffsetY: 1, Width: 4, Height: 4},
		{ID: 3, OffsetX: 5, OffsetY: 5, Width: 2, Height: 2},
	}
	output, err := day3.Part1(input)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 4, output)
}
