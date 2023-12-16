package day06

import (
	"log"
	"strconv"
	"strings"
)

func ParseRaces(s string) (races Races) {
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

func ParseSingleRace(s string) (race Race) {
	lines := strings.Split(s, "\n")

	ss := strings.Split(lines[0], "Time:")
	tt := strings.ReplaceAll(ss[1], " ", "")
	time, _ := strconv.Atoi(tt)

	ss = strings.Split(lines[1], "Distance:")
	dd := strings.ReplaceAll(ss[1], " ", "")
	distance, _ := strconv.Atoi(dd)

	return Race{time, distance}
}

func parseIntSlice(ss []string) (ii []int) {
	for i := len(ss) - 1; i >= 0; i-- {
		if in, err := strconv.Atoi(ss[i]); err == nil {
			ii = append(ii, in)
		}
	}
	return ii
}
