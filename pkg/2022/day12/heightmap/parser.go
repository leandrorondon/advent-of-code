package heightmap

import (
	"strings"
)

type Parser struct {
	scanner Scanner
}

type Scanner interface {
	Scan() bool
	Text() string
}

func NewParser(s Scanner) Parser {
	return Parser{
		scanner: s,
	}
}

func (p Parser) Parse() HeightMap {
	var m [][]byte
	i := 0
	var start, end Position
	for p.scanner.Scan() {
		line := p.scanner.Text()
		m = append(m, []byte(line))
		if s := strings.Index(line, string(START)); s >= 0 {
			start = Position{i, s}
		}
		if s := strings.Index(line, string(END)); s >= 0 {
			end = Position{i, s}
		}
		i++
	}

	m[start.X][start.Y] = 'a'
	m[end.X][end.Y] = 'z'

	return HeightMap{m, start, end}
}
