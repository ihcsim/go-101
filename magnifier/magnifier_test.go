package magnifier

import (
	"testing"
)

type testpairs struct {
	input    byte
	expected string
}

var tests = []testpairs{
	{input: 'A', expected: "A MAGNIFIED A"},
	{input: 'B', expected: "A MAGNIFIED B"},
	{input: 'C', expected: "A MAGNIFIED C"},
	{input: 'D', expected: "A MAGNIFIED D"},
	{input: 'E', expected: "A MAGNIFIED E"},
}

func TestMagnify(t *testing.T) {
	for _, test := range tests {
		actualOutput := Magnify(test.input)
		if actualOutput != test.expected {
			t.Errorf("For input: %q, Expected to get %s, but Actual is %s\n", test.input, test.expected, actualOutput)
		}
	}
}
