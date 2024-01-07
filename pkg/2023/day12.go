package puzzles2023

import (
	"fmt"
	"os"
	"time"
	
	"github.com/leandrorondon/advent-of-code/pkg/2023/day12"
)

func Day12(file string) error {
	t := time.Now()
	b, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	m := day12.Parse(string(b))
	p1 := m.SumCombinations()

	exp := m.Unfold(5)
	p2 := exp.SumCombinations()

	took := time.Now().Sub(t)
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
	fmt.Printf("(took %v)\n", took)

	return nil
}
