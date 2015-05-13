package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("I am first")
	}()

	go func() {
		fmt.Println("I am second")
	}()

	Sleep(3)
}

func Sleep(durationInMilliseconds time.Duration) {
	select {
	case <-time.After(time.Millisecond * durationInMilliseconds):
		fmt.Println("Sleep's over...")
	}
}
