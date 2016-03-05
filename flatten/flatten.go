package flatten

// Flatten converts a 2-D slice into a 1-D slice where matrix[i][j] is inserted into output[i+j]
func Flatten(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	output := matrix[0]
	for i := 1; i < len(matrix); i++ {
		output = append(output, matrix[i]...)
	}

	return output
}
