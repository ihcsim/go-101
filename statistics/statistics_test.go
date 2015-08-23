package main

import "testing"

var stats *statistics

func setUp() {
	stats = NewStatistics()
}

func TestComputeSum_GivenFloatIntegers_ReturnSum(t *testing.T) {
	setUp()

	var tests = []struct {
		inputs   []float64
		expected float64
	}{
		{inputs: []float64{1.00, 2.00, 3.00, 4.00, 5.00},
			expected: 1.00 + 2.00 + 3.00 + 4.00 + 5.00},
		{inputs: []float64{-1.00, -2.00, -3.00, -4.00, -5.00},
			expected: -1.00 + -2.00 + -3.00 + -4.00 + -5.00},
		{inputs: []float64{11.00, 12.00, 13.00, 14.00, 15.00},
			expected: 11.00 + 12.00 + 13.00 + 14.00 + 15.00},
		{inputs: []float64{16.00, 17.00, 18.00, 19.00, 20.00},
			expected: 16.00 + 17.00 + 18.00 + 19.00 + 20.00},
		{inputs: []float64{21.00, 22.00, 23.00, 24.00, 25.00},
			expected: 21.00 + 22.00 + 23.00 + 24.00 + 25.00},
	}
	for _, test := range tests {
		stats.numbers = test.inputs
		if actual := stats.computeSum(); test.expected != actual {
			t.Errorf("Expected sum to be %f, but get %f", test.expected, actual)
		}
	}
}

func TestComputeMean_GivenFloatIntegers_ReturnSum(t *testing.T) {
	setUp()

	var tests = []struct {
		inputs   []float64
		expected float64
	}{
		{inputs: []float64{1.00, 2.00, 3.00, 4.00, 5.00},
			expected: 3.00},
		{inputs: []float64{-1.00, -2.00, -3.00, -4.00, -5.00},
			expected: -3.00},
		{inputs: []float64{11.00, 12.00, 13.00, 14.00, 15.00},
			expected: 13.00},
		{inputs: []float64{16.00, 17.00, 18.00, 19.00, 20.00},
			expected: 18.00},
		{inputs: []float64{21.00, 22.00, 23.00, 24.00, 25.00},
			expected: 23.00},
	}
	for _, test := range tests {
		stats.numbers = test.inputs
		if actual := stats.computeMean(); test.expected != actual {
			t.Errorf("Expected mean to be %f, but get %f", test.expected, actual)
		}
	}
}

func TestComputeMedian_GivenOddCounts_ReturnMiddleValue(t *testing.T) {
	setUp()

	var tests = []struct {
		inputs   []float64
		expected float64
	}{
		{inputs: []float64{1.00, 2.00, 3.00, 4.00, 5.00},
			expected: 3.00},
		{inputs: []float64{-1.00, -2.00, -3.00, -4.00, -5.00},
			expected: -3.00},
		{inputs: []float64{11.00, 12.00, 13.00, 14.00, 15.00},
			expected: 13.00},
		{inputs: []float64{16.00, 17.00, 18.00, 19.00, 20.00},
			expected: 18.00},
		{inputs: []float64{21.00, 22.00, 23.00, 24.00, 25.00},
			expected: 23.00},
	}

	for _, test := range tests {
		stats.numbers = test.inputs
		if median := stats.computeMedian(); test.expected != median {
			t.Errorf("Expected median to be %f, but got %f", test.expected, median)
		}
	}
}

func TestComputeMedian_GivenEvenCounts_ReturnAverageOfMiddleValues(t *testing.T) {
	setUp()

	var tests = []struct {
		inputs   []float64
		expected float64
	}{
		{inputs: []float64{1.00, 2.00, 3.00, 4.00, 5.00, 6.00},
			expected: 3.50},
		{inputs: []float64{-1.00, -2.00, -3.00, -4.00, -5.00, -6.00},
			expected: -3.50},
		{inputs: []float64{11.00, 12.00, 13.00, 14.00, 15.00, -16.00},
			expected: 13.50},
		{inputs: []float64{16.00, 17.00, 18.00, 19.00, 20.00, -21.00},
			expected: 18.50},
		{inputs: []float64{21.00, 22.00, 23.00, 24.00, 25.00, 25.00},
			expected: 23.50},
	}

	for _, test := range tests {
		stats.numbers = test.inputs
		if median := stats.computeMedian(); test.expected != median {
			t.Errorf("Expected median to be %f, but got %f", test.expected, median)
		}
	}
}
