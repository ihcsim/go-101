package make2d

import "testing"

func TestMake2D(t *testing.T) {
	var tests = []struct {
		input    []int
		depth    int
		expected [][]int
	}{
		{input: []int{},
			depth:    0,
			expected: [][]int{},
		},
		{input: []int{1, 2},
			depth: 1,
			expected: [][]int{
				{1},
				{2}},
		},
		{input: []int{1, 2, 3, 4, 5},
			depth: 5,
			expected: [][]int{
				{1, 2, 3, 4, 5},
			},
		},
		{input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			depth: 3,
			expected: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
				{10, 0, 0},
			},
		},
		{input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			depth: 6,
			expected: [][]int{
				{1, 2, 3, 4, 5, 6},
				{7, 8, 9, 10, 0, 0},
			},
		},
		{input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			depth: 4,
			expected: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
				{17, 18, 19, 20}},
		},
	}

	for _, test := range tests {
		actual := Make2D(test.input, test.depth)
		if len(actual) != len(test.expected) {
			t.Errorf("Bad slice %v.\nExpected slice length to be %d, but got %d", test.input, len(test.expected), len(actual))
		}

		for i := 0; i < len(test.expected); i++ {
			for j := 0; j < len(test.expected[i]); j++ {
				if test.expected[i][j] != actual[i][j] {
					t.Errorf("Bad slice %v.\nExpected entry at (%d, %d) to be %d, but got %d", test.input, i, j, test.expected[i][j], actual[i][j])
				}
			}
		}
	}
}
