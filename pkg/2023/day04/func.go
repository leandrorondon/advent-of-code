package day04

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

type card struct {
	Winning []int
	Have    []int
}

func (s card) Points() int {
	matches := 0

	for _, w := range s.Winning {
		if slices.Contains(s.Have, w) {
			matches++
		}
	}

	if matches == 0 {
		return 0
	}

	return int(math.Pow(2, float64(matches-1)))
}

func ParseLines(s string) []card {
	lines := strings.Split(s, "\n")

	var cards []card

	for _, line := range lines {
		s := strings.Split(strings.Split(line, ":")[1], "|")

		winning := parseNumbers(strings.ReplaceAll(strings.TrimSpace(s[0]), "  ", " "))
		have := parseNumbers(strings.ReplaceAll(strings.TrimSpace(s[1]), "  ", " "))

		cards = append(cards, card{
			Winning: winning,
			Have:    have,
		})
	}
	return cards
}

func parseNumbers(input string) []int {
	ss := strings.Split(input, " ")
	var is []int
	for _, s := range ss {
		i, _ := strconv.Atoi(s)
		is = append(is, i)
	}
	return is
}
