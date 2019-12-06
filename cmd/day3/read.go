package main

import (
	"fmt"
	"io"
	"regexp"

	"github.com/chigley/advent2018"
)

var regexpClaim = regexp.MustCompile(`^#(\d+) @ (\d+),(\d+): (\d+)x(\d+)$`)

func ReadClaims(r io.Reader) ([]Claim, error) {
	input, err := advent2018.ReadStrings(r)
	if err != nil {
		return nil, err
	}

	var ret []Claim

	for _, in := range input {
		strMatches := regexpClaim.FindStringSubmatch(in)
		if len(strMatches) != 6 {
			return nil, fmt.Errorf("failed to parse claim '%s'", in)
		}

		matches, err := advent2018.AtoiSlice(strMatches[1:])
		if err != nil {
			return nil, err
		}

		ret = append(ret, Claim{
			ID:      ClaimID(matches[0]),
			OffsetX: matches[1],
			OffsetY: matches[2],
			Width:   matches[3],
			Height:  matches[4],
		})
	}

	return ret, nil
}
