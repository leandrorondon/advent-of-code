package day12

import (
	"strings"

	"github.com/leandrorondon/advent-of-code/pkg/cache"
)

type Records []Record

func (r *Records) SumCombinations() int {
	sum := 0
	for _, record := range *r {
		sum += record.PossibleCombinations()
	}
	return sum
}

func repeatSlice(sl []int, n int) []int {
	newsl := make([]int, n*len(sl))

	for i := 0; i < n; i++ {
		for j := range sl {
			newsl[i*len(sl)+j] = sl[j]
		}
	}
	return newsl
}

func repeatInput(input, div string, n int) string {
	ss := make([]string, n)
	for i := range ss {
		ss[i] = input
	}
	return strings.Join(ss, div)
}

func (r *Records) Unfold(n int) *Records {
	var newR Records

	for _, record := range *r {
		newR = append(newR, Record{
			Input:     repeatInput(record.Input, "?", n),
			Sequences: repeatSlice(record.Sequences, n),
		})
	}

	return &newR
}

type Record struct {
	Input     string
	Sequences []int
}

func (r *Record) calcPossible(i, j int, ca *cache.TwoD) int {
	if i >= len(r.Input) {
		if j < len(r.Sequences) {
			return 0
		}
		return 1
	}

	if v, ok := ca.Get(i, j); ok {
		return v
	}

	res := 0
	if r.Input[i] == '.' {
		res = r.calcPossible(i+1, j, ca)
	} else {
		if r.Input[i] == '?' {
			res += r.calcPossible(i+1, j, ca)
		}
		if j < len(r.Sequences) {
			count := 0
			for k := i; k < len(r.Input); k++ {
				if count > r.Sequences[j] || r.Input[k] == '.' || count == r.Sequences[j] && r.Input[k] == '?' {
					break
				}
				count += 1
			}

			if count == r.Sequences[j] {
				if i+count < len(r.Input) && r.Input[i+count] != '#' {
					res += r.calcPossible(i+count+1, j+1, ca)
				} else {
					res += r.calcPossible(i+count, j+1, ca)
				}
			}
		}
	}

	ca.Set(i, j, res)
	return res
}

func (r *Record) PossibleCombinations() int {
	ca := cache.New2D()

	return r.calcPossible(0, 0, ca)
}
