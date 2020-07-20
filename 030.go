package main

import (
  "fmt"
  "math"
)

func pow(base, exp int) int {
  return int(math.Pow(float64(base), float64(exp)))
}

func sumFifthPower(n int) int {
  sum := 0
  for n > 0 {
    sum += pow(n%10, 5)
    n /= 10
  }
  return sum
}

func main() {
  sum := 0
  for i := 2; i < 1e6; i++ {
    if sumFifthPower(i) == i {
      sum += i
    }
  }
  fmt.Println(sum)
}
