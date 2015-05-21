package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

type Person struct {
	name   string
	age    int
	income float64
}

type Address struct {
	street_number int
	street_name   string
	city          string
	province      string
	postal_code   string
}

func main() {
	fmt.Printf("1 + 2 = %d\n", add(1, 2))

	for i := 0; i <= 10; i++ {
		fmt.Printf("10 - %d = %d\n", i, subtract(10, i))
	}

	fmt.Printf("The square root of 64 is %f\n", squareRoot(64))

	fmt.Printf("Using the Newton's method, the square root of 2 is %f\n", newtonMethod(2.0))

	fmt.Printf("Go runs on %s\n", checkOS())

	fmt.Printf("Saturday is %s\n", calculateDaysAway())

	me := Person{"Ivan Sim", 35, 88000.00}
	fmt.Println(me)
	fmt.Printf("Name: %s, Age: %d, Income: %f\n", me.name, me.age, me.income)

	myAddress := Address{
		street_number: 511,
		street_name:   "Rochester Avenue",
		city:          "Coquitlam",
		province:      "BC",
		postal_code:   "V3K0A2",
	}
	fmt.Println(myAddress)

	var addressPointer *Address = &myAddress
	fmt.Println(*addressPointer)

	fav_number := 5500
	fav_number_pointer := &fav_number
	fmt.Printf("Pointer address: %d, Pointer value: %d\n", fav_number_pointer, *fav_number_pointer)

}

func add(operand1, operand2 int) (sum int) {
	return operand1 + operand2
}

func subtract(operand1, operand2 int) (subtraction int) {
	return operand1 - operand2
}

func squareRoot(operand float64) (result float64) {
	return math.Sqrt(operand)
}

func newtonMethod(operand float64) (result float64) {
	z := 1.0
	for i := 0; i < 20; i++ {
		z = z - ((z*z - operand) / 2 * z)
		fmt.Println(z)
	}
	return z
}

func checkOS() (os string) {
	switch os := runtime.GOOS; os {
	case "darwin":
		return "OS X"
	case "linux":
		return "Linux"
	default:
		return os
	}
}

func calculateDaysAway() (daysAway string) {
	today := time.Now().Weekday()
	switch saturday := time.Saturday; saturday {
	case today + 0:
		return "today!"
	case today + 1:
		return "1 day away!"
	case today + 2:
		return "2 days away!"
	default:
		return "too far away!"
	}
}
