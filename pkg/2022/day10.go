package puzzles2022

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func Day10(file string) error {
	t := time.Now()
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	cpu := NewCPU()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		cpu.exec(scanner.Text())
	}

	took := time.Now().Sub(t)
	fmt.Println("Sum of signal strengths:", cpu.strength)
	fmt.Println("CRT output:")
	cpu.CRT.print()
	fmt.Printf("(took %v)\n", took)

	return nil
}

func NewCPU() *CPU {
	return &CPU{X: 1}
}

type CPU struct {
	X        int
	cycles   int
	strength int
	CRT      CRT
}

func (cpu *CPU) exec(s string) {
	cmd := strings.Split(s, " ")

	switch cmd[0] {
	case "noop":
		cpu.cycle()
	case "addx":
		v, _ := strconv.Atoi(cmd[1])
		cpu.cycle()
		cpu.cycle()
		cpu.add(v)
	}
}

func (cpu *CPU) cycle() {
	cpu.CRT.render(cpu.cycles, cpu.X)
	cpu.cycles++
	if cpu.cycles%40 == 20 {
		cpu.strength += cpu.cycles * cpu.X
	}

}

func (cpu *CPU) add(v int) {
	cpu.X += v
}

type CRT struct {
	screen [6][40]byte
}

func (c *CRT) render(pos, sprite int) {
	row := pos / 40
	col := pos % 40

	if col >= sprite-1 && col <= sprite+1 {
		c.screen[row][col] = '#'
	} else {
		c.screen[row][col] = '.'
	}
}

func (c *CRT) print() {
	for _, r := range c.screen {
		for _, b := range r {
			fmt.Printf("%s", string(b))
		}
		fmt.Printf("\n")
	}
}
