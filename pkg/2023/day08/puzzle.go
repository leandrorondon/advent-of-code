package day08

import (
	"fmt"

	"github.com/leandrorondon/advent-of-code/pkg/math"
)

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

func (seq *Sequence) Next() (int, int) {
	next := seq.Current
	seq.Current++
	if seq.Current == len(seq.Instructions) {
		seq.Current = 0
	}

	return next, seq.Instructions[next]
}

func (seq *Sequence) Reset() {
	seq.Current = 0
}

type Node struct {
	Value     string
	Neighbors []*Node
}

func (n *Node) Steps(seq *Sequence, comp func(string) bool) (int, *Node) {
	nn := n
	steps := 0
	next := 0
	for !comp(nn.Value) || steps == 0 {
		_, next = seq.Next()
		nn = nn.Neighbors[next]
		steps++
	}

	return steps, nn
}

type NodeList []*Node

func (nl NodeList) FindPattern(seq *Sequence) []int {
	var repeat []int
	for i := range nl {
		seq.Reset()
		steps, node := nl[i].Steps(seq, CompEndingFn('Z'))
		stepsrep, _ := node.Steps(seq, CompFn(node.Value))
		if steps != stepsrep {
			panic("expected pattern")
		}

		repeat = append(repeat, stepsrep)
	}
	return repeat
}

func CompFn(dest string) func(s string) bool {
	return func(s string) bool { return s == dest }
}
func CompEndingFn(dest uint8) func(s string) bool {
	return func(s string) bool { return s[2] == dest }
}

func (nl NodeList) Steps(seq *Sequence) int {
	repeat := nl.FindPattern(seq)

	lcm := repeat[0]
	for i := 1; i < len(repeat); i++ {
		lcm = math.LCM(lcm, repeat[i])
	}

	fmt.Println(lcm)
	return lcm
}

func (nl NodeList) Print() {
	for i := range nl {
		fmt.Printf("%s ", nl[i].Value)
	}
	fmt.Printf("\n")
}
