package main

import "fmt"

func main() {
	red := make(chan int)
	green := make(chan int)
	refactor := make(chan int)

	go func() {
		for {
			red <- 1
			green <- 1
			refactor <- 1
		}
	}()

	RedGreenRefactor(red, green, refactor)
}

func RedGreenRefactor(red, green, refactor chan int) {
	for {
		select {
		case <-red:
			fmt.Println("Write a failing test...")
		case <-green:
			fmt.Println("Get the test to pass...")
		case <-refactor:
			fmt.Println("Refactor the codes...")
		}
	}
}
