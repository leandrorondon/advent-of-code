package day08

import (
	"fmt"
	"strings"
)

func getOrCreate(m map[string]*Node, v string) *Node {
	n, ok := m[v]
	if ok {
		return n
	}
	n = &Node{
		Value:     v,
		Neighbors: make([]*Node, 2),
	}
	m[v] = n
	return n
}

func ParseMap(s string) (Sequence, *Node) {
	lines := strings.Split(s, "\n")

	var instr []int
	for _, r := range lines[0] {
		instr = append(instr, idx[r])
	}

	nodeMap := make(map[string]*Node)
	var initial *Node
	for i := 2; i < len(lines); i++ {
		var nodeValue, leftValue, rightValue string
		ss := strings.ReplaceAll(strings.ReplaceAll(lines[i], ",", ""), ")", "")
		_, _ = fmt.Sscanf(ss, "%s = (%s %s", &nodeValue, &leftValue, &rightValue)

		node := getOrCreate(nodeMap, nodeValue)
		node.Neighbors[Left] = getOrCreate(nodeMap, leftValue)
		node.Neighbors[Right] = getOrCreate(nodeMap, rightValue)

		if nodeValue == Initial {
			initial = node
		}
	}

	return Sequence{Instructions: instr}, initial
}
