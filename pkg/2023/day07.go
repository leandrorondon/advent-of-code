package puzzles2023

import (
	"fmt"
	"os"
	"time"

	"github.com/leandrorondon/advent-of-code/pkg/2023/day07"
)

func Day07(file string) error {
	t := time.Now()
	b, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	plays := day07.ParsePlays(string(b), day07.CardValuesP1, day07.GetCardsAndCountsP1)
	total := plays.TotalWins()

	plays2 := day07.ParsePlays(string(b), day07.CardValuesP2, day07.GetCardsAndCountsP2)
	total2 := plays2.TotalWins()

	took := time.Now().Sub(t)
	fmt.Println("Part 1:", total)
	fmt.Println("Part 2:", total2)
	fmt.Printf("(took %v)\n", took)

	return nil
}
