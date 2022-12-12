package iterator

func New2D[T any](d [][]T) TwoD[T] {
	it := TwoD[T]{
		inited: false,
		data:   d,
		maxX:   len(d) - 1,
	}

	if len(d) > 0 {
		it.maxY = len(d[0]) - 1
	}

	return it
}

type TwoD[T any] struct {
	data   [][]T
	maxX   int
	maxY   int
	x      int
	y      int
	inited bool
}

func (it *TwoD[T]) Next() bool {
	if !it.inited {
		it.inited = true
		return true
	}

	if it.y == it.maxY {
		if it.x == it.maxX {
			return false
		}

		it.y = 0
		it.x++
		return true
	}

	it.y++

	return true
}

func (it *TwoD[T]) Get() T {
	return it.data[it.x][it.y]
}

func (it *TwoD[T]) X() int {
	return it.x
}

func (it *TwoD[T]) Y() int {
	return it.y
}
