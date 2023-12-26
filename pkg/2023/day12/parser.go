package day12

import (
	"strconv"
	"strings"
)

func Parse(s string) Records {
	lines := strings.Split(s, "\n")

	var records Records
	for i := range lines {
		ss := strings.Split(lines[i], " ")
		nn := strings.Split(ss[1], ",")
		seq := make([]int, len(nn))
		for j := range nn {
			seq[j], _ = strconv.Atoi(nn[j])
		}

		records = append(records, Record{
			Input:     ss[0],
			Sequences: seq,
		})
	}

	return records
}
