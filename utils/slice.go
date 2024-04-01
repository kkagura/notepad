package utils

func Unshift[T any](s []T, v T) []T {
	s = append([]T{v}, s...)
	return s
}
