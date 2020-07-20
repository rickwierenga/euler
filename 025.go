package main

import (
  "fmt"
  "math/big"
)

func main() {
  a := big.NewInt(1)
  b := big.NewInt(2)//2?
  i := 2
  for len(a.String()) < 1000 {
    a, b = b, a.Add(a, b)
    i++
  }
  fmt.Println(i)
}
