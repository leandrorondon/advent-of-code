package puzzles2023

import (
	"fmt"
	"github.com/leandrorondon/advent-of-code/pkg/2023/day06"
	"math"
	"os"
	"time"
)

func Day06(file string) error {
	t := time.Now()
	b, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	races := day06.ParseRaces(string(b))
	margin := races.MarginOfError()

	race := day06.ParseSingleRace(string(b))
	fmt.Println(race, math.MaxInt)
	waysToWin := race.WaysToWin()

	took := time.Now().Sub(t)
	fmt.Println("Part 1:", margin)
	fmt.Println("Part 2:", waysToWin)
	fmt.Printf("(took %v)\n", took)

	return nil
}
