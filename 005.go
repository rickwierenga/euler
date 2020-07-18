package main

import "fmt"

func gcd(a int, b int) int {
  // https://en.wikipedia.org/wiki/Euclidean_algorithm
  if b == 0 {
    return a
  } else {
    return gcd(b, a % b)
  }
}

func main() {
  res := 1

  // Product of all nonmutual prime factors.
  for i := 2; i < 20; i++ {
    res = (res * i) / gcd(res, i)
  }

  fmt.Println(res)
}
