package puzzles2023

import (
	"fmt"
	"os"
	"time"

	"github.com/leandrorondon/advent-of-code/pkg/2023/day11"
)

func Day11(file string) error {
	t := time.Now()
	b, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	m := day11.Parse(string(b))

	exp := m.Expand()
	p1 := exp.SumDistances()

	took := time.Now().Sub(t)
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", 0)
	fmt.Printf("(took %v)\n", took)

	return nil
}
