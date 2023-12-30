package day14

import (
	"fmt"
)

const (
	rock  = 'O'
	cube  = '#'
	empty = '.'
)

type Map struct {
	m [][]byte
}

func (m *Map) TiltNorth() {
	for c := range m.m[0] {
		m.rollColNorth(c)
	}
}

func (m *Map) Load() int {
	load := 0
	rows := len(m.m)
	for r := range m.m {
		for c := range m.m[r] {
			if m.m[r][c] == rock {
				load += rows - r
			}
		}
	}
	return load
}

func (m *Map) emptySpaceSouth(col int, start int) (int, bool) {
	for i := start; i < len(m.m); i++ {
		if m.m[i][col] == empty {
			return i, true
		}
	}

	return 0, false
}

func (m *Map) rollColNorth(c int) {
	rocksToRoll := 0
	rollTo, ok := m.emptySpaceSouth(c, 0)
	if !ok {
		return
	}

	for r := rollTo + 1; r < len(m.m); {
		if m.m[r][c] == rock {
			rocksToRoll++
			m.m[r][c] = empty
			r++
			continue
		}

		if m.m[r][c] == empty {
			r++
			continue
		}

		if m.m[r][c] == cube {
			m.rollRocksNorth(c, rollTo, rocksToRoll)
			rocksToRoll = 0
			rollTo, ok = m.emptySpaceSouth(c, r+1)
			if !ok {
				break
			}
			r = rollTo + 1
		}
	}

	m.rollRocksNorth(c, rollTo, rocksToRoll)
}

func (m *Map) rollRocksNorth(c, rollTo, rocksToRoll int) {
	if rocksToRoll == 0 {
		return
	}

	for i := 0; i < rocksToRoll; i++ {
		m.m[rollTo+i][c] = rock
	}
}

func (m *Map) Print() {
	for i := range m.m {
		fmt.Println(string(m.m[i]))
	}
	fmt.Println()
}
