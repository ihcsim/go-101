package make2d

// Make2D converts slice into a 2D matrix with inner slices of length depth.
func Make2D(slice []int, depth int) [][]int {
	if len(slice) == 0 || depth == 0 {
		return [][]int{}
	}

	var numRows int
	if len(slice)%depth != 0 {
		numRows = len(slice)/depth + 1
	} else {
		numRows = len(slice) / depth
	}
	matrix := make([][]int, numRows)

	for i := 0; i < len(matrix); i++ {
		start := i * depth

		// if this row's depth exceeded total number of entries,
		// add padding.
		var end, padding int = start + depth, 0
		if end > len(slice) {
			padding = end - len(slice)
			end = len(slice)
		}

		matrix[i] = slice[start:end]
		for j := 0; j < padding; j++ {
			matrix[i] = append(matrix[i], 0)
		}
	}

	return matrix
}
