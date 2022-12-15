package parser

import (
	"fmt"
	"github.com/leandrorondon/advent-of-code/pkg/2022/day15/grid"
)

type Parser struct {
	scanner Scanner
}

type Scanner interface {
	Scan() bool
	Text() string
}

func NewParser(s Scanner) *Parser {
	return &Parser{
		scanner: s,
	}
}

func (p *Parser) ScanPairs() []grid.SensonBeaconPair {
	var pairs []grid.SensonBeaconPair

	for p.scanner.Scan() {
		text := p.scanner.Text()
		if text == "" {
			continue
		}

		var sx, sy, bx, by int
		fmt.Sscanf(text, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		pair := grid.SensonBeaconPair{
			Sensor: grid.Coordinate{X: sx, Y: sy},
			Beacon: grid.Coordinate{X: bx, Y: by},
		}

		pairs = append(pairs, pair)
	}

	return pairs
}
