package monkey

import (
	"fmt"
)

func New(id int) *Monkey {
	return &Monkey{
		id: id,
	}
}

type Operand int

const (
	Old Operand = iota
	Value
)

type Monkey struct {
	id           int
	inspections  int
	items        []int
	operator     string
	operand      Operand
	operandValue int
	adjustLevel  int
	divisor      int
	monkeyTrue   int
	monkeyFalse  int
}

func (m *Monkey) Turn(monkeys []*Monkey) {
	for i := range m.items {
		m.inspect(i)
		m.adjust(i)

		dest := m.monkeyFalse
		if m.items[i]%m.divisor == 0 {
			dest = m.monkeyTrue
		}

		m.throw(i, dest, monkeys)
	}

	m.items = []int{}
}

func (m *Monkey) Receive(item int) {
	m.items = append(m.items, item)
}

func (m *Monkey) Print() {
	fmt.Printf("Monkey %d: %v\n", m.id, m.items)
}

func (m *Monkey) PrintInspections() {
	fmt.Printf("Monkey %d inspected items %d times.\n", m.id, m.inspections)
}

func (m *Monkey) Divisor() int {
	return m.divisor
}

func (m *Monkey) Inspections() int {
	return m.inspections
}

func (m *Monkey) SetAdjustLevel(level int) {
	m.adjustLevel = level
}

func (m *Monkey) inspect(i int) {
	v := m.items[i]
	if m.operand == Value {
		v = m.operandValue
	}
	m.operation(i, v)
	m.inspections++
}

func (m *Monkey) operation(i, v int) {
	switch m.operator {
	case "+":
		m.items[i] += v
	case "-":
		m.items[i] -= v
	case "*":
		m.items[i] *= v
	case "/":
		m.items[i] /= v
	}
}

func (m *Monkey) adjust(i int) {
	m.items[i] /= m.adjustLevel
}

func (m *Monkey) throw(i, dest int, monkeys []*Monkey) {
	monkeys[dest].Receive(m.items[i])
}
