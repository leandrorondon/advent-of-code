package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Shape int

const (
	Rock Shape = iota
	Paper
	Scissors
)

var mapToShape = map[string]Shape{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

var shapeScore = map[Shape]int{
	Rock:     1,
	Paper:    2,
	Scissors: 3,
}

var shapeBeats = map[Shape]Shape{
	Rock:     Scissors,
	Paper:    Rock,
	Scissors: Paper,
}

var shapeLoses = map[Shape]Shape{
	Scissors: Rock,
	Rock:     Paper,
	Paper:    Scissors,
}

type Outcome int

const (
	Lose Outcome = iota
	Draw
	Win
)

var mapToOutcome = map[string]Outcome{
	"X": Lose,
	"Y": Draw,
	"Z": Win,
}

var outcomeScore = map[Outcome]int{
	Lose: 0,
	Draw: 3,
	Win:  6,
}

func main() {
	t := time.Now()
	b, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(b), "\n")

	var score1, score2 int
	for _, line := range lines {
		if line == "" {
			continue
		}
		s := strings.Split(line, " ")
		opponent := mapToShape[s[0]]
		player := mapToShape[s[1]]
		outcome := mapToOutcome[s[1]]

		score1 += roundOutcome(player, opponent) + shapeScore[player]
		score2 += shapeScore[findShape(opponent, outcome)] + outcomeScore[outcome]
	}

	took := time.Now().Sub(t)
	fmt.Println("Score 1:", score1)
	fmt.Println("Score 2:", score2)
	fmt.Printf("(took %v)\n", took)
}

func roundOutcome(player, opponent Shape) int {
	if player == opponent {
		return 3
	}

	if shapeBeats[player] == opponent {
		return 6
	}

	return 0
}

func findShape(opponent Shape, outcome Outcome) Shape {
	switch outcome {
	case Win:
		return shapeLoses[opponent]
	case Lose:
		return shapeBeats[opponent]
	default:
		// Draw
		return opponent
	}
}
