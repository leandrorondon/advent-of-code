package day20

import (
	"container/ring"
	"os"
	"strconv"
	"strings"
)

type Parser struct {
	file string
}

type Scanner interface {
	Scan() bool
	Text() string
}

func NewParser(f string) *Parser {
	return &Parser{
		file: f,
	}
}

func (p *Parser) Parse() ([]*ring.Ring, int) {
	b, err := os.ReadFile(p.file)
	if err != nil {
		return nil, 0
	}
	lines := strings.Split(string(b), "\n")

	r := ring.New(len(lines) - 1)
	sl := make([]*ring.Ring, len(lines)-1)

	var zero int
	for i, line := range lines {
		if line == "" {
			continue
		}
		v, _ := strconv.Atoi(line)
		r.Value = v
		sl[i] = r
		if v == 0 {
			zero = i
		}
		r = r.Next()
	}

	return sl, zero
}
