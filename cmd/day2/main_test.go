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
