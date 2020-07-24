package main

import (
  "fmt"
  "strconv"
  "math"
)

func isPandigital(n int) bool {
  seen := map[int]bool{}
  for n > 0 {
    if seen[n%10] {
      return false
    }
    seen[n%10] = true
    n /= 10
  }
  if seen[0] { return false }
  for i := 1; i <= 9; i++ {
    if !seen[i] {
      return false
    }
  }
  return true
}

func concatenatedProduct(x, n int) int {
  res := ""
  for i := 1; i <= n; i++ {
    res += strconv.Itoa(x*i)
  }
  y, _ := strconv.Atoi(res)
  return y
}

func main() {
  max := 0
  for n := 2; n <= 9; n++ {
    for x := 1; x < int(math.Pow(10.0, float64(8/n+1))); x++ {
      y := concatenatedProduct(x, n)
      if isPandigital(y) && y > max {
        max = y
      }
    }
  }
  fmt.Println(max)
}
