package puzzles2023

import (
	"fmt"
	"github.com/leandrorondon/advent-of-code/pkg/2023/day12"
	"os"
	"time"
)

func Day12(file string) error {
	t := time.Now()
	b, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	m := day12.Parse(string(b))
	p1 := m.SumCombinations()

	took := time.Now().Sub(t)
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", 1)
	fmt.Printf("(took %v)\n", took)

	return nil
}
