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
