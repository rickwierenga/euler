package main

import (
  "fmt"
  "math"
)

func isPrime(x int) bool {
  if x == 1 { return false }
  for i := 2; i < int(math.Sqrt(float64(x))) + 1; i++ {
    if x % i == 0 {
      return false
    }
  }
  return true
}

func isTruncatablePrime(n int) bool {
  if !isPrime(n) {
    return false
  }

  numDigits := int(math.Floor(math.Log10(float64(n)))) + 1

  // left -> right
  for j := 1; j < numDigits; j++ {
    if !isPrime(n / int(math.Pow(10.0, float64(j)))) {
      return false
    }
  }

  // right -> left
  for j := 1; j < numDigits; j++ {
    if !isPrime(n % int(math.Pow(10.0, float64(j)))) {
      return false
    }
  }

  return true
}

func main() {
  sum := 0
  num := 0
  i := 10
  for num < 11 {
    if isTruncatablePrime(i) {
      sum += i
      num++
      fmt.Println(i)
    }
    i++
  }
  fmt.Println(sum)
}
