package puzzles2022

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/leandrorondon/advent-of-code/pkg/2022/day12/heightmap"
)

func Day12(file string) error {
	t := time.Now()
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	hm := heightmap.NewParser(scanner).Parse()

	t1 := time.Now()
	part1 := hm.ShortestPath(hm.Start, hm.End)
	took1 := time.Now().Sub(t1)

	t2 := time.Now()
	part2 := hm.ShortestPathFromHeight('a', hm.End)
	took2 := time.Now().Sub(t2)

	took := time.Now().Sub(t)
	fmt.Printf("- Part 1: %d (took %v)\n", part1, took1)
	fmt.Printf("- Part 2: %d (took %v)\n", part2, took2)
	fmt.Printf("(took %v)\n", took)

	return nil
}
