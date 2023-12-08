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

	totalCards := 0
	for i, card := range cards {
		matches := card.Matches()
		for k := 1; k <= matches && (i+k) < len(cards); k++ {
			cards[i+k].Copies += card.Copies
		}

		totalCards += card.Copies
	}

	took := time.Now().Sub(t)
	fmt.Println("Part 1:", cardPointSum)
	fmt.Println("Part 2:", totalCards)
	fmt.Printf("(took %v)\n", took)

	return nil
}
