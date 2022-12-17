package graph

type node[T comparable] struct {
	V    T
	Dist int
}

type Iterator[T comparable] interface {
	Next(T) []T
	Constraint(from, to T) bool
}

func ShortestPath[T comparable](from, to T, it Iterator[T]) int {
	visited := make(map[T]bool)
	visited[from] = true

	var queue []node[T]
	queue = append(queue, node[T]{V: from})

	for len(queue) > 0 {
		current := queue[0]

		if current.V == to {
			return current.Dist
		}

		queue = queue[1:]

		for _, next := range it.Next(current.V) {
			if !visited[next] && !it.Constraint(current.V, next) {
				visited[next] = true
				queue = append(queue, node[T]{V: next, Dist: current.Dist + 1})
			}
		}
	}

	return -1
}
