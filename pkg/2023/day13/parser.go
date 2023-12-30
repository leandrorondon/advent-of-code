package day13

import (
	"log"
	"strconv"
	"strings"
)

func Parse(s string) Patterns {
	s = strings.ReplaceAll(s, "#", "1")
	s = strings.ReplaceAll(s, ".", "0")

	lines := strings.Split(s, "\n")

	var ps Patterns
	var p Pattern

	for i := range lines {
		if lines[i] == "" {
			p.Cols = buildCols(p.Rows, len(lines[i-1]))
			p.RowSize = len(p.Cols)
			p.ColSize = len(p.Rows)
			ps = append(ps, p)
			p = Pattern{}
			continue
		}

		n, err := strconv.ParseInt(lines[i], 2, 64)
		if err != nil {
			log.Fatal(err)
		}

		p.Rows = append(p.Rows, n)
	}
	p.Cols = buildCols(p.Rows, len(lines[len(lines)-1]))
	p.RowSize = len(p.Cols)
	p.ColSize = len(p.Rows)
	ps = append(ps, p)

	return ps
}

func buildCols(rows []int64, size int) []int64 {
	cols := make([]int64, size)
	for c := range cols {
		for r := range rows {
			rot := len(cols) - c - 1
			bit := int64(1 << rot)
			cols[c] += ((bit & rows[r]) >> rot) << (len(rows) - r - 1)
		}
	}
	return cols
}
