package main

import "testing"

var stats *statistics

func setUp() {
	precision := 4
	stats = NewStatistics(precision)
}

func TestCompute_GivenEmptyInputsArray_ReturnAnError(t *testing.T) {
	setUp()

	inputs := []float64{}
	stats.Compute(inputs)
	if stats.err == nil {
		t.Errorf("Expected error didn't occur")
	}
}

func TestCompute_GivenFloatIntegers_CorrectlyComputeAllStats(t *testing.T) {
	setUp()

	var tests = []struct {
		inputs                    []float64
		expectedMean              float64
		expectedMedian            float64
		expectedSum               float64
		expectedStandardDeviation float64
		expectedModes             []float64
	}{
		{inputs: []float64{1.00, 2.00, 3.00, 3.00, 4.00, 5.00},
			expectedMean:              3.00,
			expectedMedian:            3.00,
			expectedStandardDeviation: 1.4142,
			expectedSum:               18.00,
			expectedModes:             []float64{3.00}},
	}

	for _, test := range tests {
		stats.Compute(test.inputs)

		if stats.err != nil {
			t.Errorf("Expected no errors to occur")
		}

		if stats.mean != test.expectedMean {
			t.Errorf("Expected mean to be %f, but got %f", test.expectedMean, stats.mean)
		}

		if stats.median != test.expectedMedian {
			t.Errorf("Expected median to be %f, but got %f", test.expectedMedian, stats.median)
		}

		if stats.sum != test.expectedSum {
			t.Errorf("Expected sum to be %f, but got %f", test.expectedSum, stats.sum)
		}

		if stats.standardDeviation != test.expectedStandardDeviation {
			t.Errorf("Expected stnadard deviation to be %f, but got %f", test.expectedStandardDeviation, stats.standardDeviation)
		}
	}
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
		if actual, _ := stats.computeSum(); test.expected != actual {
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
		if actual, _ := stats.computeMean(); test.expected != actual {
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
		if median, _ := stats.computeMedian(); test.expected != median {
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
		if median, _ := stats.computeMedian(); test.expected != median {
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
		if standardDeviation, _ := stats.computeStandardDeviation(); test.expected != standardDeviation {
			t.Errorf("Expected standard deviation to be %f, but got %f", test.expected, standardDeviation)
		}
	}
}

func TestComputeMean_GivenEmptyInputsArray_ReturnAnError(t *testing.T) {
	setUp()

	stats.numbers = []float64{}
	_, err := stats.computeMean()
	if err == nil {
		t.Errorf("Expected error to have occurred")
	}
}

func TestComputeSum_GivenEmptyInputsArray_ReturnAnError(t *testing.T) {
	setUp()

	stats.numbers = []float64{}
	_, err := stats.computeSum()
	if err == nil {
		t.Errorf("Expected error to have occurred")
	}
}

func TestComputeMedian_GivenEmptyInputsArray_ReturnAnError(t *testing.T) {
	setUp()

	stats.numbers = []float64{}
	_, err := stats.computeMedian()
	if err == nil {
		t.Errorf("Expected error to have occurred")
	}
}

func TestComputeStandardDeviation_GivenEmptyInputsArray_ReturnAnError(t *testing.T) {
	setUp()

	stats.numbers = []float64{}
	_, err := stats.computeStandardDeviation()
	if err == nil {
		t.Errorf("Expected error to have occurred")
	}
}

func TestComputeMode_GivenInputsWithOneMode_ReturnCorrectMode(t *testing.T) {
	setUp()

	var tests = []struct {
		inputs   []float64
		expected float64
	}{
		{inputs: []float64{1.0, 2.0, 3.0, 4.0, 5.0, 5.0, 6.0, 6.0, 6.0, 7.0, 7.0, 8.0, 9.0, 10.0, 10.0},
			expected: 6.0},
		{inputs: []float64{-1.0, -2.0, -2.0, -2.0, -2.0, -3.0, -4.0, -5.0, -6.0, -7.0, -8.0, -8.0, -9.0, -10.0},
			expected: -2.0},
		{inputs: []float64{10.00, 20.00, 30.00, 30.00, 40.00, 40.00, 40.00, 40.00, 50.00, 60.00, 70.00, 70.00},
			expected: 40.00},
		{inputs: []float64{100.00, 100.00, 100.00, 200.00, 300.00, 400.00, 400.00, 500.00, 500.00},
			expected: 100.0},
	}

	for _, test := range tests {
		stats.numbers = test.inputs
		// expecting only one mode
		if actuals, _ := stats.computeModes(); actuals[0] != test.expected {
			t.Errorf("Expected mode to be %f, but got %f", test.expected, actuals[0])
		}
	}
}

func TestComputeMode_GivenInputsWithMultipleModes_ReturnAllModes(t *testing.T) {
	setUp()

	var tests = []struct {
		inputs   []float64
		expected []float64
	}{
		{inputs: []float64{1.0, 2.0, 3.0, 4.0, 5.0, 5.0, 6.0, 6.0, 6.0, 7.0, 7.0, 7.0, 8.0, 9.0, 10.0, 10.0},
			expected: []float64{6.0, 7.0}},
		{inputs: []float64{-1.0, -2.0, -3.0, -4.0, -5.0, -5.0, -6.0, -6.0, -6.0, -7.0, -7.0, -7.0, -8.0, -9.0, -10.0, -10.0},
			expected: []float64{-7.0, -6.0}},
		{inputs: []float64{10.0, 20.0, 30.0, 40.0, 50.0, 50.0, 60.0, 60.0, 70.0, 70.0, 70.0, 70.0, 80.0, 90.0, 100.0, 100.0},
			expected: []float64{70.0}},
		{inputs: []float64{100.0, 100.0, 200.0, 200.0, 300.0, 300.0, 400.0, 500.0, 500.0},
			expected: []float64{100.0, 200.0, 300.0, 500.0}},
	}

	for _, test := range tests {
		stats.numbers = test.inputs
		actuals, _ := stats.computeModes()
		for _, expected := range test.expected {
			ok := false
			for _, actual := range actuals {
				if actual == expected {
					ok = true
				}
			}

			if !ok {
				t.Errorf("Expected modes to include %f, but not found. Actual modes: %v", expected, actuals)
			}
		}
	}
}
