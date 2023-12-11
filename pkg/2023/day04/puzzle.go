package day04

import (
	"math"
	"slices"
)

type card struct {
	Winning []int
	Have    []int
	Copies  int
}

func (s card) Points() int {
	matches := s.Matches()

	if matches == 0 {
		return 0
	}

	return int(math.Pow(2, float64(matches-1)))
}

func (s card) Matches() int {
	matches := 0

	for _, w := range s.Winning {
		if slices.Contains(s.Have, w) {
			matches++
		}
	}

	return matches
}
