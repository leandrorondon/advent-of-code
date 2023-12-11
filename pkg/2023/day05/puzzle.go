package day05

import (
	"github.com/leandrorondon/advent-of-code/pkg/math"
)

type Range struct {
	Dest   int64
	Source int64
	Range  int64
}

type Map []Range

func (m Map) Convert(input int64) int64 {
	for i := range m {
		if math.BetweenInclusive(input, m[i].Source, m[i].Source+m[i].Range-1) {
			return m[i].Dest + (input - m[i].Source)
		}
	}

	// not in range
	return input
}

type Maps []Map

func (ms Maps) GetLocation(seed int64) int64 {
	tmp := seed
	for i := range ms {
		tmp = ms[i].Convert(tmp)
	}
	return tmp
}
