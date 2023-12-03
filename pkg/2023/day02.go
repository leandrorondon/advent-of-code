package puzzles2023

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var maxSet = gameSet{R: 12, G: 13, B: 14}

func Day02(file string) error {
	t := time.Now()
	b, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	lines := strings.Split(string(b), "\n")

	sum := 0

game:
	for _, line := range lines {
		game := parseLine(line)

		for _, set := range game.Sets {
			if !set.isPossible(maxSet) {
				continue game
			}
		}

		// is possible
		sum += game.ID
	}

	took := time.Now().Sub(t)
	fmt.Println("Part 1:", sum)
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

func parseSets(s string) []gameSet {
	ss := strings.Split(s, ";")

	sets := make([]gameSet, len(ss))

	for i, set := range ss {
		sets[i] = parseSet(set)
	}

	return sets
}

func parseSet(s string) gameSet {
	ss := strings.Split(s, ",")

	var set gameSet

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

type gameSet struct {
	R int
	G int
	B int
}

func (s gameSet) isPossible(max gameSet) bool {
	return s.R <= max.R && s.G <= max.G && s.B <= max.B
}

type rgbGame struct {
	ID   int
	Sets []gameSet
}
