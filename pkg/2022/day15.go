package puzzles2022

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/leandrorondon/advent-of-code/pkg/2022/day15/grid"
	"github.com/leandrorondon/advent-of-code/pkg/2022/day15/parser"
)

func Day15(file string) error {
	t := time.Now()
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	input1 := 2000000
	input2 := 4000000
	if strings.Contains(file, "_test.txt") {
		input1 = 20
		input2 = 20
	}

	scanner := bufio.NewScanner(f)
	p := parser.NewParser(scanner)
	pairs := p.ScanPairs()
	g := grid.New(pairs)

	t1 := time.Now()
	part1 := g.CountNoBeaconInRow(input1, grid.NewSimpleProjector())
	took1 := time.Now().Sub(t1)
	fmt.Printf("- Part 1: %d (took %v)\n", part1, took1)

	t2 := time.Now()
	part2 := g.FindTuningFrequency(input2, grid.NewLimitedProjector(input2))
	took2 := time.Now().Sub(t2)
	fmt.Printf("- Part 2: %d (took %v)\n", part2, took2)

	took := time.Now().Sub(t)
	fmt.Printf("(took %v)\n", took)

	return nil
}
