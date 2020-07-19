package main

import (
  "fmt"
  "math/big"
  "strconv"
)

func main() {
  var x = new(big.Int)
  x.MulRange(1,100)
  str := x.String()
  sum := 0
  for i := 0; i < len(str); i++ {
    y, _ := strconv.Atoi(string(str[i]))
    sum += y
  }
  fmt.Println(sum)
}
