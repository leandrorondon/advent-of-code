package monkey

import (
	"fmt"
	"math/big"
)

func NewBig(m *Monkey) *BigMonkey {
	bm := &BigMonkey{
		id:           m.id,
		operator:     m.operator,
		operand:      m.operand,
		operandValue: int64(m.operandValue),
		monkeyTrue:   m.monkeyTrue,
		monkeyFalse:  m.monkeyFalse,
	}

	bm.divisor = big.NewInt(int64(m.divisor))

	for i := range m.items {
		item := big.NewInt(int64(m.items[i]))
		bm.items = append(bm.items, item)
	}

	return bm
}

type BigMonkey struct {
	id           int
	inspections  int
	items        []*big.Int
	operator     string
	operand      Operand
	operandValue int64
	adjustLevel  *big.Int
	divisor      *big.Int
	monkeyTrue   int
	monkeyFalse  int
}

func (m *BigMonkey) Turn(monkeys []*BigMonkey) {
	for i := range m.items {
		m.inspect(i)
		m.adjust(i)

		dest := m.monkeyFalse
		copy := big.NewInt(m.items[i].Int64())
		if copy.Mod(m.items[i], m.divisor).Int64() == 0 {
			dest = m.monkeyTrue
		}

		m.throw(i, dest, monkeys)
	}

	m.items = []*big.Int{}
}

func (m *BigMonkey) Receive(item *big.Int) {
	m.items = append(m.items, item)
}

func (m *BigMonkey) Print() {
	fmt.Printf("Monkey %d: ", m.id)
	for _, item := range m.items {
		fmt.Printf("%s ", item.String())
	}
	fmt.Printf("\n")
}

func (m *BigMonkey) PrintInspections() {
	fmt.Printf("Monkey %d inspected items %d times.\n", m.id, m.inspections)
}

func (m *BigMonkey) Inspections() int {
	return m.inspections
}

func (m *BigMonkey) SetAdjustLevel(level int) {
	m.adjustLevel = big.NewInt(int64(level))
}

func (m *BigMonkey) inspect(i int) {
	if m.operand == Value {
		v := big.NewInt(m.operandValue)
		m.operation(i, v)
	} else {
		m.operation(i, m.items[i])
	}

	m.inspections++
}

func (m *BigMonkey) operation(i int, v *big.Int) {
	switch m.operator {
	case "+":
		m.items[i].Add(m.items[i], v)
	case "-":
		m.items[i].Sub(m.items[i], v)
	case "*":
		m.items[i].Mul(m.items[i], v)
	case "/":
		m.items[i].Div(m.items[i], v)
	}
}

func (m *BigMonkey) adjust(i int) {
	m.items[i].Mod(m.items[i], m.adjustLevel)
}

func (m *BigMonkey) throw(i, dest int, monkeys []*BigMonkey) {
	monkeys[dest].Receive(m.items[i])
}
