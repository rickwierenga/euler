package main

import "fmt"

func largestPrimeFactor(n int) (int) {
  x := 2

  for x * x < n {
    for n % x == 0 {
      n /= x
   }
   x++
  }

  return n
}

func main() {
  fmt.Println(largestPrimeFactor(600851475143))
}
