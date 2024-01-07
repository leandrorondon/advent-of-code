package puzzles2023

import (
	"fmt"
	"os"
	"time"

	"github.com/leandrorondon/advent-of-code/pkg/2023/day17"
)

func Day17(file string) error {
	t := time.Now()
	input, _ := os.ReadFile(file)
	arr := day17.Parse(string(input))

	p1 := day17.MinLoss(arr, 1, 3)
	p2 := day17.MinLoss(arr, 4, 10)

	took := time.Since(t)
	fmt.Printf("(took %v)\n", took)
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)

	return nil
}
