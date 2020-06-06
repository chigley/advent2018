package main_test

import (
	"log"
	"os"
	"testing"

	day4 "github.com/chigley/advent2018/cmd/day4"
	"github.com/stretchr/testify/assert"
)

func TestDay1(t *testing.T) {
	inputFile, err := os.Open("testdata/input")
	if err != nil {
		t.Fatal(err)
	}
	defer inputFile.Close()

	guards, err := day4.ReadEvents(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 26281, guards.Part1())
	assert.Equal(t, 73001, guards.Part2())
}
