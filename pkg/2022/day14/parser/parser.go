package parser

import (
	"strconv"
	"strings"

	"github.com/leandrorondon/advent-of-code/pkg/2022/day14/grid"
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

func (p *Parser) ScanRocks() []grid.Path {
	var paths []grid.Path

	for p.scanner.Scan() {
		text := p.scanner.Text()
		if text == "" {
			continue
		}

		var path grid.Path
		coords := strings.Split(text, " -> ")
		for _, c := range coords {
			xy := strings.Split(c, ",")
			x, _ := strconv.Atoi(xy[0])
			y, _ := strconv.Atoi(xy[1])
			path = append(path, grid.Coordinate{X: x, Y: y})
		}

		paths = append(paths, path)
	}
	return paths
}
