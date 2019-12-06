package main_test

import (
	"os"
	"testing"

	"github.com/chigley/advent2018"
	day2 "github.com/chigley/advent2018/cmd/day2"
	"github.com/stretchr/testify/assert"
)

func TestDay1(t *testing.T) {
	inputFile, err := os.Open("testdata/input")
	if err != nil {
		t.Fatal(err)
	}
	defer inputFile.Close()

	input, err := advent2018.ReadStrings(inputFile)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 7470, day2.Part1(input))

	part2, err := day2.Part2(input)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "kqzxdenujwcstybmgvyiofrrd", part2)
}

func TestPart1(t *testing.T) {
	input := []string{
		"abcdef",
		"bababc",
		"abbcde",
		"abcccd",
		"aabcdd",
		"abcdee",
		"ababab",
	}
	assert.Equal(t, 12, day2.Part1(input))
}

func TestPart2(t *testing.T) {
	input := []string{
		"abcde",
		"fghij",
		"klmno",
		"pqrst",
		"fguij",
		"axcye",
		"wvxyz",
	}
	output, err := day2.Part2(input)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "fgij", output)
}
