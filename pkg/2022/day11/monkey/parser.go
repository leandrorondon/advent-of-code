package monkey

import (
	"fmt"
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

func (p Parser) ParseMonkey() (m *Monkey) {
	for p.scanner.Scan() {
		line := p.scanner.Text()
		if line == "" {
			break
		}

		switch {
		case strings.HasPrefix(line, "Monkey"):
			var id int
			fmt.Sscanf(line, "Monkey %d:", &id)
			m = New(id)
		case strings.Contains(line, "Starting items:"):
			p.parseItems(m, line)
		case strings.Contains(line, "Operation:"):
			p.parseOperation(m, line)
		case strings.Contains(line, "Test:"):
			p.parseTest(m, line)
		case strings.Contains(line, "If true:"),
			strings.Contains(line, "If false:"):
			p.parseDest(m, line)
		}
	}

	return m
}

func (p Parser) parseItems(m *Monkey, line string) {
	var items []int
	for _, item := range strings.Split(line[18:], ", ") {
		v, _ := strconv.Atoi(item)
		items = append(items, int(v))
	}
	m.items = items
}

func (p Parser) parseOperation(m *Monkey, line string) {
	s := strings.Split(strings.Split(line, " = ")[1], " ")
	if len(s) != 3 {
		return
	}

	m.operator = s[1]

	if s[2] == "old" {
		m.operand = Old
		return
	}
	m.operand = Value
	v, _ := strconv.Atoi(s[2])
	m.operandValue = int(v)
}

func (p Parser) parseTest(m *Monkey, line string) {
	s := strings.Split(line, " divisible by ")
	if len(s) != 2 {
		m.divisor = 1
		return
	}
	m.divisor, _ = strconv.Atoi(s[1])

}

func (p Parser) parseDest(m *Monkey, line string) {
	dest := &m.monkeyTrue
	if strings.Contains(line, "If false") {
		dest = &m.monkeyFalse
	}

	s := strings.Split(line, "monkey ")
	if len(s) != 2 {
		*dest = -1
		return
	}

	*dest, _ = strconv.Atoi(s[1])
}
