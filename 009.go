package main

import (
  "fmt"
  "math"
)

func main() {
  for a := 1; a < 1000; a++ {
    for b := a; b < (1000 - a); b++ {
      c := math.Sqrt(float64(a * a + b * b))
      if math.Trunc(c) == c && a + b + int(c) == 1000 {
        fmt.Println(a * b * int(c))
        break
      }
    }
  }
}
