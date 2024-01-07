package day17

import (
	"strings"
)

func Parse(s string) [][]int {
	lines := strings.Split(s, "\n")
	arr := make([][]int, len(lines))
	for y := range lines {
		arr[y] = make([]int, len(lines[y]))
		for x := range lines[y] {
			arr[y][x] = int(lines[y][x] - '0')
		}
	}

	return arr
}
