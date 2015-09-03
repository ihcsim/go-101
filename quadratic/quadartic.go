package quadratic

import (
	"errors"
	"math"
	"math/cmplx"
)

// coefficients has fields that represent the quadratic coefficient, linear coefficient
// and constant of a quadratic equation.
type coefficients struct {
	quadratic complex128
	linear    complex128
	constant  complex128
}

// equation represents a quadratic equation.
type equation struct {
	coefficients
	precision int
}

func NewEquation(quadratic, linear, constant complex128, precision int) *equation {
	return &equation{
		coefficients{quadratic: quadratic, linear: linear, constant: constant},
		precision}
}

// solution has fields that represents the answer to a quadratic equation
type solution struct {
	x1 complex128
	x2 complex128
}

func NewSolution(x1, x2 complex128) *solution {
	return &solution{
		x1: x1,
		x2: x2,
	}
}

// Solve attempts to solve the quadratic equation made up of the given coefficients using
// the quadratic formula.
func Solve(e *equation) (s *solution, err error) {
	if e.quadratic == 0 {
		return nil, errors.New("The quadratic coefficient cannot be 0.")
	}

	x1 := (-e.linear + discriminant(e.coefficients)) / divisor(e.coefficients)
	x2 := (-e.linear - discriminant(e.coefficients)) / divisor(e.coefficients)

	s = &solution{
		x1: complex(roundToPrecision(real(x1), e.precision), roundToPrecision(imag(x1), e.precision)),
		x2: complex(roundToPrecision(real(x2), e.precision), roundToPrecision(imag(x2), e.precision)),
	}
	return
}

func discriminant(c coefficients) complex128 {
	return cmplx.Sqrt(c.linear*c.linear - 4*c.quadratic*c.constant)
}

func divisor(c coefficients) complex128 {
	return 2 * c.quadratic
}

func roundToPrecision(f float64, precision int) float64 {
	multiplier := math.Pow(10, float64(precision))
	return float64(int(f*multiplier)) / multiplier
}
