package advent2018

import (
	"bufio"
	"io"
	"strconv"
)

func ReadInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	var ret []int

	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		ret = append(ret, n)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return ret, nil
}
