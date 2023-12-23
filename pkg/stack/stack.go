package stack

type Stack[T any] struct {
	elems []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{nil}
}

func (s *Stack[T]) Push(v T) {
	s.elems = append(s.elems, v)
}

func (s *Stack[T]) Prepend(v T) {
	s.elems = append([]T{v}, s.elems...)
}

func (s *Stack[T]) Pop() T {
	n := len(s.elems) - 1
	b := s.elems[n]
	s.elems = s.elems[:n]
	return b
}

func (s *Stack[T]) Top() T {
	return s.elems[len(s.elems)-1]
}

func (s *Stack[T]) PopMultiple(n int) []T {
	size := len(s.elems)
	m := s.elems[size-n : size]
	s.elems = s.elems[:size-n]
	return m
}

func (s *Stack[T]) PushMultiple(m []T) {
	s.elems = append(s.elems, m...)
}

func (s *Stack[T]) Size() int {
	return len(s.elems)
}
