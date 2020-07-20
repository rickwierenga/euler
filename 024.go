package main

import "fmt"

func fact(n int) int {
  x := 1
  for i := 1; i <= n; i++ {
    x *= i
  }
  return x
}

func remove(slice []int, i int) []int {
  copy(slice[i:], slice[i+1:])
  return slice[:len(slice)-1]
}

func main() {
  numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
  N := len(numbers)
  res := []int{}
  remain := 1000000 - 1

  for i := 1; i < N; i++ {
    j := remain / fact(N - i)
    remain = remain % fact(N - i)
    res = append(res, numbers[j])
    numbers = remove(numbers, j)
  }

  for i := 0; i < len(res); i++ {
    fmt.Printf("%d", res[i])
  }

  for i := 0; i < len(numbers); i++ {
    fmt.Printf("%d", numbers[i])
  }

  fmt.Printf("\n")
}
