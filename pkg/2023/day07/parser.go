package day07

import (
	"strconv"
	"strings"
)

func ParsePlays(s string) Plays {
	lines := strings.Split(s, "\n")

	var plays Plays
	for _, line := range lines {
		ss := strings.Split(line, " ")
		bid, _ := strconv.Atoi(ss[1])
		plays = append(plays, Play{
			Hand: NewHand(ss[0]),
			Bid:  bid,
		})
	}

	return plays
}
