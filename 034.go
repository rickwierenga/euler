package main

import "fmt"

func fact(n int) int {
  res := 1
  for n > 0 {
    res *= n
    n -= 1
  }
  return res
}

func isCurious(n int) bool {
  sum := 0
  i := n
  for i > 0 {
    sum += fact(i % 10)
    i /= 10
  }
  return sum == n
}

func main() {
  sum := 0
  for i := 10; i < 100000; i++ {
    if isCurious(i) {
      sum += i
    }
  }
  fmt.Println(sum)
}
