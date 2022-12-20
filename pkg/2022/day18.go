package puzzles2022

import (
	"bufio"
	"fmt"
	"github.com/leandrorondon/advent-of-code/pkg/2022/day18"
	"os"
	"time"
)

func Day18(file string) error {
	t := time.Now()
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	p := day18.NewParser(scanner)
	droplets := p.ScanDroplets()
	droplets2 := make([]day18.XYZ, len(droplets))
	copy(droplets2, droplets)

	t1 := time.Now()
	part1 := day18.SurfaceArea(droplets)
	took1 := time.Now().Sub(t1)
	fmt.Printf("- Part 1: %d (took %v)\n", part1, took1)

	t2 := time.Now()
	part2 := day18.ExternalSurface(droplets2)
	took2 := time.Now().Sub(t2)
	fmt.Printf("- Part 2: %d (took %v)\n", part2, took2)

	took := time.Now().Sub(t)
	fmt.Printf("(took %v)\n", took)
	return nil
}
