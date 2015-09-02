package quadratic

import "testing"

// Verified using calculator at http://www.math.com/students/calculators/source/quadratic.htm
func TestSolveEquation_ReturnsTheCorrectSolution(t *testing.T) {
	precision := 4
	var tests = []struct {
		input    *equation
		expected *solution
	}{
		{input: &equation{
			coefficients{quadratic: 1.00, linear: -4.00, constant: 4.00},
			precision},
			expected: &solution{x1: 2.00, x2: 2.00}},
		{input: &equation{
			coefficients{quadratic: 1.00, linear: 10.00, constant: 25.00},
			precision},
			expected: &solution{x1: -5.00, x2: -5.00}},
		{input: &equation{
			coefficients{quadratic: 1.00, linear: 0.00, constant: -9.00},
			precision},
			expected: &solution{x1: 3.00, x2: -3.00}},
		{input: &equation{
			coefficients{quadratic: 1 + 0i, linear: 0.00, constant: -81.00},
			precision},
			expected: &solution{x1: 9.00, x2: -9.00}},
		{input: &equation{
			coefficients{quadratic: 2 + 0i, linear: 8.00, constant: 8.00},
			precision},
			expected: &solution{x1: -2.00, x2: -2.00}},
		{input: &equation{
			coefficients{quadratic: 2.00, linear: 16.00, constant: 32.00},
			precision},
			expected: &solution{x1: -4.00, x2: -4.00}},
		{input: &equation{
			coefficients{quadratic: 2 + 0i, linear: 32.00, constant: 64.00},
			precision},
			expected: &solution{x1: -2.3431, x2: -13.6568}},
	}

	for _, test := range tests {
		actual, err := SolveEquation(test.input)
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

func TestSolveQuadraticEquation_WhenQuadraticCoefficentIsZero_ReturnsAnError(t *testing.T) {}
func TestSolveQuadraticEquation_WhenDiscriminatIsNegative(t *testing.T)                    {}
