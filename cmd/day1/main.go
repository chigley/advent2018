package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chigley/advent2018"
)

func main() {
	input, err := advent2018.ReadInts(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(Part1(input))
	fmt.Println(Part2(input))
}

func Part1(input []int) int {
	return advent2018.SumInts(input)
}

func Part2(input []int) int {
	seen := map[int]struct{}{
		0: {},
	}

	var freq int
	for {
		for _, n := range input {
			freq += n
			if _, ok := seen[freq]; ok {
				return freq
			}
			seen[freq] = struct{}{}
		}
	}
}
