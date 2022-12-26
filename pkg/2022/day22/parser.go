package day22

import (
	"os"
	"strconv"
	"strings"
)

type Parser struct {
	file string
}

func NewParser(f string) *Parser {
	return &Parser{
		file: f,
	}
}

func (p *Parser) Parse() ([][]byte, []Instruction) {
	b, err := os.ReadFile(p.file)
	if err != nil {
		return nil, nil
	}

	blocks := strings.Split(string(b), "\n\n")

	return p.parseMap(blocks[0]), p.parseInstructions(blocks[1])
}

func (p *Parser) parseMap(text string) [][]byte {
	lines := strings.Split(text, "\n")
	max := 0
	for _, l := range lines {
		if len(l) > max {
			max = len(l)
		}
	}

	m := make([][]byte, len(lines))
	for i := range m {
		m[i] = make([]byte, max)
	}

	for i, l := range lines {
		copy(m[i], l)
	}

	return m
}

func (p *Parser) parseInstructions(text string) []Instruction {
	var instructions []Instruction
	for i := 0; i < len(text); i++ {
		if text[i] == '\n' {
			break
		}

		if text[i] == 'L' || text[i] == 'R' {
			instructions = append(instructions, Instruction{Action: text[i]})
			continue
		}

		start := i
		for i+1 < len(text) && text[i+1] >= '0' && text[i+1] <= '9' {
			i++
		}
		steps, _ := strconv.Atoi(text[start : i+1])
		instructions = append(instructions, Instruction{Steps: steps})

	}
	return instructions
}
