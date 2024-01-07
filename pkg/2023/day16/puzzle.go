package day16

import (
	"fmt"
	"github.com/leandrorondon/advent-of-code/pkg/bits"
	"github.com/leandrorondon/advent-of-code/pkg/math"
	"github.com/leandrorondon/advent-of-code/pkg/stack"
	"strings"
)

type Mirror struct {
	R             int
	C             int
	Type          byte
	left          *Mirror
	right         *Mirror
	up            *Mirror
	down          *Mirror
	receivedLeft  bool
	receivedRight bool
	receivedUp    bool
	receivedDown  bool
}

func (m *Mirror) Print() {
	if m == nil {
		return
	}
	fmt.Printf("Mirror at %d,%d: %s\n", m.R, m.C, string(m.Type))
	if m.left != nil {
		fmt.Printf("  left: %d,%d %s\n", m.left.R, m.left.C, string(m.left.Type))
	}
	if m.right != nil {
		fmt.Printf("  right: %d,%d %s\n", m.right.R, m.right.C, string(m.right.Type))
	}
	if m.up != nil {
		fmt.Printf("  up: %d,%d %s\n", m.up.R, m.up.C, string(m.up.Type))
	}
	if m.down != nil {
		fmt.Printf("  down: %d,%d %s\n", m.down.R, m.down.C, string(m.down.Type))
	}
}

type Map struct {
	mirrors   [][]*Mirror
	energised []*bits.Bits
}

type Direction int

const (
	down Direction = iota
	up
	left
	right
)

type Step struct {
	Mirror *Mirror
	Coming Direction
}

func (m *Map) Reset() {
	for i := range m.energised {
		m.energised[i].Reset()
	}
	for r := range m.mirrors {
		for c := range m.mirrors[r] {
			if m.mirrors[r][c] != nil {
				m.mirrors[r][c].receivedDown = false
				m.mirrors[r][c].receivedUp = false
				m.mirrors[r][c].receivedLeft = false
				m.mirrors[r][c].receivedRight = false
			}
		}
	}
}

func (m *Map) Ray(r, c int, coming Direction) int {
	m.Reset()

	st := stack.New[*Step]()
	var initial *Mirror
	switch coming {
	case left:
		initial = m.initialLeft(r)
	case right:
		initial = m.initialRight(r)
	case up:
		initial = m.initialUp(c)
	case down:
		initial = m.initialDown(c)
	}

	if initial == nil {
		return m.countEnergised()
	}

	step := &Step{Mirror: initial, Coming: coming}
	st.Push(step)

	for step != nil {
		m.handleStep(st, step)
		step = st.Pop()
	}

	return m.countEnergised()
}

func (m *Map) FindMax() int {
	largest := 0

	for r := range m.mirrors {
		m.Ray(r, 0, left)
		count := m.countEnergised()
		m.Ray(r, 0, right)
		count2 := m.countEnergised()
		largest = max(largest, count, count2)
	}

	for c := range m.mirrors[0] {
		m.Ray(0, c, up)
		count := m.countEnergised()
		m.Ray(0, c, down)
		count2 := m.countEnergised()
		largest = max(largest, count, count2)
	}

	return largest
}

func (m *Map) handleStep(st *stack.Stack[*Step], step *Step) {
	switch step.Coming {
	case down:
		switch step.Mirror.Type {
		case '-':
			m.goLeftRight(step.Mirror, st)
		case '|':
			m.goUp(step.Mirror, st)
		case '/':
			m.goRight(step.Mirror, st)
		case '\\':
			m.goLeft(step.Mirror, st)
		}
	case up:
		switch step.Mirror.Type {
		case '-':
			m.goLeftRight(step.Mirror, st)
		case '|':
			m.goDown(step.Mirror, st)
		case '/':
			m.goLeft(step.Mirror, st)
		case '\\':
			m.goRight(step.Mirror, st)
		}
	case left:
		switch step.Mirror.Type {
		case '-':
			m.goRight(step.Mirror, st)
		case '|':
			m.goUpDown(step.Mirror, st)
		case '/':
			m.goUp(step.Mirror, st)
		case '\\':
			m.goDown(step.Mirror, st)
		}
	case right:
		switch step.Mirror.Type {
		case '-':
			m.goLeft(step.Mirror, st)
		case '|':
			m.goUpDown(step.Mirror, st)
		case '/':
			m.goDown(step.Mirror, st)
		case '\\':
			m.goUp(step.Mirror, st)
		}
	}
}

func (m *Map) goLeftRight(mirror *Mirror, st *stack.Stack[*Step]) {
	m.goLeft(mirror, st)
	m.goRight(mirror, st)
}

func (m *Map) goUpDown(mirror *Mirror, st *stack.Stack[*Step]) {
	m.goUp(mirror, st)
	m.goDown(mirror, st)
}

func (m *Map) goRight(mirror *Mirror, st *stack.Stack[*Step]) {
	if mirror.right != nil {
		if !mirror.right.receivedLeft {
			m.travel(mirror, mirror.right)
			st.Push(&Step{mirror.right, left})
			mirror.right.receivedLeft = true
		}
	} else {
		m.escapeRight(mirror)
	}
}

func (m *Map) goLeft(mirror *Mirror, st *stack.Stack[*Step]) {
	if mirror.left != nil {
		if !mirror.left.receivedRight {
			m.travel(mirror, mirror.left)
			st.Push(&Step{mirror.left, right})
			mirror.left.receivedRight = true
		}
	} else {
		m.escapeLeft(mirror)
	}
}

func (m *Map) goDown(mirror *Mirror, st *stack.Stack[*Step]) {
	if mirror.down != nil {
		if !mirror.down.receivedUp {
			m.travel(mirror, mirror.down)
			st.Push(&Step{mirror.down, up})
			mirror.down.receivedUp = true
		}
	} else {
		m.escapeDown(mirror)
	}
}

func (m *Map) goUp(mirror *Mirror, st *stack.Stack[*Step]) {
	if mirror.up != nil {
		if !mirror.up.receivedDown {
			m.travel(mirror, mirror.up)
			st.Push(&Step{mirror.up, down})
			mirror.up.receivedDown = true
		}
	} else {
		m.escapeUp(mirror)
	}
}

func (m *Map) initialLeft(r int) *Mirror {
	for i := 0; i < len(m.mirrors[r]); i++ {
		m.energised[r].Set(i)
		if m.mirrors[r][i] != nil {
			return m.mirrors[r][i]
		}
	}
	return nil
}

func (m *Map) initialRight(r int) *Mirror {
	for i := len(m.mirrors[r]) - 1; i >= 0; i-- {
		m.energised[r].Set(i)
		if m.mirrors[r][i] != nil {
			return m.mirrors[r][i]
		}
	}
	return nil
}

func (m *Map) initialUp(c int) *Mirror {
	for i := 0; i < len(m.mirrors); i++ {
		m.energised[i].Set(c)
		if m.mirrors[i][c] != nil {
			return m.mirrors[i][c]
		}
	}
	return nil
}

func (m *Map) initialDown(c int) *Mirror {
	for i := len(m.mirrors) - 1; i >= 0; i-- {
		m.energised[i].Set(c)
		if m.mirrors[i][c] != nil {
			return m.mirrors[i][c]
		}
	}
	return nil
}

func (m *Map) travel(from, to *Mirror) {
	minR := math.Min(from.R, to.R)
	maxR := math.Max(from.R, to.R)
	minC := math.Min(from.C, to.C)
	maxC := math.Max(from.C, to.C)

	for r := minR; r <= maxR; r++ {
		for c := minC; c <= maxC; c++ {
			m.energised[r].Set(c)
		}
	}

}

func (m *Map) escapeLeft(from *Mirror) {
	for c := from.C; c >= 0; c-- {
		m.energised[from.R].Set(c)
	}
}

func (m *Map) escapeRight(from *Mirror) {
	for c := from.C; c < len(m.mirrors[0]); c++ {
		m.energised[from.R].Set(c)
	}
}

func (m *Map) escapeUp(from *Mirror) {
	for r := from.R; r >= 0; r-- {
		m.energised[r].Set(from.C)
	}
}

func (m *Map) escapeDown(from *Mirror) {
	for r := from.R; r < len(m.mirrors); r++ {
		m.energised[r].Set(from.C)
	}
}

func (m *Map) countEnergised() int {
	sum := 0
	for i := range m.energised {
		sum += m.energised[i].OnesCount()
	}
	return sum
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func (m *Map) PrintEnergised() {
	for i := range m.energised {
		fmt.Println(strings.ReplaceAll(strings.ReplaceAll(Reverse(m.energised[i].String()), "0", "."), "1", "#"))
	}
}
