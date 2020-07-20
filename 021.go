package main

import "fmt"

func d(n int) int {
  var sum int
  for i := 2; i * i <= n; i++ {
    if n % i == 0 {
      sum += i
      if !(i * i == n) {
        sum += (n / i)
      }
    }
  }
  return sum + 1
}

func isAmicable(n int) bool {
  return n == d(d(n)) && !(d(n) == n)
}

func main() {
  sum := 0
  for i := 0; i < 10000; i++ {
    if isAmicable(i) {
      sum += i
    }
  }
  fmt.Println(sum)
}
