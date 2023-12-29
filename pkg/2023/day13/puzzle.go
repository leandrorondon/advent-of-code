package day13

import "fmt"

type Patterns []Pattern

func (ps Patterns) Summarise() int {
	sum := 0
	for i := range ps {
		sum += ps[i].Summarise()
	}
	return sum
}

type Pattern struct {
	Rows    []string
	Columns []string
}

func (p *Pattern) Summarise() int {
	for i := 1; i < len(p.Columns); i++ {
		if p.mirrorPrevious(p.Columns, i) {
			return i
		}
	}

	for i := 1; i < len(p.Rows); i++ {
		if p.mirrorPrevious(p.Rows, i) {
			return i * 100
		}
	}

	return 0
}

func (p *Pattern) mirrorPrevious(ss []string, idx int) bool {
	ileft := idx - 1
	for ileft >= 0 && idx < len(ss) {
		if ss[ileft] != ss[idx] {
			return false
		}
		ileft--
		idx++
	}

	return true
}

func (p *Pattern) PrintHighlightRow(n int) {
	for i := range p.Rows {
		c := " "
		if n == i {
			c = "^"
		} else if n == i+1 {
			c = "v"
		}
		fmt.Printf("  %2d%s%s\n", i+1, c, p.Rows[i])
	}
	fmt.Println("")
}

func (p *Pattern) PrintHighlightCol(n int) {
	fmt.Print("  ")
	for i := range p.Columns {
		c := " "
		if n == i {
			c = "<"
		} else if n == i+1 {
			c = ">"
		}
		fmt.Printf("%s", c)
	}
	fmt.Printf("\n")
	for i := range p.Rows {
		fmt.Printf("  %s\n", p.Rows[i])
	}
	fmt.Println()
}
