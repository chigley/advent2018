package advent2018

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func ReadInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	var ret []int

	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("advent2018: atoi: %w", err)
		}
		ret = append(ret, n)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("advent2018: scanner: %w", err)
	}

	return ret, nil
}

func ReadStrings(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	var ret []string

	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("advent2018: scanner: %w", err)
	}

	return ret, nil
}
