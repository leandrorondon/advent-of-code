package puzzles2023

import (
	"fmt"
	"github.com/leandrorondon/advent-of-code/pkg/2023/day04"
	"os"
	"time"
)

func Day04(file string) error {
	t := time.Now()
	b, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	cards := day04.ParseLines(string(b))

	cardPointSum := 0

	for _, card := range cards {
		ps := card.Points()
		cardPointSum += ps
	}

	took := time.Now().Sub(t)
	fmt.Println("Part 1:", cardPointSum)
	fmt.Printf("(took %v)\n", took)

	return nil
}
