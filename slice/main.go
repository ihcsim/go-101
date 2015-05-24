package main

import "fmt"

func main() {

	// append two slices
	slice1 := []int{0, 1, 2, 3, 4}
	slice2 := []int{5, 6, 7, 8}
	fmt.Printf("%T\n", slice2)
	fmt.Println(append(slice1, slice2...))
}
