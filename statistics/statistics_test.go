package main

import "testing"

var stats *statistics

func setUp() {
	precision := 4
	stats = NewStatistics(precision)
}

func TestComputeSum_GivenFloatIntegers_ReturnSum(t *testing.T) {
	setUp()

	var tests = []struct {
		inputs   []float64
		expected float64
	}{
		{inputs: []float64{1.00, 2.00, 3.00, 4.00, 5.00},
			expected: 15.00},
		{inputs: []float64{-1.00, -2.00, -3.00, -4.00, -5.00},
			expected: -15.00},
		{inputs: []float64{11.10, 12.23, 13.56, 14.88, 15.02},
			expected: 66.79},
		{inputs: []float64{16.30, 17.12, 18.95, 19.44, 20.77},
			expected: 92.58},
		{inputs: []float64{21.44, 22.22, 23.56, 24.43, 25.11},
			expected: 116.76},
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
		{inputs: []float64{11.10, 12.23, 13.56, 14.88, 15.02},
			expected: 13.3580},
		{inputs: []float64{16.30, 17.12, 18.95, 19.44, 20.77},
			expected: 18.5159},
		{inputs: []float64{21.44, 22.22, 23.56, 24.43, 25.11},
			expected: 23.3520},
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

func TestComputeStandardDeviation_ReturnStandardDeviation(t *testing.T) {
	setUp()

	var tests = []struct {
		inputs   []float64
		expected float64
	}{
		{inputs: []float64{1.00, 2.00, 3.00, 4.00, 5.00, 6.00},
			expected: 1.8708},
		{inputs: []float64{-1.00, -2.00, -3.00, -4.00, -5.00, -6.00},
			expected: 1.8708},
		{inputs: []float64{11.00, 12.00, 13.00, 14.00, 15.00, 16.00, 17.00},
			expected: 2.1602},
		{inputs: []float64{16.00, 17.00, 18.00, 19.00, 20.00, 21.00, 22.00, 23.00},
			expected: 2.4494},
		{inputs: []float64{21.00, 22.00, 23.00, 24.00, 25.00},
			expected: 1.5811},
	}

	for _, test := range tests {
		stats.numbers = test.inputs
		if standardDeviation := stats.computeStandardDeviation(); test.expected != standardDeviation {
			t.Errorf("Expected standard deviation to be %f, but got %f", test.expected, standardDeviation)
		}
	}
}

func TestComputeMean_GivenEmptyInputsArray(t *testing.T)              {}
func TestComputeStandardDeviation_GivenEmptyInputsArray(t *testing.T) {}
