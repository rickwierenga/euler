package main

import "fmt"

func isPandigital(n int) bool {
  seen := map[int]bool{}
  for n > 0 {
    if seen[n%10] {
      return false
    }
    seen[n%10] = true
    n /= 10
  }
  for i := 1; i <= 9; i++ {
    if !seen[i] {
      return false
    }
  }
  return true
}

func main() {
  sum := 0
  seen := map[int]bool{}

  // 2,3 digit
  for i := 10; i < 100; i++ {
    for j := 100; j < 1000; j++ {
      if i*j < 1e4 && !seen[i*j] && isPandigital(i * 1e7 + j * 1e4 + i * j) {
        seen[i*j] = true
        sum += i*j
      }
    }
  }

  // 1,4 digit
  for i := 1; i < 10; i++ {
    for j := 1000; j < 10000; j++ {
      if i*j < 1e4 && !seen[i*j] && isPandigital(i * 1e8 + j * 1e4 + i * j) {
        seen[i*j] = true
        sum += i*j
      }
    }
  }

  fmt.Println(sum)
}
