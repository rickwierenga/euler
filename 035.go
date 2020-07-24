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

func rotate(n int) []int {
  res := []int{n}
  numDigits := math.Floor(math.Log10(float64(n)))
  x := int(math.Pow(10, numDigits))
  for i := 0; i < int(numDigits); i++ {
    res = append(res, (res[i]%10)*x + res[i] / 10)
  }
  return res
}

func isCircPrime(n int) bool {
  for _, i := range rotate(n) {
    if !isPrime(i) {
      return false
    }
  }
  return true
}

func main() {
  num := 0
  for i := 2; i<1e6; i++ {
    if isCircPrime(i) {
      num++
    }
  }
  fmt.Println(num)
}
