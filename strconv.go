package advent2018

import (
	"strconv"
)

func AtoiSlice(input []string) ([]int, error) {
	ret := make([]int, len(input))
	for i, str := range input {
		n, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		ret[i] = n
	}
	return ret, nil
}
