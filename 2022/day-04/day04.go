package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	t := time.Now()
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	contains := 0
	overlaps := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		ranges := strings.Split(line, ",")
		r1 := NewRange(ranges[0])
		r2 := NewRange(ranges[1])

		if r1.Contains(r2) || r2.Contains(r1) {
			contains++
			overlaps++
			continue
		}

		if r1.Overlaps(r2) {
			overlaps++
		}

	}

	took := time.Now().Sub(t)
	fmt.Println("Fully contained:", contains)
	fmt.Println("Overlaps       :", overlaps)
	fmt.Printf("(took %v)\n", took)
}

type Range struct {
	Lower int
	Upper int
}

func (r Range) Contains(r2 Range) bool {
	return r.Lower <= r2.Lower && r.Upper >= r2.Upper
}

func (r Range) Overlaps(r2 Range) bool {
	return (r.Upper >= r2.Lower && r.Upper <= r2.Upper) ||
		(r2.Upper >= r.Lower && r2.Upper <= r.Upper)
}

func NewRange(r string) Range {
	s := strings.Split(r, "-")
	l, _ := strconv.Atoi(s[0])
	u, _ := strconv.Atoi(s[1])

	return Range{
		Lower: l,
		Upper: u,
	}
}
