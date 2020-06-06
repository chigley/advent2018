package main

import (
	"errors"
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

	part1 := Part1(input)

	part2, err := Part2(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(part1)
	fmt.Println(part2)
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

func Part2(input []string) (string, error) {
	for i, id1 := range input {
		for j, id2 := range input {
			if i >= j {
				continue
			}

			differingIndex, err := differingIndex(id1, id2)
			if err != nil {
				return "", fmt.Errorf("differingIndex: %w", err)
			}
			if differingIndex != -1 {
				return id1[:differingIndex] + id1[differingIndex+1:], nil
			}
		}
	}
	return "", errors.New("no solution found")
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

func differingIndex(s1, s2 string) (int, error) {
	if len(s1) != len(s2) {
		return 0, fmt.Errorf("box IDs %s and %s have different lengths", s1, s2)
	}

	differingIndex := -1

	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			continue
		}

		if differingIndex != -1 {
			return -1, nil
		}

		differingIndex = i
	}

	return differingIndex, nil
}
