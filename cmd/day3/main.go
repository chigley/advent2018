package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

const fabricSize = 1000

type ClaimID int

type Claim struct {
	ID      ClaimID
	OffsetX int
	OffsetY int
	Width   int
	Height  int
}

type squareInch struct {
	firstOccupier ClaimID
	count         int
}

func main() {
	claims, err := ReadClaims(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	part1, part2, err := Day3(claims)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(part1)
	fmt.Println(part2)
}

func Day3(claims []Claim) (part1, part2 int, err error) {
	var fabric [fabricSize][fabricSize]squareInch
	var contested int

	intact := make(map[ClaimID]struct{})

	for _, c := range claims {
		intact[c.ID] = struct{}{}

		for x := c.OffsetX; x < c.OffsetX+c.Width; x++ {
			for y := c.OffsetY; y < c.OffsetY+c.Height; y++ {
				if x >= fabricSize || y >= fabricSize {
					return 0, 0, errors.New("claim exceeds bounds of fabric")
				}

				fabric[x][y].count++
				if fabric[x][y].count == 1 {
					fabric[x][y].firstOccupier = c.ID
				} else {
					if fabric[x][y].count == 2 {
						contested++
						delete(intact, fabric[x][y].firstOccupier)
					}
					delete(intact, c.ID)
				}
			}
		}
	}

	if len(intact) != 1 {
		return 0, 0, fmt.Errorf("got %d part 2 solutions, expected 1", len(intact))
	}
	var intactID ClaimID
	for intactID = range intact {
	}
	return contested, int(intactID), nil
}
