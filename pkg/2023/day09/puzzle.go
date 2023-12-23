package day09

import "fmt"

type History []int

func (h History) PrintHistoryLevels() {
	levels := h.Levels()

	for _, l := range levels {
		fmt.Println(l)
	}
}

func (h History) Levels() []History {
	var levels []History
	levels = append(levels, h)
	level := 0
	size := len(h) - 1
	allZeroes := false
	for !allZeroes {
		allZeroes = true
		hh := make(History, size)
		for i := 0; i < size; i++ {
			hh[i] = levels[level][i+1] - levels[level][i]
			if hh[i] != 0 {
				allZeroes = false
			}
		}

		levels = append(levels, hh)
		level++
		size--
	}

	return levels
}

func (h History) Extrapolate() int {
	levels := h.Levels()

	last := len(levels) - 1
	levels[last] = append(levels[last], 0)
	for i := len(levels) - 2; i >= 0; i-- {
		last = len(levels[i]) - 1
		levels[i] = append(levels[i], levels[i][last]+levels[i+1][last])
	}

	return levels[0][last+1]
}

func (h History) ExtrapolateBack() int {
	levels := h.Levels()

	last := len(levels) - 1
	levels[last] = append([]int{0}, levels[last]...)

	for i := len(levels) - 2; i >= 0; i-- {
		levels[i] = append([]int{
			levels[i][0] - levels[i+1][0],
		}, levels[i]...)
	}

	return levels[0][0]
}

type Histories []History

func (hs Histories) SumExtrapolations() int {
	sum := 0
	for i := range hs {
		sum += hs[i].Extrapolate()
	}

	return sum
}

func (hs Histories) SumExtrapolationsBack() int {
	sum := 0
	for i := range hs {
		sum += hs[i].ExtrapolateBack()
	}

	return sum
}
