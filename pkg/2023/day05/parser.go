package day05

import (
	"strconv"
	"strings"
)

func ParseLines(s string) ([]int64, Maps) {
	lines := strings.Split(s, "\n")

	ss := strings.Split(lines[0], "seeds: ")
	seeds := strings.Split(ss[1], " ")
	var iseeds []int64
	for _, seed := range seeds {
		i, _ := strconv.ParseInt(seed, 10, 64)
		iseeds = append(iseeds, i)
	}

	mapsLines := parseMaps(lines[2:])

	return iseeds, mapsLines
}

func parseMaps(lines []string) Maps {
	var maps Maps

	var buffer Map
	for _, line := range lines {
		if line == "" && len(buffer) > 0 {
			maps = append(maps, buffer)
			buffer = make([]Range, 0)
		}

		ss := strings.Split(line, " ")
		if len(ss) == 3 {
			n := parseNumbers(ss)
			buffer = append(buffer, Range{
				Dest:   n[0],
				Source: n[1],
				Range:  n[2],
			})
		}
	}

	return maps
}

func parseNumbers(ss []string) []int64 {
	var is []int64
	for _, s := range ss {
		i, _ := strconv.ParseInt(s, 10, 64)
		is = append(is, i)
	}
	return is
}
