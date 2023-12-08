package day03

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/leandrorondon/advent-of-code/pkg/math"
)

type symbol struct {
	Value           string
	Row             int
	Col             int
	AdjacentNumbers []int
}

func (s symbol) IsNumber() bool {
	return unicode.IsDigit(rune(s.Value[0]))
}

func (s symbol) IsGear() bool {
	return s.Value == "*" && len(s.AdjacentNumbers) == 2
}

func (s symbol) GearRatio() int {
	if !s.IsGear() {
		return 0
	}

	return s.AdjacentNumbers[0] * s.AdjacentNumbers[1]
}

func (s symbol) HasAdjacentSymbol(symbols [][]*symbol) bool {
	// Check previous row
	if s.Row > 0 && s.hasAdjacentInRow(symbols[s.Row-1]) {
		return true
	}

	// Check current row
	if s.hasAdjacentInRow(symbols[s.Row]) {
		return true
	}

	// Check next row
	if s.Row+1 < len(symbols) && s.hasAdjacentInRow(symbols[s.Row+1]) {
		return true
	}

	return false
}

func (s symbol) hasAdjacentInRow(row []*symbol) bool {
	for i, sym := range row {
		if sym.IsNumber() {
			continue
		}

		if s.Row == sym.Row && s.Col == sym.Col {
			continue
		}

		if math.BetweenInclusive(sym.Col, s.Col-1, s.Col+len(s.Value)) {
			v, _ := strconv.Atoi(s.Value)
			row[i].AdjacentNumbers = append(row[i].AdjacentNumbers, v)
			return true
		}

		// past number
		if sym.Col > s.Col+len(s.Value) {
			break
		}
	}

	return false
}

func ParseLines(s string) [][]*symbol {
	lines := strings.Split(s, "\n")
	regex := regexp.MustCompile(`(\d+|[^.])`)
	var symbols [][]*symbol

	for i, line := range lines {
		matches := regex.FindAllStringSubmatchIndex(line, -1)

		var symbolsLine []*symbol

		for _, match := range matches {
			start := match[0]
			end := match[1]
			value := line[start:end]

			symbolsLine = append(symbolsLine, &symbol{
				Value: value,
				Row:   i,
				Col:   start,
			})
		}

		symbols = append(symbols, symbolsLine)
	}

	return symbols
}
