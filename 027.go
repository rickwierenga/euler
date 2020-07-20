package main

import (
  "fmt"
  "math"
)

func abs(x int) int {
  return int(math.Abs(float64(x)))
}

func isPrime(x int) bool {
  for i := 2; i < int(math.Sqrt(float64(x))) + 1; i++ {
    if x % i == 0 {
      return false
    }
  }
  return true
}

func main() {
  maxA := 0
  maxB := 0
  maxN := 0

  // a,b must be odd.
  for a := -999; a <= 999; a += 2 {
    for b := -999; b <= 999; b += 2 {
      n := 0
      for isPrime(abs(n * n + a * n + b)) {
        n++
      }

      if n > maxN {
        maxA = a
        maxB = b
        maxN = n
      }
    }
  }

  fmt.Println(maxA, maxB, maxA * maxB)
}
