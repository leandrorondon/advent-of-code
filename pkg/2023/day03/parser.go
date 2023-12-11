package day03

import (
	"regexp"
	"strings"
)

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
