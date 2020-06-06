package main

import (
	"fmt"
	"io"
	"regexp"
	"sort"
	"strconv"

	"github.com/chigley/advent2018"
)

const (
	fallsAsleep = "falls asleep"
	wakesUp     = "wakes up"
)

var (
	regexpEvent      = regexp.MustCompile(`(\d+)] (.+)$`)
	regexpBeginShift = regexp.MustCompile(`^Guard #(\d+) begins shift$`)
)

func ReadEvents(r io.Reader) (Guards, error) {
	input, err := advent2018.ReadStrings(r)
	if err != nil {
		return nil, fmt.Errorf("ReadStrings: %w", err)
	}

	// Lexicographical ordering on ISO 8601 dates coincides with the
	// chronological order
	sort.Strings(input)

	ret := make(Guards)
	guardID, sleepingFrom := GuardID(-1), -1

	for _, in := range input {
		matches := regexpEvent.FindStringSubmatch(in)
		if len(matches) != 3 {
			return nil, fmt.Errorf("failed to parse event '%s'", in)
		}
		minutes, err := strconv.Atoi(matches[1])
		if err != nil {
			return nil, fmt.Errorf("atoi: %w", err)
		}
		event := matches[2]

		if event == fallsAsleep {
			sleepingFrom = minutes
		} else if event == wakesUp {
			guardArr := ret[guardID]
			for i := sleepingFrom; i < minutes; i++ {
				guardArr[i]++
			}
			ret[guardID] = guardArr
		} else {
			matches = regexpBeginShift.FindStringSubmatch(event)
			if len(matches) != 2 {
				return nil, fmt.Errorf("failed to parse event '%s'", in)
			}
			guardInt, err := strconv.Atoi(matches[1])
			if err != nil {
				return nil, fmt.Errorf("atoi: %w", err)
			}
			guardID = GuardID(guardInt)
		}
	}

	return ret, nil
}
