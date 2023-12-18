package day08

var idx = map[rune]int{
	'L': 0,
	'R': 1,
}

const (
	Left    = 0
	Right   = 1
	Initial = "AAA"
	Final   = "ZZZ"
)

type Sequence struct {
	Current      int
	Instructions []int
}

func (seq *Sequence) Next() int {
	next := seq.Current
	seq.Current++
	if seq.Current == len(seq.Instructions) {
		seq.Current = 0
	}

	return seq.Instructions[next]
}

type Node struct {
	Value     string
	Neighbors []*Node
}

func (n *Node) Steps(dest string, seq *Sequence) int {
	if n.Value == dest {
		return 0
	}

	return n.Neighbors[seq.Next()].Steps(dest, seq) + 1
}
