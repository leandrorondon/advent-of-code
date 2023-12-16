package day06

import (
	"log"
	"strconv"
	"strings"
)

func ParseLines(s string) (races Races) {
	lines := strings.Split(s, "\n")

	ss := strings.Split(lines[0], "Time:")
	times := parseIntSlice(strings.Split(strings.TrimSpace(ss[1]), " "))

	ss = strings.Split(lines[1], "Distance:")
	distances := parseIntSlice(strings.Split(strings.TrimSpace(ss[1]), " "))

	if len(distances) != len(times) {
		log.Panicln("different sizes", len(distances), len(times))
	}

	for i := range times {
		races = append(races, Race{times[i], distances[i]})
	}

	return races
}

func parseIntSlice(ss []string) (ii []int) {
	for i := len(ss) - 1; i >= 0; i-- {
		if in, err := strconv.Atoi(ss[i]); err == nil {
			ii = append(ii, in)
		}
	}
	return ii
}
