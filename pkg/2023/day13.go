package puzzles2023

import (
	"fmt"
	"os"
	"time"

	"github.com/leandrorondon/advent-of-code/pkg/2023/day13"
)

func Day13(file string) error {
	t := time.Now()
	b, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	patterns := day13.Parse(string(b))

	t2 := time.Now()
	p1 := patterns.Summarise()

	took := time.Since(t)
	processing := time.Since(t2)
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", 2)
	fmt.Printf("(took %v, %v processing)\n", took, processing)

	return nil
}
