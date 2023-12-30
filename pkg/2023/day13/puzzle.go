package day13

type Patterns []Pattern

func (ps Patterns) Summarise() int {
	sum := 0
	for i := range ps {
		sum += ps[i].Summarise()
	}
	return sum
}

type Pattern struct {
	Rows    []int64
	Columns []int64
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

func (p *Pattern) mirrorPrevious(ss []int64, idx int) bool {
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
