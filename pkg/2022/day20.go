package puzzles2022

import (
	"fmt"
	"time"

	"github.com/leandrorondon/advent-of-code/pkg/2022/day20"
)

func Day20(file string) error {
	t := time.Now()

	p := day20.NewParser(file)
	sl, zero := p.Parse()

	t1 := time.Now()
	g1, g2, g3 := day20.GrooveCoordinates(sl, zero)
	part1 := g1 + g2 + g3
	took1 := time.Now().Sub(t1)
	fmt.Printf("- Part 1: %d (took %v)\n", part1, took1)

	t2 := time.Now()
	sl, zero = p.Parse()
	g1, g2, g3 = day20.GrooveCoordinatesWithKey(sl, zero, 811589153)
	part2 := g1 + g2 + g3
	took2 := time.Now().Sub(t2)
	fmt.Printf("- Part 2: %d (took %v)\n", part2, took2)

	took := time.Now().Sub(t)
	fmt.Printf("(took %v)\n", took)
	return nil
}
