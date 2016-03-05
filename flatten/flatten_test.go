package flatten

import "testing"

func TestFlatten(t *testing.T) {
	var tests = []struct {
		input    [][]int
		expected []int
	}{
		{input: [][]int{{}}, expected: []int{}},
		{input: [][]int{{1, 2, 3, 4, 5}},
			expected: []int{1, 2, 3, 4, 5},
		},
		{input: [][]int{
			{1, 2, 3, 4, 5, 6},
			{7, 8, 9, 10, 11, 12},
			{13, 14, 15, 16, 17, 18},
			{19, 20}},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		},
		{input: [][]int{
			{10},
			{9, 8},
			{7, 6, 5},
			{4, 3, 2, 1}},
			expected: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
	}

	for _, test := range tests {
		actual := Flatten(test.input)

		if len(test.expected) != len(actual) {
			t.Errorf("Slice length mismatch. Expected %d, but got %d", len(test.expected), len(actual))
		}

		for index, value := range actual {
			if test.expected[index] != value {
				t.Errorf("Slice entries mismatched. Expected %d, but got %d", test.expected[index], value)
			}
		}
	}
}
