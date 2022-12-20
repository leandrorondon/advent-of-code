package day18

import (
	"fmt"
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

func (p *Parser) ScanDroplets() []XYZ {
	var droplets []XYZ

	for p.scanner.Scan() {
		text := p.scanner.Text()
		if text == "" {
			continue
		}

		var droplet XYZ
		fmt.Sscanf(text, "%d,%d,%d", &droplet.X, &droplet.Y, &droplet.Z)
		droplets = append(droplets, droplet)
	}

	return droplets
}
