package priorityqueue

type PQ[T any] struct {
	elems []any
	less  less[T]
}

type less[T any] func(T, T) bool

func New[T any](fn less[T]) *PQ[T] {
	return &PQ[T]{
		less: fn,
	}
}

func (q *PQ[T]) Push(v any) {
	q.elems = append(q.elems, v)
}

func (q *PQ[T]) Pop() any {
	if len(q.elems) == 0 {
		return *new(T)
	}
	n := len(q.elems) - 1
	b := q.elems[n]
	q.elems = q.elems[:n]
	return b
}

func (q *PQ[T]) Len() int {
	return len(q.elems)
}

func (q *PQ[T]) Less(i, j int) bool {
	return q.less(q.elems[i].(T), q.elems[j].(T))
}

func (q *PQ[T]) Swap(i, j int) {
	q.elems[i], q.elems[j] = q.elems[j], q.elems[i]
}
