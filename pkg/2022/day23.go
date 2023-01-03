package puzzles2022

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/leandrorondon/advent-of-code/pkg/2022/day23"
)

func Day23(file string) error {
	t := time.Now()

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	p := day23.NewParser(scanner)
	elfs := p.Parse()
	elfs2 := make([]day23.Elf, len(elfs))
	copy(elfs2, elfs)

	t1 := time.Now()
	grove := day23.NewGrove(elfs)
	grove.RunRounds(10)
	part1 := grove.EmptySpaces()
	took1 := time.Now().Sub(t1)
	fmt.Printf("- Part 1: %d (took %v)\n", part1, took1)

	t2 := time.Now()
	grove = day23.NewGrove(elfs2)
	part2 := grove.RunAll()
	took2 := time.Now().Sub(t2)
	fmt.Printf("- Part 2: %d (took %v)\n", part2, took2)

	took := time.Now().Sub(t)
	fmt.Printf("(took %v)\n", took)

	return nil
}
