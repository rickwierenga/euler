package main

import (
  "fmt"
  "math/big"
)

func factorial(n int64) *big.Int {
  var x = new(big.Int)
  x.MulRange(1, n)
  return x
}

func combination(n int64, k int64) *big.Int {
  var num = factorial(n)
  var den = new(big.Int)
  den.Mul(factorial(n - k), factorial(k))
  var res = new(big.Int)
  return res.Div(num, den)
}

func main() {
  fmt.Println(combination(20 + 20, 20))
}
