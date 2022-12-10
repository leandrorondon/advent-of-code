package puzzles2022

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func Day09(file string) error {
	t := time.Now()
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	start := Coord{0, 0}

	rope1 := make([]Coord, 2)
	visited1 := make(map[Coord]bool)
	visited1[start] = true

	rope2 := make([]Coord, 10)
	visited2 := make(map[Coord]bool)
	visited2[start] = true

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		if len(s) == 0 {
			continue
		}

		steps, _ := strconv.Atoi(s[1])
		for i := 0; i < steps; i++ {
			step(rope1, visited1, s[0])
			step(rope2, visited2, s[0])
		}

	}

	took := time.Now().Sub(t)
	fmt.Println("Part 1:", len(visited1))
	fmt.Println("Part 2:", len(visited2))
	fmt.Printf("(took %v)\n", took)

	return nil
}

type Coord struct {
	X int
	Y int
}

func (c *Coord) Move(dx, dy int) {
	c.X += dx
	c.Y += dy
}

func (c *Coord) Touching(other Coord) bool {
	return math.Abs((float64)(c.X-other.X)) <= 1 && math.Abs((float64)(c.Y-other.Y)) <= 1
}

func (c *Coord) Follow(other Coord) {
	dx := other.X - c.X
	dy := other.Y - c.Y

	if dx > 0 {
		c.X++
	} else if dx < 0 {
		c.X--
	}

	if dy > 0 {
		c.Y++
	} else if dy < 0 {
		c.Y--
	}
}

var dc = map[string]Coord{
	"R": {1, 0},
	"L": {-1, 0},
	"U": {0, 1},
	"D": {0, -1},
}

func step(rope []Coord, visited map[Coord]bool, dir string) {
	rope[0].Move(dc[dir].X, dc[dir].Y)

	for i := 1; i < len(rope); i++ {
		if rope[i].Touching(rope[i-1]) {
			break
		}

		rope[i].Follow(rope[i-1])
	}

	visited[rope[len(rope)-1]] = true
}

func print(rope []Coord, visited map[Coord]bool, from, to int) {
	fmt.Printf("\n")

	for i := to; i >= from; i-- {
		lineVisited := ""
		linePos := ""
		for j := from; j <= to; j++ {
			c := Coord{j, i}

			if visited[c] {
				lineVisited += "#"
			} else {
				if i == 0 && j == 0 {
					lineVisited += "+"
				} else if i == 0 {
					lineVisited += "-"
				} else if j == 0 {
					lineVisited += "|"
				} else {
					lineVisited += "."
				}
			}

			p := "."
			if i == 0 && j == 0 {
				p = "+"
			} else if i == 0 {
				p = "-"
			} else if j == 0 {
				p = "|"
			}

			if c.X == 0 && c.Y == 0 {
				p = "s"
			}

			if c.X == rope[0].X && c.Y == rope[0].Y {
				p = "H"
			} else {
				for k := 1; k < len(rope); k++ {
					if c.X == rope[k].X && c.Y == rope[k].Y {
						p = strconv.Itoa(k)
						break
					}
				}
			}
			linePos += p
		}
		fmt.Printf("%s   %s\n", lineVisited, linePos)
	}
}
