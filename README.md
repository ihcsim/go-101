Golang Tutorial
===============

Some basics on golang including answers to exercises on the golang tour.

To get all the codes, clone this repository to your `$GOPATH` workspace,

```
$ go get github.com/ihcsim/golang-tutorial
```

Exercise | Package | Exercise URL
-------- | ------- | ------------
Loops and Functions | `loop` | https://tour.golang.org/flowcontrol/8
Slices | `picture` | https://tour.golang.org/moretypes/14
Maps | `maps` | https://tour.golang.org/moretypes/19
Fibonacci closure | `closure` | https://tour.golang.org/moretypes/22
Stringers | `interface` | https://tour.golang.org/methods/7
Errors | `error` | https://tour.golang.org/methods/9
Readers | `reader` | https://tour.golang.org/methods/11
rot13Reader | `rot13reader` | https://tour.golang.org/methods/12
Http Handlers | `httphandlers` | https://tour.golang.org/methods/14
Images | `images` | https://tour.golang.org/methods/16
Equivalent Binary Trees | `binarytree` | https://tour.golang.org/concurrency/7

To run the answer of an exercise,

```
$ cd github.com/ihcsim/golang-tutorial/<folder_of_exercise>
$ go run main.go
```
 
To fetch dependencies:

1. For picture exercise, `$ go get golang.org/x/tour/pic`
2. For maps exercise, `$ go get golang.org/x/tour/wc`
3. For reader exercise, `$ go get golang.org/x/tour/reader`

# Other Examples

Package | Description
------- | -----------
`stack` | A stack data structure demonstrating the usage of `interface{}`.
`discountcalculator` | A discount calculator demonstrating the usage of function values and closures.
`bigdigits` | A program that takes a numeric input and magnifies the number to stdout. Run `$GOPATH/bin/bigdigits --help` for usages.
