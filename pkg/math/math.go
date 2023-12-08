package math

import "math"

func Abs[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float64 | float32](v T) T {
	return T(math.Abs(float64(v)))
}

func Max[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float64 | float32](v1, v2 T) T {
	if v2 > v1 {
		return v2
	}

	return v1
}

func MaxS[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float64 | float32](v ...T) T {
	var m T
	if len(v) < 1 {
		return m
	}

	m = v[0]
	for i := 1; i < len(v); i++ {
		m = Max(m, v[i])
	}

	return m
}

func Min[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float64 | float32](v1, v2 T) T {
	if v2 < v1 {
		return v2
	}

	return v1
}

func MinMax[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float64 | float32](v1, v2 T) (T, T) {
	if v2 < v1 {
		return v2, v1
	}

	return v1, v2
}

func Sum[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float64 | float32](arr []T) T {
	sum := T(0)
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	return sum
}

func BetweenInclusive[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float64 | float32](v, b, e T) bool {
	return v >= b && v <= e
}
