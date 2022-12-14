package puzzles2022

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/leandrorondon/advent-of-code/pkg/2022/day16/parser"
	"github.com/leandrorondon/advent-of-code/pkg/2022/day16/volcano"
)

func Day16(file string) error {
	t := time.Now()
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	p := parser.NewParser(scanner)
	valves := p.ScanValves()

	t1 := time.Now()
	krakatoa := volcano.New(valves)
	part1, _ := krakatoa.HighestPossiblePressureReleased(valves["AA"], 30)
	took1 := time.Now().Sub(t1)
	fmt.Printf("- Part 1: %d (took %v)\n", part1, took1)

	t2 := time.Now()
	_, all := krakatoa.HighestPossiblePressureReleased(valves["AA"], 26)
	part2 := krakatoa.FindBestCombination(all)
	took2 := time.Now().Sub(t2)
	fmt.Printf("- Part 2: %d (took %v)\n", part2, took2)

	took := time.Now().Sub(t)
	fmt.Printf("(took %v)\n", took)

	return nil
}
