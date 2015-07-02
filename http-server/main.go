package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	counter := new(Counter)
	err := http.ListenAndServe("localhost:7000", counter)
	if err != nil {
		log.Fatal("Server Failed!")
	}
}

type Counter struct {
	counts       int
	notification chan *http.Request
}

func (counter *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	counter.counts++
	fmt.Fprintf(w, "counter = %d\n", counter.counts)
}
