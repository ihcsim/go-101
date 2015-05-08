package main

import "fmt"

func main(){
  f := fibonacci()
  for i := 0; i < 10; i++{
    fmt.Println(f())
  }
}

func fibonacci() func() int{
  last := 0
  current := 0
  return func() int {
    if current == 0 {
      current = current + 1
      return 1
    }
  
    next := last + current
    last, current = current, next
    return next
  }
}
