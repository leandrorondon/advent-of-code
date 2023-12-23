package puzzles2023

import (
	"fmt"
	"os"
	"time"

	"github.com/leandrorondon/advent-of-code/pkg/2023/day10"
)

func Day10(file string) error {
	t := time.Now()
	b, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	m := day10.ParseMap(string(b))

	m.BuildNeighbourhood()

	p1 := m.Farthest(m.Start)

	took := time.Now().Sub(t)
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", 2)
	fmt.Printf("(took %v)\n", took)

	return nil
}
