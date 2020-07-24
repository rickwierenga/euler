package main

import (
  "fmt"
  "math"
)

func isPrime(x int) bool {
  for i := 2; i < int(math.Sqrt(float64(x))) + 1; i++ {
    if x % i == 0 {
      return false
    }
  }
  return true
}

func isPandigital(n int) bool {
  seen := map[int]bool{}
  l := int(math.Log10(float64(n))) + 1
  for n > 0 {
    if seen[n%10] {
      return false
    }
    seen[n%10] = true
    n /= 10
  }
  if seen[0] { return false }
  for i := 1; i <= l; i++ {
    if !seen[i] {
      return false
    }
  }
  return true
}

func main() {
  fmt.Println("This one takes a while")
  for i := int(1e9) - 1; i > 0; i -= 2 {
    if isPandigital(i) && isPrime(i) {
      fmt.Println(i)
      break
    }
  }
}
