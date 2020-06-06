package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chigley/advent2018"
)

type (
	GuardID int
	Guards  map[GuardID][60]int
)

func main() {
	guards, err := ReadEvents(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(guards.Part1())
	fmt.Println(guards.Part2())
}

func (gs Guards) Part1() int {
	guardID, maxSleep := GuardID(-1), -1
	for g, freqs := range gs {
		if sleep := advent2018.SumInts(freqs[:]); sleep > maxSleep {
			guardID = g
			maxSleep = sleep
		}
	}

	minute, maxFreq := -1, -1
	for m, freq := range gs[guardID] {
		if freq > maxFreq {
			minute = m
			maxFreq = freq
		}
	}

	return int(guardID) * minute
}

func (gs Guards) Part2() int {
	minute, guardID, maxFreq := -1, GuardID(-1), -1
	for g, freqs := range gs {
		for m, freq := range freqs {
			if freq > maxFreq {
				minute = m
				guardID = g
				maxFreq = freq
			}
		}
	}
	return int(guardID) * minute
}
