package main

import (
  "fmt"
  "strconv"
  "strings"
)

func max(x, y int) int {
  if x > y {
    return x
  }
  return y
}

func main() {
  pyramidString := `75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23`
  var pyramid [][]int

  // Parse
  rows := strings.Split(pyramidString, "\n")
  for i, row := range rows {
    pyramid = append(pyramid, []int{})
    for _, node := range strings.Split(row, " ") {
      x, _ := strconv.Atoi(node)
      pyramid[i] = append(pyramid[i], x)
    }
  }

  // Bottom up search, reduce each row to the highest value path.
  for i := len(pyramid) - 1; i > 0; i-- {
    parentNodes := pyramid[i-1]
    nodes := pyramid[i]
    for j := 0; j < len(parentNodes); j++ {
      parentNodes[j] = parentNodes[j] + max(nodes[j], nodes[j+1])
    }
  }
  fmt.Println(pyramid[0][0])
}
