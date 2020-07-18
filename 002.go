package main

import "fmt"

func main() {
  a, b := 0, 1
  c := a + b
  sum := 0

  for c <= 4000000 {
    if c % 2 == 0 {
      sum += c
    }

    a, b = b, c
    c = a + b
  }

  fmt.Println(sum)
}
