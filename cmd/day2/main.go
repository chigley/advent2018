package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chigley/advent2018"
)

func main() {
	input, err := advent2018.ReadStrings(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(Part1(input))
}

func Part1(input []string) int {
	var two, three int

	for _, str := range input {
		exactlyTwo, exactlyThree := countChars(str)
		if exactlyTwo {
			two++
		}
		if exactlyThree {
			three++
		}
	}

	return two * three
}

func countChars(s string) (exactlyTwo, exactlyThree bool) {
	counts := make(map[rune]int)
	for _, r := range s {
		counts[r]++
	}

	for _, count := range counts {
		if count == 2 {
			exactlyTwo = true
		} else if count == 3 {
			exactlyThree = true
		}
		if exactlyTwo && exactlyThree {
			return
		}
	}

	return
}
