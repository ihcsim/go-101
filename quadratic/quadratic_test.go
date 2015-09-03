package quadratic

import "testing"

// Verified using calculator at http://www.math.com/students/calculators/source/quadratic.htm
func TestSolve_GivenRealCoefficients_ReturnsTheCorrectSolution(t *testing.T) {
	precision := 4
	var tests = []struct {
		input    *equation
		expected *solution
	}{
		{input: NewEquation(1.00, -4.00, 4.00, precision),
			expected: NewSolution(2.00, 2.00)},
		{input: NewEquation(1.00, 10.00, 25.00, precision),
			expected: NewSolution(-5.00, -5.00)},
		{input: NewEquation(1.00, 0.00, -9.00, precision),
			expected: NewSolution(3.00, -3.00)},
		{input: NewEquation(1.00, 0.00, -81.00, precision),
			expected: NewSolution(9.00, -9.00)},
		{input: NewEquation(2.00, 8.00, 8.00, precision),
			expected: NewSolution(-2.00, -2.00)},
		{input: NewEquation(2.00, 16.00, 32.00, precision),
			expected: NewSolution(-4.00, -4.00)},
		{input: NewEquation(2.00, 32.00, 64.00, precision),
			expected: NewSolution(-2.3431, -13.6568)},
		{input: NewEquation(-3.00, 10.00, 20.00, precision),
			expected: NewSolution(-1.4065, 4.7398)},
	}

	for _, test := range tests {
		actual, err := Solve(test.input)
		if err != nil {
			t.Errorf("Unexpected error occurred", err)
		}
		if actual.x1 != test.expected.x1 {
			t.Errorf("Expected x1 to be %f, but got %f", test.expected.x1, actual.x1)
		}

		if actual.x2 != test.expected.x2 {
			t.Errorf("Expected x2 to be %f, but got %f", test.expected.x2, actual.x2)
		}
	}
}

func TestSolve_GivenZeroQuadraticCoefficent_ReturnsAnError(t *testing.T) {
	_, err := Solve(NewEquation(0.00, 2.00, 1.00, 4))
	if err == nil {
		t.Errorf("Expected error didn't occur.")
	}
}

// Verify using calculator at http://www.mathwarehouse.com/quadratic/quadratic-formula-calculator.php
func TestSolve_WhenDiscriminatIsNegative_ReturnsComplexNumberSolution(t *testing.T) {
	precision := 4
	var tests = []struct {
		input    *equation
		expected *solution
	}{
		{input: NewEquation(1.00, 2.00, 4.00, precision),
			expected: NewSolution(complex(-1.0000, 1.7320), complex(-1.0000, -1.7320)),
		},
		{input: NewEquation(2.00, 4.00, 16.00, precision),
			expected: NewSolution(complex(-1.0000, 2.6457), complex(-1.0000, -2.6457)),
		},
		{input: NewEquation(-8.00, 5.00, -4.00, precision),
			expected: NewSolution(complex(0.3125, -0.6343), complex(0.3125, 0.6343)),
		},
		{input: NewEquation(4.00, 10.00, 10.00, precision),
			expected: NewSolution(complex(-1.2500, 0.9682), complex(-1.2500, -0.9682)),
		},
		{input: NewEquation(100.00, 81.00, 23.00, precision),
			expected: NewSolution(complex(-0.4050, 0.2568), complex(-0.4050, -0.2568)),
		},
		{input: NewEquation(50.00, -24.00, 13.00, precision),
			expected: NewSolution(complex(0.2400, 0.4498), complex(0.2400, -0.4498)),
		},
		{input: NewEquation(12.00, -14.00, 20.00, precision),
			expected: NewSolution(complex(0.5833, 1.1516), complex(0.5833, -1.1516)),
		},
	}

	for _, test := range tests {
		actual, err := Solve(test.input)
		if err != nil {
			t.Errorf("Unexpected error ocuurred. ", err)
		}

		if actual.x1 != test.expected.x1 {
			t.Errorf("Expected x1 to be %.4f, but got %.4f", test.expected.x1, actual.x1)
		}

		if actual.x2 != test.expected.x2 {
			t.Errorf("Expected x2 to be %.4f, but got %.4f", test.expected.x2, actual.x2)
		}
	}
}
