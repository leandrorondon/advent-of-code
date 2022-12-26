package day21

type OpValue struct {
	Name     string
	Operator string
	Left     bool
	Value    int
}

type OpSeq []OpValue

func (ops OpSeq) Reverse(input int) int {
	v := input
	for i := len(ops) - 1; i >= 0; i-- {
		v = ops.inverseOperation(ops[i].Value, v, ops[i].Operator, ops[i].Left)
	}
	return v
}

func (ops OpSeq) inverseOperation(v1, v2 int, op string, left bool) int {
	switch op {
	case "+":
		return v2 - v1
	case "-":
		if left {
			return v1 - v2
		}
		return v1 + v2
	case "*":
		return v2 / v1
	case "/":
		if left {
			return v1 / v2
		}
		return v1 * v2
	}
	return 0
}

func inverse_eval(a, b int, op string, left bool) int {
	var res int
	if op == "+" {
		res = b - a
	} else if op == "-" {
		if left { // a on left of -
			res = (b - a) / -1
		} else {
			res = b + a
		}
	} else if op == "*" {
		res = b / a
	} else if op == "/" {
		if left { // a on left of /
			res = a / b
		} else {
			res = b * a
		}
	}
	return res
}

func FindEquality(monkeys map[string]*Monkey, root, target string) int {
	current := monkeys[monkeys[target].Parent]
	prev := target
	var ops OpSeq

	for current.Name != root {
		var op OpValue
		op.Name = current.Name
		op.Operator = current.Operator

		var v int
		if monkeys[current.M1].Name == prev {
			v = monkeys[current.M2].Say(monkeys)
		} else {
			op.Left = true
			v = monkeys[current.M1].Say(monkeys)
		}
		op.Value = v
		ops = append(ops, op)

		prev = current.Name
		current = monkeys[current.Parent]
	}

	var equality int
	if monkeys[root].M1 == ops[len(ops)-1].Name {
		equality = monkeys[monkeys[root].M2].Say(monkeys)
	} else {
		equality = monkeys[monkeys[root].M1].Say(monkeys)
	}

	return ops.Reverse(equality)
}
