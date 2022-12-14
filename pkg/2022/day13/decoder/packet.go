package decoder

import "fmt"

type Packet []Value

func (p Packet) Compare(p2 Packet) int {
	var res int
	for i, v := range p {
		if i > len(p2)-1 {
			return 1
		}

		res = v.Compare(p2[i])
		if res != 0 {
			return res
		}
	}

	if len(p) < len(p2) {
		return -1
	}

	return 0
}

func (p Packet) String() string {
	if len(p) == 0 {
		return "[]"
	}

	s := "["
	for _, v := range p {
		s += fmt.Sprintf("%s,", v.String())
	}

	return s[:len(s)-1] + "]"
}
