package day04

import (
	"strconv"
	"strings"
)

func ParseLines(s string) []*card {
	lines := strings.Split(s, "\n")

	var cards []*card

	for _, line := range lines {
		s := strings.Split(strings.Split(line, ":")[1], "|")

		winning := parseNumbers(strings.ReplaceAll(strings.TrimSpace(s[0]), "  ", " "))
		have := parseNumbers(strings.ReplaceAll(strings.TrimSpace(s[1]), "  ", " "))

		cards = append(cards, &card{
			Winning: winning,
			Have:    have,
			Copies:  1,
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
