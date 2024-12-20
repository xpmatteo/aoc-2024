package matrix

func New[T any](rows, cols int) [][]T {
	var result = make([][]T, rows)
	for r := range rows {
		result[r] = make([]T, cols)
	}
	return result
}
