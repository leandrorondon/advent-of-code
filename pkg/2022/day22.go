package puzzles2022

import (
	"fmt"
	"strings"
	"time"

	"github.com/leandrorondon/advent-of-code/pkg/2022/day22"
)

func Day22(file string) error {
	t := time.Now()

	p := day22.NewParser(file)
	m, instructions := p.Parse()

	t1 := time.Now()
	board := day22.NewBoard(m, instructions, day22.NewWrapper2D(m))
	part1 := board.Run()
	took1 := time.Now().Sub(t1)
	fmt.Printf("- Part 1: %d (took %v)\n", part1, took1)

	t2 := time.Now()
	if strings.Contains(file, "_test.txt") {
		board = day22.NewBoard(m, instructions, day22.NewWrapper3DT(m))
	} else {
		board = day22.NewBoard(m, instructions, day22.NewWrapper3D(m))
	}

	part2 := board.Run()
	took2 := time.Now().Sub(t2)
	fmt.Printf("- Part 2: %d (took %v)\n", part2, took2)

	took := time.Now().Sub(t)
	fmt.Printf("(took %v)\n", took)

	return nil
}
