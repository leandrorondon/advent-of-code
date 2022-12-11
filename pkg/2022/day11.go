package puzzles2022

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/leandrorondon/advent-of-code/pkg/2022/day11/monkey"
)

func Day11(file string) error {
	t := time.Now()
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	p := monkey.NewParser(scanner)
	var monkeys []*monkey.Monkey
	for m := p.ParseMonkey(); m != nil; m = p.ParseMonkey() {
		monkeys = append(monkeys, m)
	}

	rounds := 20
	for round := 1; round <= rounds; round++ {
		for i := 0; i < len(monkeys); i++ {
			monkeys[i].Turn(monkeys)
		}
	}

	var inspections []int
	for i := 0; i < len(monkeys); i++ {
		inspections = append(inspections, monkeys[i].Inspections())
	}
	sort.Ints(inspections)

	level := inspections[len(inspections)-1] * inspections[len(inspections)-2]

	took := time.Now().Sub(t)
	fmt.Println("- Level of monkey business Part 1:", level)
	fmt.Println("- Part 2:", 1)
	fmt.Printf("(took %v)\n", took)

	return nil
}
