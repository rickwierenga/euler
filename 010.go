package main

import (
  "fmt"
)

func sieve(n int) []bool {
  // Sieve of Eratosthenes
  primes := make([]bool, n+1)
  for i := 2; i <= n; i++ {
    primes[i] = true
  }

  for p := 2; p*p < n; p++ {
    if primes[p] {
      for i := 2 * p; i <= n; i += p {
        primes[i] = false
      }
    }
  }

  return primes
}

func main() {
  sum := 0
  for i, x := range sieve(2000000) {
    if x {
      sum += i
    }
  }
  fmt.Println(sum)
}
