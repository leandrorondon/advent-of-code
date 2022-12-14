package puzzles2022

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/leandrorondon/advent-of-code/pkg/2022/day14/cave"
	"github.com/leandrorondon/advent-of-code/pkg/2022/day14/grid"
	"github.com/leandrorondon/advent-of-code/pkg/2022/day14/parser"
	"github.com/leandrorondon/advent-of-code/pkg/2022/day14/physics"
)

func Day14(file string) error {
	t := time.Now()
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	p := parser.NewParser(scanner)
	rockPaths := p.ScanRocks()

	t1 := time.Now()
	ph := physics.NewPhysics()
	dangerousCave := cave.New(grid.Coordinate{X: 500}, rockPaths, ph)
	part1 := dangerousCave.FillSand()
	took1 := time.Now().Sub(t1)
	fmt.Printf("- Part 1: %d (took %v)\n", part1, took1)

	t2 := time.Now()
	floorLevel := 9
	bigph := physics.NewBigPhysics(floorLevel)
	bigCave := cave.New(grid.Coordinate{X: 500}, rockPaths, bigph)
	bigCave.SetFloor(bigCave.Grid.Max.Y + 2)
	part2 := bigCave.FillSand()
	took2 := time.Now().Sub(t2)

	took := time.Now().Sub(t)

	fmt.Printf("- Part 2: %d (took %v)\n", part2, took2)
	fmt.Printf("(took %v)\n", took)

	return nil
}
