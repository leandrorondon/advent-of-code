package day21

type Monkey struct {
	Name     string
	Value    int
	IsKnown  bool
	Operator string
	M1       string
	M2       string
	Parent   string
}

func (m Monkey) Say(monkeys map[string]*Monkey) int {
	if !m.IsKnown {
		v1 := monkeys[m.M1].Say(monkeys)
		v2 := monkeys[m.M2].Say(monkeys)

		monkeys[m.M1].Parent = m.Name
		monkeys[m.M2].Parent = m.Name
		return m.operation(v1, v2)
	}

	return m.Value
}

func (m Monkey) operation(v1, v2 int) int {
	switch m.Operator {
	case "+":
		return v1 + v2
	case "-":
		return v1 - v2
	case "*":
		return v1 * v2
	case "/":
		return v1 / v2
	}
	return 0
}
