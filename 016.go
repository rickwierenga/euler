package main

import (
  "fmt"
  "strconv"
  "math/big"
)

func main() {
  var base = big.NewInt(2)
  var exp = big.NewInt(1000)
  var x = new(big.Int)
  x.Exp(base, exp, nil)
  str := x.String()

  res := 0
  for i := 0; i < len(str); i++ {
    y, _ := strconv.Atoi(string(str[i]))
    res += y
  }
  fmt.Println(res)
}
