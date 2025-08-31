package utils

func MapSlice[T, U any](s []T, f func(T) U) (result []U) {
	for _, v := range s {
		result = append(result, f(v))
	}
	return
}
