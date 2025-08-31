package utils

func DiffInt(a, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}
