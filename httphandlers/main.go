package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/string", String("I'm a frayed knot."))

	someStructure := Struct{Greeting: "Hello", Punct: ",", Who: "Gophers!"}
	http.Handle("/struct", someStructure)
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (s Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s%s %s", s.Greeting, s.Punct, s.Who)
}
