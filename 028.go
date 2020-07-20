package main

import "fmt"

func sum(n int) int {
  if n == 0 {
    return 1
  }
  return 4 * (2 * n + 1) * (2 * n + 1) - 12 * n + sum(n - 1)
}

func main() {
  fmt.Println(sum(500))
}

