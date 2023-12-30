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

func (m *Map) Cycle(n int) {
	loadMap := make(map[string]int)
	useMap := true

	for i := 0; i < n; i++ {
		m.TiltNorth()
		m.TiltWest()
		m.TiltSouth()
		m.TiltEast()

		if useMap {
			s := m.str()
			last, ok := loadMap[s]
			if !ok {
				loadMap[s] = i
				continue
			}
			useMap = false
			cycle := i - last
			i = n - (n-i)%cycle
		}
	}
}

func (m *Map) str() string {
	s := ""
	for i := range m.m {
		s += string(m.m[i])
	}
	return s
}

func (m *Map) TiltNorth() {
	for c := range m.m[0] {
		m.rollColNorth(c)
	}
}

func (m *Map) TiltSouth() {
	for c := range m.m[0] {
		m.rollColSouth(c)
	}
}

func (m *Map) TiltWest() {
	for r := range m.m {
		m.rollRowWest(r)
	}
}

func (m *Map) TiltEast() {
	for r := range m.m {
		m.rollRowEast(r)
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

func (m *Map) emptySpaceNorth(col int, start int) (int, bool) {
	for i := start; i >= 0; i-- {
		if m.m[i][col] == empty {
			return i, true
		}
	}

	return 0, false
}

func (m *Map) emptySpaceWest(row int, start int) (int, bool) {
	for i := start; i >= 0; i-- {
		if m.m[row][i] == empty {
			return i, true
		}
	}

	return 0, false
}

func (m *Map) emptySpaceEast(row int, start int) (int, bool) {
	for i := start; i < len(m.m[row]); i++ {
		if m.m[row][i] == empty {
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

func (m *Map) rollColSouth(c int) {
	rocksToRoll := 0
	rollTo, ok := m.emptySpaceNorth(c, len(m.m)-1)
	if !ok {
		return
	}

	for r := rollTo - 1; r >= 0; {
		if m.m[r][c] == rock {
			rocksToRoll++
			m.m[r][c] = empty
			r--
			continue
		}

		if m.m[r][c] == empty {
			r--
			continue
		}

		if m.m[r][c] == cube {
			m.rollRocksSouth(c, rollTo, rocksToRoll)
			rocksToRoll = 0
			rollTo, ok = m.emptySpaceNorth(c, r-1)
			if !ok {
				break
			}
			r = rollTo - 1
		}
	}

	m.rollRocksSouth(c, rollTo, rocksToRoll)
}

func (m *Map) rollRocksSouth(c, rollTo, rocksToRoll int) {
	if rocksToRoll == 0 {
		return
	}

	for i := 0; i < rocksToRoll; i++ {
		m.m[rollTo-i][c] = rock
	}
}

func (m *Map) rollRowWest(r int) {
	rocksToRoll := 0
	rollTo, ok := m.emptySpaceEast(r, 0)
	if !ok {
		return
	}

	for c := rollTo + 1; c < len(m.m[r]); {
		if m.m[r][c] == rock {
			rocksToRoll++
			m.m[r][c] = empty
			c++
			continue
		}

		if m.m[r][c] == empty {
			c++
			continue
		}

		if m.m[r][c] == cube {
			m.rollRocksWest(r, rollTo, rocksToRoll)
			rocksToRoll = 0
			rollTo, ok = m.emptySpaceEast(r, c+1)
			if !ok {
				break
			}
			c = rollTo + 1
		}
	}

	m.rollRocksWest(r, rollTo, rocksToRoll)
}

func (m *Map) rollRocksWest(r, rollTo, rocksToRoll int) {
	if rocksToRoll == 0 {
		return
	}

	for i := 0; i < rocksToRoll; i++ {
		m.m[r][rollTo+i] = rock
	}
}

func (m *Map) rollRowEast(r int) {
	rocksToRoll := 0
	rollTo, ok := m.emptySpaceWest(r, len(m.m[r])-1)
	if !ok {
		return
	}

	for c := rollTo - 1; c >= 0; {
		if m.m[r][c] == rock {
			rocksToRoll++
			m.m[r][c] = empty
			c--
			continue
		}

		if m.m[r][c] == empty {
			c--
			continue
		}

		if m.m[r][c] == cube {
			m.rollRocksEast(r, rollTo, rocksToRoll)
			rocksToRoll = 0
			rollTo, ok = m.emptySpaceWest(r, c-1)
			if !ok {
				break
			}
			c = rollTo - 1
		}
	}

	m.rollRocksEast(r, rollTo, rocksToRoll)
}

func (m *Map) rollRocksEast(r, rollTo, rocksToRoll int) {
	if rocksToRoll == 0 {
		return
	}

	for i := 0; i < rocksToRoll; i++ {
		m.m[r][rollTo-i] = rock
	}
}

func (m *Map) Print() {
	for i := range m.m {
		fmt.Println(string(m.m[i]))
	}
	fmt.Println()
}
