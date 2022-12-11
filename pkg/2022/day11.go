package puzzles2022

import (
	"bufio"
	"fmt"
	"math/big"
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

	// Parse file and build Monkey list
	scanner := bufio.NewScanner(f)
	p := monkey.NewParser(scanner)
	var monkeys []*monkey.Monkey
	for m := p.ParseMonkey(); m != nil; m = p.ParseMonkey() {
		monkeys = append(monkeys, m)
		m.SetAdjustLevel(3)
	}

	var bigMonkeys []*monkey.BigMonkey
	div := lcmMonkeys(monkeys)
	for _, m := range monkeys {
		newM := monkey.NewBig(m)
		newM.SetAdjustLevel(div)
		bigMonkeys = append(bigMonkeys, newM)
	}

	// Part 1: 20 rounds using normal Monkey
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
	level1 := inspections[len(inspections)-1] * inspections[len(inspections)-2]

	// Part 2: 10000 rounds using BigMonkey
	rounds = 10000
	for round := 1; round <= rounds; round++ {
		for i := 0; i < len(bigMonkeys); i++ {
			bigMonkeys[i].Turn(bigMonkeys)
		}
	}

	inspections = []int{}
	for i := 0; i < len(monkeys); i++ {
		inspections = append(inspections, bigMonkeys[i].Inspections())
	}
	sort.Ints(inspections)
	b := big.NewInt(int64(inspections[len(inspections)-1]))
	b2 := big.NewInt(int64(inspections[len(inspections)-2]))
	b.Mul(b, b2)

	took := time.Now().Sub(t)
	fmt.Println("- Level of monkey business Part 1:", level1)
	fmt.Println("- Level of monkey business Part 2:", b.String())
	fmt.Printf("(took %v)\n", took)

	return nil
}

func lcmMonkeys(m []*monkey.Monkey) int {
	if len(m) < 2 {
		return 1
	}

	result := m[0].Divisor()

	for i := 1; i < len(m); i++ {
		result = lcm(result, m[i].Divisor())
	}

	return result
}

func lcm(x, y int) int {
	return x * y / gcd(x, y)
}

func gcd(x, y int) int {
	for y != 0 {
		t := y
		y = x % y
		x = t
	}
	return x
}
