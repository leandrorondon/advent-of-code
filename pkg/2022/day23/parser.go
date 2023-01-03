package day23

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

func (p *Parser) Parse() []Elf {
	var elfs []Elf

	row := 0
	for p.scanner.Scan() {
		line := p.scanner.Text()
		if line == "" {
			break
		}

		for col, c := range line {
			if c == '\n' {
				panic("panic")
			}

			if c == '.' {
				continue
			}

			elfs = append(elfs, Elf{Position: Coordinate{row, col}})
		}

		row++
	}

	return elfs
}
