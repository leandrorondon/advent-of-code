package puzzles2022

import (
	"bufio"
	"fmt"
	"github.com/leandrorondon/advent-of-code/pkg/2022/day17/tetris"
	"os"
	"time"
)

const ChamberWidth = 7

func Day17(file string) error {
	t := time.Now()
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	jetPattern := scanner.Bytes()

	t1 := time.Now()
	jetGen := tetris.NetJetGenerator(jetPattern)
	rockGen := tetris.NewRockGenerator()
	chamber := tetris.NewChamber(ChamberWidth, jetGen, rockGen, false)
	tetris.Simulate(chamber, 2022)
	part1 := chamber.Height
	took1 := time.Now().Sub(t1)
	fmt.Printf("- Part 1: %d (took %v)\n", part1, took1)

	t2 := time.Now()
	jetGen = tetris.NetJetGenerator(jetPattern)
	rockGen = tetris.NewRockGenerator()
	chamber = tetris.NewChamber(ChamberWidth, jetGen, rockGen, false)
	found1, found2 := tetris.FindSimilarPatterns(chamber)
	cycle := found2.Rocks - found1.Rocks
	incrementPerCycle := found2.Height - found1.Height
	n := 1_000_000_000_000
	cycles := (n - found2.Rocks) / cycle
	remaining := (n - found2.Rocks) % cycle
	tetris.Simulate(chamber, remaining)

	part2 := cycles*incrementPerCycle + chamber.Height
	took2 := time.Now().Sub(t2)
	fmt.Printf("- Part 2: %d (took %v)\n", part2, took2)

	took := time.Now().Sub(t)
	fmt.Printf("(took %v)\n", took)
	return nil
}
