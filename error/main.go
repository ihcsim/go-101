package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %6.2f", float64(e))
}
