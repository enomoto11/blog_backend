package utils

func MapSlice[T, O any](i []T, f func(T) O) []O {
	n := make([]O, len(i))
	for i, e := range i {
		n[i] = f(e)
	}
	return n
}
