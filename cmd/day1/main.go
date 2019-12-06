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
}

func Part1(input []int) int {
	return advent2018.SumInts(input)
}
