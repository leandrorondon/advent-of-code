package day14

import (
	"strings"
)

func Parse(s string) *Map {
	lines := strings.Split(s, "\n")
	m := make([][]byte, len(lines))
	for i := range lines {
		m[i] = []byte(lines[i])
	}

	return &Map{m: m}
}
