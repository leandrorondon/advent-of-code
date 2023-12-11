package puzzles2023

import (
	"fmt"
	"github.com/leandrorondon/advent-of-code/pkg/2023/day05"
	"math"
	"os"
	"time"
)

func Day05(file string) error {
	t := time.Now()
	b, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	seeds, maps := day05.ParseLines(string(b))

	lowest := int64(math.MaxInt64)
	for _, seed := range seeds {
		location := maps.GetLocation(seed)
		if location < lowest {
			lowest = location
		}
	}

	took := time.Now().Sub(t)
	fmt.Println("Part 1:", lowest)
	fmt.Println("Part 2:", 0)
	fmt.Printf("(took %v)\n", took)

	return nil
}
