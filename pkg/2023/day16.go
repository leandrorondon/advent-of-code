package puzzles2023

import (
	"fmt"
	"github.com/leandrorondon/advent-of-code/pkg/2023/day16"
	"os"
	"time"
)

func Day16(file string) error {
	t := time.Now()
	b, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	m := day16.Parse(string(b))

	p1 := m.Ray(0, 0, day16.Direction(2))
	p2 := m.FindMax()

	took := time.Since(t)
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
	fmt.Printf("(took %v)\n", took)

	return nil
}
