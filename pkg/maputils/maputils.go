package maputils

import (
	"cmp"
	"slices"
)

func Copy[K comparable, V any](in map[K]V) map[K]V {
	out := make(map[K]V)

	for k, v := range in {
		out[k] = v
	}

	return out
}

func KeysToSlice[K cmp.Ordered, V any](in map[K]V) []K {
	out := make([]K, len(in))

	i := 0
	for k := range in {
		out[i] = k
		i++
	}
	slices.Sort(out)

	return out
}
