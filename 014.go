package main

import "fmt"

func chain(n int) int {
  i := 0
  for n != 1 {
    if n % 2 == 0 {
      n /= 2
    } else {
      n = 3 * n + 1
    }
    i++
  }
  return i + 1
}

func main() {
  max := 0
  maxI := 0
  for i := 1; i < 1000000; i++ {
    l := chain(i)
    if l > max {
      max = l
      maxI = i
    }
  }
  fmt.Println(maxI)
}
