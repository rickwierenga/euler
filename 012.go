package main

import "fmt"

func numFactors(n int) int {
  i := 1
  num := 0
  for i * i < n {
    if n % i == 0 {
      num++
    }
    i++
  }
  return 2 * num
}

func triangleNum(n int) int {
  return (n + 1) * n / 2
}

func main() {
  i := 1
  for numFactors(triangleNum(i)) < 500 {
    i++
  }
  fmt.Println(triangleNum(i))
}
