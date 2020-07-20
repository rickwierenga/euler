package main

import "fmt"

func main() {
  //algo: https://math.stackexchange.com/questions/46864/best-way-of-computing-the-decimal-representation-of-a-fraction-with-an-arbitrary
  max := 0

  for d := 1; d < 1000; d++ {
    if max > d { continue }
    seen := map[int]bool{}
    // Always starts with "0." since d > 1, so skip first iter.
    num := 10
    r := 0

    for !seen[num] {
      r = num / d
      seen[num] = true
      num = 10 * (num - r * d)
    }

    if len(seen) > max {
      max = len(seen)
    }
  }

  fmt.Println(max+1)
}
