package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	go Ping(c1)

	c2 := make(chan string)
	go Pong(c2)

	for {
		select {
		case msg := <-c1:
			fmt.Println(msg)
		case msg := <-c2:
			fmt.Println(msg)
		default:
			fmt.Println("No data yet....")
			time.Sleep(time.Second * 1)
		}
	}

	var input string
	fmt.Scanln(&input)

}

func Ping(c chan string) {
	for {
		c <- "Ping!"
		time.Sleep(time.Second * 2)
	}
}

func Pong(c chan string) {
	for {
		c <- "Pong!"
		time.Sleep(time.Second * 4)
	}
}
