package day21

import (
	"strconv"
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

func (p *Parser) Parse() map[string]*Monkey {
	m := make(map[string]*Monkey)

	for p.scanner.Scan() {
		line := p.scanner.Text()
		if line == "" {
			break
		}

		var op, m1, m2 string
		var value int

		s := strings.Split(line, ": ")
		name := s[0]

		s = strings.Split(s[1], " ")

		known := true
		if len(s) == 1 {
			value, _ = strconv.Atoi(s[0])
		} else {
			m1 = s[0]
			op = s[1]
			m2 = s[2]
			known = false
		}

		m[name] = &Monkey{
			Name:     name,
			Value:    value,
			IsKnown:  known,
			M1:       m1,
			M2:       m2,
			Operator: op,
		}
	}

	return m
}
