package unique

import "testing"

func TestUniqueInt(t *testing.T) {
	var tests = []struct {
		input    []int
		expected []int
	}{
		{input: []int{1, 2, 3, 4, 5}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{2, 3, 4, 5, 6}, expected: []int{2, 3, 4, 5, 6}},
		{input: []int{1, 1, 2, 3, 4, 5}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{1, 1, 2, 2, 2}, expected: []int{1, 2}},
		{input: []int{1, 1, 1, 1, 1}, expected: []int{1}},
		{input: []int{1, 1, 2, 2, 3, 3}, expected: []int{1, 2, 3}},
		{input: []int{0}, expected: []int{0}},
	}

	for _, test := range tests {
		actual := UniqueInt(test.input)
		if len(test.expected) != len(actual) {
			t.Errorf("Bad input %v.\nExpected length of result slice to be %d, but got %d", test.input, len(test.expected), len(actual))
		}

		for i := 0; i < len(test.expected); i++ {
			if test.expected[i] != actual[i] {
				t.Errorf("Bad input %v.\nExpected value at %d to be %d, but got %d", test.input, i, test.expected[i], actual[i])
			}
		}
	}
}
