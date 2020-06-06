package advent2018

import (
	"fmt"
	"strconv"
)

func AtoiSlice(input []string) ([]int, error) {
	ret := make([]int, len(input))
	for i, str := range input {
		n, err := strconv.Atoi(str)
		if err != nil {
			return nil, fmt.Errorf("advent2018: atoi: %w", err)
		}
		ret[i] = n
	}
	return ret, nil
}
