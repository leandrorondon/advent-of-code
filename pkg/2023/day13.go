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

	p1 := patterns.Summarise()
	p2 := patterns.SummariseSmudge()

	took := time.Since(t)
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
	fmt.Printf("(took %v)\n", took)

	return nil
}
