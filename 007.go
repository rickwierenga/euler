package main

import (
  "fmt"
  "math"
)

func is_prime(x int) bool {
  for i := 2; i < int(math.Sqrt(float64(x))) + 1; i++ {
    if x % i == 0 {
      return false
    }
  }
  return true
}

func main() {
  num_primes := 0
  n := 0

  for num_primes <= 10001 {
    n += 1
    if is_prime(n) {
      num_primes += 1
    }
  }

  fmt.Println(n)
}
