package day09

import (
	"strconv"
	"strings"
)

func ParseHistories(s string) Histories {
	lines := strings.Split(s, "\n")

	histories := make(Histories, len(lines))

	for i, l := range lines {
		ss := strings.Split(l, " ")
		h := make(History, len(ss))
		for j, s := range ss {
			n, _ := strconv.Atoi(s)
			h[j] = n
		}
		histories[i] = h
	}

	return histories
}
