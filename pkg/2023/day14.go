package puzzles2023

import (
	"fmt"
	"os"
	"time"

	"github.com/leandrorondon/advent-of-code/pkg/2023/day14"
)

func Day14(file string) error {
	t := time.Now()
	b, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	m := day14.Parse(string(b))
	m.TiltNorth()
	p1 := m.Load()

	//m.Reset()
	m.Cycle(1000000000)
	p2 := m.Load()

	took := time.Since(t)
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
	fmt.Printf("(took %v)\n", took)

	return nil
}
