package day13

type Patterns []Pattern

func (ps Patterns) Summarise() int {
	sum := 0
	for i := range ps {
		sum += ps[i].Summarise()
	}
	return sum
}

func (ps Patterns) SummariseSmudge() int {
	sum := 0
	for i := range ps {
		sum += ps[i].SummariseSmudge()
	}
	return sum
}

type Pattern struct {
	Rows    []int64
	RowSize int
	Cols    []int64
	ColSize int
}

func (p *Pattern) summarise(arr []int64, ignore int) int {
	for i := 1; i < len(arr); i++ {
		if p.mirrorPrevious(arr, i) && i != ignore {
			return i
		}
	}

	return 0
}

func (p *Pattern) Summarise() int {
	if summary := p.summarise(p.Cols, 0); summary > 0 {
		return summary
	}

	return p.summarise(p.Rows, 0) * 100
}

func (p *Pattern) summariseSmudge(arr []int64, size, mult, original int) int {
	for i := 0; i < len(arr); i++ {
		orig := arr[i]
		for b := 0; b < size; b++ {
			arr[i] = orig ^ (1 << b)
			summary := p.summarise(arr, original/mult)
			if summary > 0 {
				return summary * mult
			}
		}
		arr[i] = orig
	}

	return 0
}

func (p *Pattern) SummariseSmudge() int {
	original := p.Summarise()

	summary := p.summariseSmudge(p.Cols, p.ColSize, 1, original)
	if summary > 0 {
		return summary
	}

	summary = p.summariseSmudge(p.Rows, p.RowSize, 100, original)

	return summary
}

func (p *Pattern) mirrorPrevious(ss []int64, idx int) bool {
	prev := idx - 1
	for prev >= 0 && idx < len(ss) {
		if ss[prev] != ss[idx] {
			return false
		}
		prev--
		idx++
	}

	return true
}
