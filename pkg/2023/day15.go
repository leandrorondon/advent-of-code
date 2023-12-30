package puzzles2023

import (
	"fmt"
	"os"
	"time"

	"github.com/leandrorondon/advent-of-code/pkg/2023/day15"
)

func Day15(file string) error {
	t := time.Now()
	b, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	p1 := day15.HashP1(b)
	p2 := day15.TotalP2(b)

	took := time.Since(t)
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
	fmt.Printf("(took %v)\n", took)

	return nil
}
