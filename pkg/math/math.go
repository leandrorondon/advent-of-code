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
	var max T
	if len(v) < 1 {
		return max
	}

	max = v[0]
	for i := 1; i < len(v); i++ {
		max = Max(max, v[i])
	}

	return max
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
