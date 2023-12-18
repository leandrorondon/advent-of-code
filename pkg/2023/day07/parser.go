package day07

import (
	"strconv"
	"strings"
)

func ParsePlays(s string, values map[rune]int, counter counterFunc) Plays {
	lines := strings.Split(s, "\n")

	rev := make(map[int]rune)
	for k, v := range values {
		rev[v] = k
	}

	var plays Plays
	plays.values = values
	for _, line := range lines {
		ss := strings.Split(line, " ")
		bid, _ := strconv.Atoi(ss[1])
		plays.plays = append(plays.plays, Play{
			Hand: NewHand(ss[0], values, rev, counter),
			Bid:  bid,
		})
	}

	return plays
}
