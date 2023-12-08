package puzzles2023

import (
	"fmt"
	"github.com/leandrorondon/advent-of-code/pkg/2023/day03"
	"os"
	"strconv"
	"time"
)

func Day03(file string) error {
	t := time.Now()
	b, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	symbols := day03.ParseLines(string(b))

	partNumberSum := 0

	for _, symbolsLine := range symbols {
		for _, sym := range symbolsLine {
			if sym.IsNumber() {
				if sym.HasAdjacentSymbol(symbols) {
					value, _ := strconv.Atoi(sym.Value)
					partNumberSum += value
				}
			}
		}
	}

	took := time.Now().Sub(t)
	fmt.Println("Part 1:", partNumberSum)
	fmt.Printf("(took %v)\n", took)

	return nil
}
