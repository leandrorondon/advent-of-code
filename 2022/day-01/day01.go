package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	b, _ := os.ReadFile("input-1.txt")
	lines := strings.Split(string(b), "\n")

	var a []int
	acc := 0
	for _, line := range lines {
		if line == "" {
			a = append(a, acc)
			acc = 0
			continue
		}

		v, _ := strconv.Atoi(line)
		acc += v
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})

	fmt.Println("Top 1:", a[0])
	fmt.Println("Top 3:", a[0]+a[1]+a[2])
}
