package main

import (
  "fmt"
  "math"
)

func sol(p int) [][]int {
  res := [][]int{}
  for a := 1; a < p; a++ {
    for b := a; b < p; b++ {
      // Must be a pythagorean triple.
      c := math.Sqrt(float64(a*a) + float64(b*b))
      if float64(p-a-b) == c {
        res = append(res, []int{a, b, int(c)})
      }
    }
  }
  return res
}

func main() {
  pMax := 0
  lMax := 0
  for p := 25; p <= 1000; p++ {
    l := len(sol(p))
    if l > lMax {
      pMax = p
      lMax = l
    }
  }
  fmt.Println(pMax)
}
