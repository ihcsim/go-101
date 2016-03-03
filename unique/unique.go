package unique

// UniqueInt eliminates all duplicates in input
func UniqueInt(input []int) []int {
	var result []int

	seen := make(map[int]struct{})
	for _, value := range input {
		if _, ok := seen[value]; !ok {
			seen[value] = struct{}{}
			result = append(result, value)
		}
	}

	return result
}
