package day17

import (
	"container/heap"
	"github.com/leandrorondon/advent-of-code/pkg/priorityqueue"
	"image"
	"math"
)

const (
	up = iota
	down
	left
	right
)

var turn = map[int]int{
	right: down,
	down:  right,
	left:  up,
	up:    left,
}

var dirs = []image.Point{
	{0, -1},
	{0, 1},
	{-1, 0},
	{1, 0},
}

type Step struct {
	Pos  image.Point
	Dir  int
	Loss int
}

func valid(arr [][]int, p image.Point) bool {
	return p.X >= 0 && p.X < len(arr) && p.Y > 0 && p.Y < len(arr[0])
}

func MinLoss(arr [][]int, mmin, mmax int) int {
	start := image.Point{0, 0}
	end := image.Point{len(arr) - 1, len(arr[0]) - 1}
	rows := len(arr)
	cols := len(arr[0])
	visited := make([][]int8, rows)
	for i := range visited {
		visited[i] = make([]int8, cols)
	}

	q := priorityqueue.New(func(s1, s2 Step) bool { return s1.Loss < s2.Loss })

	heap.Push(q, Step{Pos: start, Dir: right, Loss: 0})
	heap.Push(q, Step{Pos: start, Dir: down, Loss: 0})

	for q.Len() > 0 {
		step := heap.Pop(q).(Step)

		if step.Pos == end {
			return step.Loss
		}

		if visited[step.Pos.X][step.Pos.Y]&(1<<step.Dir) != 0 {
			continue
		}
		visited[step.Pos.X][step.Pos.Y] |= 1 << step.Dir

		for i := -mmax; i <= mmax; i++ {
			if i > -mmin && i < mmin {
				continue
			}

			n := step.Pos.Add(dirs[step.Dir].Mul(i))
			if !valid(arr, n) {
				continue
			}

			loss := 0
			s := int(math.Copysign(1, float64(i)))
			for j := s; j != i+s; j += s {
				p := step.Pos.Add(dirs[step.Dir].Mul(j))
				loss += arr[p.X][p.Y]
			}

			heap.Push(q, Step{n, turn[step.Dir], step.Loss + loss})
		}
	}
	return -1
}
