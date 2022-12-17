package parser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/leandrorondon/advent-of-code/pkg/2022/day16/valve"
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

func (p *Parser) ScanValves() map[string]*valve.Valve {
	m := make(map[string]*valve.Valve)

	r, err := regexp.Compile(`^Valve (?P<label>\w+) has flow rate=(?P<rate>\d+); tunnels? leads? to valves? (?P<leads>.+)`)
	if err != nil {
		fmt.Println(err)
		return m
	}

	id := 0
	for p.scanner.Scan() {
		text := p.scanner.Text()
		if text == "" {
			continue
		}

		match := r.FindStringSubmatch(text)
		if len(match) < 4 {
			continue
		}

		label := match[1]
		rate, _ := strconv.Atoi(match[2])
		leads := strings.Split(match[3], ", ")

		v, ok := m[label]
		if !ok {
			v = &valve.Valve{
				Label: label,
			}
			m[label] = v
		}
		v.ID = id
		id++

		var next []*valve.Valve
		for _, lbl := range leads {
			if existing, ok := m[lbl]; ok {
				next = append(next, existing)
				continue
			}

			newValve := &valve.Valve{Label: lbl}
			next = append(next, newValve)
			m[lbl] = newValve
		}

		v.Rate = rate
		v.Tunnels = next
	}

	return m
}
