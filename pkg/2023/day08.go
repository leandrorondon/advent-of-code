package puzzles2023

import (
	"fmt"
	"os"
	"time"

	"github.com/leandrorondon/advent-of-code/pkg/2023/day08"
)

func Day08(file string) error {
	t := time.Now()
	b, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	sequence, initial := day08.ParseMap(string(b))
	steps, _ := initial.Steps(&sequence, day08.CompFn(day08.Final))

	sequence, initials := day08.ParseMap2(string(b))
	steps2 := initials.Steps(&sequence)

	took := time.Now().Sub(t)
	fmt.Println("Part 1:", steps)
	fmt.Println("Part 2:", steps2)
	fmt.Printf("(took %v)\n", took)

	return nil
}
