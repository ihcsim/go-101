package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	counter := newCounter()
	go func() {
		<-counter.notification
		fmt.Println("Received a request")
	}()

	fmt.Println("Starting server at localhost:7000")
	err := http.ListenAndServe("localhost:7000", counter)
	if err != nil {
		log.Fatal("Server Failed!")
	}
}

type Counter struct {
	counts       int
	notification chan int
}

func newCounter() *Counter {
	return &Counter{
		counts:       0,
		notification: make(chan int),
	}
}

func (counter *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	counter.notification <- 1
	counter.counts++
	fmt.Fprintf(w, "counter = %d\n", counter.counts)
}
