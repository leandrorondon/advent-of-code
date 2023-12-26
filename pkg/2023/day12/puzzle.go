package day12

import (
	"strings"

	"github.com/leandrorondon/advent-of-code/pkg/math"
)

type Records []Record

func (r *Records) SumCombinations() int {
	sum := 0
	for _, record := range *r {
		sum += record.PossibleCombinations()
	}
	return sum
}

func trimMultiple(s string) string {
	news := strings.ReplaceAll(s, "..", ".")
	for news != s {
		s = news
		news = strings.ReplaceAll(s, "..", ".")
	}
	return s
}

type Record struct {
	Input     string
	Sequences []int
}

func (r *Record) PossibleCombinations() int {
	sum := 0
combinations:
	for _, possible := range r.BuildCombinations() {
		s := trimMultiple(possible)
		s = strings.TrimSuffix(strings.TrimPrefix(s, "."), ".")
		ss := strings.Split(s, ".")
		if len(ss) != len(r.Sequences) {
			continue
		}
		for i := range ss {
			if len(ss[i]) != r.Sequences[i] {
				continue combinations
			}
		}
		sum++
	}
	return sum
}

func (r *Record) BuildCombinations() []string {
	knownDamaged := strings.Count(r.Input, "#")
	needDamage := math.Sum(r.Sequences) - knownDamaged
	return generateCombinations(r.Input, needDamage)
}

func generateCombinations(s string, numberToReplace int) []string {
	positions := findUnknowns(s)
	var combinations []string
	replace(s, positions, 0, numberToReplace, &combinations)
	return combinations
}

func replace(s string, positions []int, currentIndex, numberToReplace int, combinations *[]string) {
	if numberToReplace == 0 {
		s = strings.ReplaceAll(s, "?", ".")
		*combinations = append(*combinations, s)
		return
	}

	for i := currentIndex; i < len(positions); i++ {
		strCopy := make([]rune, len(s))
		copy(strCopy, []rune(s))
		strCopy[positions[i]] = '#'
		replace(string(strCopy), positions, i+1, numberToReplace-1, combinations)
	}
}

func findUnknowns(str string) []int {
	var positions []int
	for i, c := range str {
		if c == '?' {
			positions = append(positions, i)
		}
	}
	return positions
}
