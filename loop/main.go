package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println("Sqrt of 1:", Sqrt(1))
	fmt.Println("Sqrt of 2:", Sqrt(2))
	fmt.Println("Sqrt of 4:", Sqrt(4))
	fmt.Println("Sqrt of 16:", Sqrt(16))
	fmt.Println("Sqrt of 25:", Sqrt(25))
}

func Sqrt(input float64) float64 {
	const TOLERANCE, START = 0.000001, 1.0
	var result, temp = START, float64(0.0)
	for {
		temp = result - ((result*result - input) / (2 * result))
		if diff := math.Abs(result - temp); diff < TOLERANCE {
			break
		} else {
			result = temp
		}
	}

	truncatedStr := fmt.Sprintf("%.6f", result)
	result, err := strconv.ParseFloat(truncatedStr, 6)
	if err != nil {
	}
	return result
}
