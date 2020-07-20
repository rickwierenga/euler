package main

import (
  "fmt"
  "math/big"
)

func pow(base, exp int) *big.Int {
  return new(big.Int).Exp(big.NewInt(int64(base)), big.NewInt(int64(exp)), nil)
}

func main() {
  seen := map[string]bool{}
  for a := 2; a <= 100; a++ {
    for b := 2; b <= 100; b++ {
      p := pow(a, b).String()
      if !seen[p] {
        seen[p] = true
      }
    }
  }
  fmt.Println(len(seen))
}
