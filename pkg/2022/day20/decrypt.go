package day20

import (
	"container/ring"
	"fmt"
)

func GrooveCoordinates(sl []*ring.Ring, zero int) (int, int, int) {
	n := len(sl) - 1
	for _, node := range sl {
		r := node.Prev()
		removed := r.Unlink(1)
		r = r.Move(removed.Value.(int) % n)
		r.Link(removed)
	}

	p1000 := sl[zero].Move(1000)
	p2000 := p1000.Move(1000)
	p3000 := p2000.Move(1000)

	return p1000.Value.(int), p2000.Value.(int), p3000.Value.(int)
}

func GrooveCoordinatesWithKey(sl []*ring.Ring, zero int, key int) (int, int, int) {

	for _, node := range sl {
		node.Value = node.Value.(int) * key
	}

	var g1, g2, g3 int
	for i := 0; i < 10; i++ {
		g1, g2, g3 = GrooveCoordinates(sl, zero)
	}

	return g1, g2, g3
}

func printRing(r *ring.Ring) {
	current := r
	for ; current.Next() != r; current = current.Next() {
		fmt.Printf("%d, ", current.Value.(int))
	}
	fmt.Printf("%d\n", current.Value.(int))
}
