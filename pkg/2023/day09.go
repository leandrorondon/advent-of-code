package puzzles2023

import (
	"fmt"
	"os"
	"time"

	"github.com/leandrorondon/advent-of-code/pkg/2023/day09"
)

func Day09(file string) error {
	t := time.Now()
	b, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	histories := day09.ParseHistories(string(b))
	sum := histories.SumExtrapolations()

	took := time.Now().Sub(t)
	fmt.Println("Part 1:", sum)
	fmt.Println("Part 2:", 1)
	fmt.Printf("(took %v)\n", took)

	return nil
}
