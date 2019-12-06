package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

const fabricSize = 1000

type Claim struct {
	ID      int
	OffsetX int
	OffsetY int
	Width   int
	Height  int
}

func main() {
	claims, err := ReadClaims(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	part1, err := Part1(claims)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(part1)
}

func Part1(claims []Claim) (int, error) {
	var fabric [fabricSize][fabricSize]int
	var contested int

	for _, c := range claims {
		for x := c.OffsetX; x < c.OffsetX+c.Width; x++ {
			for y := c.OffsetY; y < c.OffsetY+c.Height; y++ {
				if x >= fabricSize || y >= fabricSize {
					return 0, errors.New("claim exceeds bounds of fabric")
				}

				fabric[x][y]++
				if fabric[x][y] == 2 {
					contested++
				}
			}
		}
	}

	return contested, nil
}
