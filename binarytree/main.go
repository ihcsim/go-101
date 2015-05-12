package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	RecursiveWalk(t, ch)
	close(ch)
}

func RecursiveWalk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		RecursiveWalk(t.Left, ch)
	}

	ch <- t.Value

	if t.Right != nil {
		RecursiveWalk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	channel1 := make(chan int)
	go Walk(t1, channel1)

	channel2 := make(chan int)
	go Walk(t2, channel2)

	isTrue := true
	for {
		val1, ok1 := <-channel1
		val2, ok2 := <-channel2

		if ok1 && ok2 && val1 != val2 {
			isTrue = false
			break
		}

		if (!ok1 && !ok2) || (ok1 != ok2) {
			break
		}
	}

	return isTrue
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for value := range ch {
		fmt.Println(value)
	}

	fmt.Printf("tree1 is the same as tree2? %t\n", Same(tree.New(1), tree.New(1)))
}
