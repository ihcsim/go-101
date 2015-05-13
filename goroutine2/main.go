package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go ping(c)
	go pong(c)
	go print(c)

	var input string
	fmt.Scanln(&input)
}

func ping(c chan<- string) {
	for i := 0; i < 10; i++ {
		c <- "Ping!"
	}
}

func pong(c chan<- string) {
	for i := 0; i < 10; i++ {
		c <- "Pong!"
	}
}

func print(c <-chan string) {
	for {
		fmt.Println(<-c)
		time.Sleep(time.Second * 1)
	}
}
