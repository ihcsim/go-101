package main

import "testing"

func TestSoundex(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{input: "Robert", expected: "R163"},
		{input: "Rupert", expected: "R163"},
		{input: "Ashcraft", expected: "A261"},
		{input: "Ashcroft", expected: "A261"},
		{input: "Mary", expected: "M600"},
		{input: "Simon", expected: "S550"},
		{input: "Allen", expected: "A450"},
		{input: "Sabestian", expected: "S123"},
		{input: "Laura", expected: "L600"},
		{input: "Tymczak", expected: "T522"},
		{input: "Elizabeth", expected: "E421"},
		{input: "Ivan", expected: "I150"},
		{input: "Jennifer", expected: "J516"},
		{input: "Ash", expected: "A200"},
		{input: "Loretta", expected: "L630"},
	}

	for _, test := range tests {
		if actual := Soundex(test.input); actual != test.expected {
			t.Errorf("Expected encoding of %q to be %q, but got %q", test.input, test.expected, actual)
		}
	}
}
