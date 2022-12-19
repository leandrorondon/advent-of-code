package tetris

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/leandrorondon/advent-of-code/pkg/math"
)

type Chamber struct {
	Height        int
	Width         int
	chamber       []int
	current       *Rock
	currentInLine []int
	jetGen        *JetGenerator
	rockGen       *RockGenerator
	print         bool
	leftLimit     int
}

func NewChamber(width int, jetGen *JetGenerator, rockGen *RockGenerator, print bool) *Chamber {
	c := &Chamber{
		Width:     width,
		leftLimit: 1 << (width + 1),
		print:     print,
		rockGen:   rockGen,
		jetGen:    jetGen,
	}

	return c
}

func (c *Chamber) SpawnRock() {
	c.current = c.rockGen.Generate()
	c.current.Position.X = 2
	c.current.Position.Y = c.Height + 3

	c.currentInLine = make([]int, c.current.Rows)
	for i := 0; i < c.current.Rows; i++ {
		c.currentInLine[i] = c.current.Pattern[i] << (c.Width - c.current.Cols - 2)
	}
}

func (c *Chamber) SettleCurrentRock() {
	if c.current == nil {
		return
	}

	rock := c.current
	inLine := c.currentInLine
	c.Height = math.Max(c.Height, rock.Position.Y+rock.Rows)
	rock.Rest = true
	c.current = nil
	c.currentInLine = nil

	for i := rock.Position.Y; i < rock.Position.Y+rock.Rows; i++ {
		if i == len(c.chamber) {
			c.chamber = append(c.chamber, 0)
		}

		c.chamber[i] |= inLine[i-rock.Position.Y]
	}
}

func (c *Chamber) Print() {
	if c.currentInLine != nil {
		for y := len(c.currentInLine) - 1; y >= 0; y-- {
			fmt.Printf("%4d |%s|\n", c.current.Position.Y+y, c.toLine(c.currentInLine[y], "@"))
		}

		for y := c.current.Position.Y; y > len(c.chamber); y-- {
			fmt.Printf("%4d |%s|\n", y-1, c.toLine(0, "."))
		}
	}

	if len(c.chamber) > 0 {
		for y := len(c.chamber) - 1; y >= 0; y-- {
			fmt.Printf("%4d |%s|\n", y, c.toLine(c.chamber[y], "#"))
		}
	}
	fmt.Printf("   F +%s+\n", strings.Repeat("-", c.Width))
}

func (c *Chamber) Tick() {
	if c.current == nil {
		return
	}

	// Fall
	if c.current.Position.Y > c.Height {
		c.current.Position.Y--
		return
	}

	if c.current.Position.Y == 0 {
		c.SettleCurrentRock()
		return
	}

	linesToCompare := c.Height - c.current.Position.Y + 1
	linesToCompare = math.Min(c.current.Rows, linesToCompare)

	for i := 0; i < linesToCompare; i++ {
		if c.currentInLine[i]&c.chamber[c.current.Position.Y-1+i] > 0 {
			c.SettleCurrentRock()
			return
		}
	}

	c.current.Position.Y--

	return
}

func (c *Chamber) CurrentRockResting() bool {
	if c.current == nil {
		return true
	}

	return c.current.Rest
}

func (c *Chamber) Jet() {
	jet := c.jetGen.Generate()

	switch jet {
	case '>':
		c.jetRight()
	case '<':
		c.jetLeft()
	default:
		panic(string(jet))
	}
}

func (c *Chamber) toLine(n int, r string) string {
	s := strconv.FormatInt(int64(n), 2)
	v, _ := strconv.Atoi(s)
	l := strings.Repeat("0", c.Width-len(s))
	s = fmt.Sprintf("%s%d", l, v)
	s = strings.ReplaceAll(s, "1", r)
	s = strings.ReplaceAll(s, "0", ".")
	return s
}

func (c *Chamber) jetRight() {
	newLines := make([]int, len(c.currentInLine))
	for i, l := range c.currentInLine {
		if l&1 > 0 {
			return
		}

		newL := l >> 1
		if c.current.Position.Y+i < c.Height {
			if newL&c.chamber[c.current.Position.Y+i] > 0 {
				return
			}
		}

		newLines[i] = newL
	}
	c.currentInLine = newLines
	c.current.Position.X++
}

func (c *Chamber) jetLeft() {
	newLines := make([]int, len(c.currentInLine))
	if c.current.Position.X == 0 {
		return
	}

	for i, l := range c.currentInLine {
		newL := l << 1

		if newL&c.leftLimit > 0 {
			return
		}

		if c.current.Position.Y+i < c.Height {
			if newL&c.chamber[c.current.Position.Y+i] > 0 {
				return
			}
		}

		newLines[i] = newL
	}
	c.currentInLine = newLines
	c.current.Position.X--
}

func (c *Chamber) NextRockIndex() int {
	return c.rockGen.Next()
}

func (c *Chamber) NextJetIndex() int {
	return c.jetGen.Next()
}
