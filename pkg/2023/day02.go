package puzzles2023

import (
	"fmt"
	"github.com/leandrorondon/advent-of-code/pkg/ternary"
	"os"
	"strconv"
	"strings"
	"time"
)

var maxSet = rgb{R: 12, G: 13, B: 14}

func Day02(file string) error {
	t := time.Now()
	b, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	lines := strings.Split(string(b), "\n")

	sum := 0

game1:
	for _, line := range lines {
		game := parseLine(line)

		for _, set := range game.Sets {
			if !set.isPossible(maxSet) {
				continue game1
			}
		}

		// is possible
		sum += game.ID
	}

	took := time.Now().Sub(t)
	fmt.Println("Part 1:", sum)
	fmt.Printf("(took %v)\n", took)

	sum = 0

	for _, line := range lines {
		game := parseLine(line)

		var minSet rgb
		for _, set := range game.Sets {
			minSet.R = ternary.If(set.R > minSet.R, set.R, minSet.R)
			minSet.G = ternary.If(set.G > minSet.G, set.G, minSet.G)
			minSet.B = ternary.If(set.B > minSet.B, set.B, minSet.B)
		}

		power := minSet.R * minSet.G * minSet.B
		sum += power
	}

	took = time.Now().Sub(t)
	fmt.Println("Part 2:", sum)
	fmt.Printf("(took %v)\n", took)

	return nil
}

func parseLine(s string) rgbGame {
	ss := strings.Split(s, ":")

	return rgbGame{
		ID:   parseGameID(ss[0]),
		Sets: parseSets(ss[1]),
	}
}

func parseGameID(s string) int {
	ss := strings.Split(s, " ")
	i, _ := strconv.Atoi(ss[1])
	return i
}

func parseSets(s string) []rgb {
	ss := strings.Split(s, ";")

	sets := make([]rgb, len(ss))

	for i, set := range ss {
		sets[i] = parseSet(set)
	}

	return sets
}

func parseSet(s string) rgb {
	ss := strings.Split(s, ",")

	var set rgb

	for _, colourSet := range ss {
		colour := strings.Split(strings.TrimSpace(colourSet), " ")
		switch colour[1] {
		case "red":
			set.R, _ = strconv.Atoi(colour[0])
		case "green":
			set.G, _ = strconv.Atoi(colour[0])
		case "blue":
			set.B, _ = strconv.Atoi(colour[0])
		}
	}

	return set
}

type rgb struct {
	R int
	G int
	B int
}

func (s rgb) isPossible(max rgb) bool {
	return s.R <= max.R && s.G <= max.G && s.B <= max.B
}

type rgbGame struct {
	ID   int
	Sets []rgb
}
