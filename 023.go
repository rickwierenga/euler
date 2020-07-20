package main

import "fmt"

func sumFactors(n int) int {
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

func isAbundant(n int) bool {
  return sumFactors(n) > n
}

// Sum of abundant numbers?
func isSum(n int, ab []int) bool {
  i := 0
  j := len(ab) - 1
  for i <= j {
    diff := ab[i] + ab[j]
    if diff == n {
      return true
    } else if diff > n {
      j--
    } else { // diff < n
      i++
    }
  }
  return false
}

func main() {
  ab := []int{}
  sum := 0
  for i := 0; i < 28123; i++ {
    if !isSum(i, ab) {
      sum += i
    }
    if isAbundant(i) {
      ab = append(ab, i)
    }
  }
  fmt.Println(sum)
}
