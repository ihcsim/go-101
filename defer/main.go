package main

import "fmt"

func main() {
	defer untrace(trace("foo"))
	foo()
}

func foo() {
	fmt.Println("Executing foo")

	defer untrace(trace("bar"))
	bar()
}

func bar() {
	fmt.Println("Executing bar")
}

func trace(s string) string {
	fmt.Println("Tracing " + s)
	return s
}

func untrace(s string) {
	fmt.Println("Untracing " + s)
}
