package puzzles2023

import (
	"fmt"
	"github.com/leandrorondon/advent-of-code/pkg/2023/day06"
	"os"
	"time"
)

func Day06(file string) error {
	t := time.Now()
	b, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	races := day06.ParseLines(string(b))
	margin := races.MarginOfError()

	took := time.Now().Sub(t)
	fmt.Println("Part 1:", margin)
	fmt.Println("Part 2:", 1)
	fmt.Printf("(took %v)\n", took)

	return nil
}
