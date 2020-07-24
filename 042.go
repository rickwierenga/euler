package main

import (
  "fmt"
	"strings"
  "os"
  "sort"
  "io/ioutil"
)

func value(name string) int {
  s := 0
  for i := 0; i < len(name); i++ {
    s += int(name[i]) - 64
  }
  return s
}

func contains(a []int, b int) bool {
  i := sort.SearchInts(a, b)
  return i < len(a) && a[i] == b
}

func main() {
  // Load names from file.
	file, err := os.Open("data/042.txt")
	if err != nil {
			panic(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
  if err != nil {
    panic(err)
  }
	var words []string = strings.Split(string(b), "\n")

  // Pre-calculate triangle numbers
  triangleNumbers := []int{}
  for n := 1; n < 50; n++ {
    triangleNumbers = append(triangleNumbers, n * (n + 1) / 2)
  }

  // Count
  i := 0
  for _, word := range words {
    if contains(triangleNumbers, value(word)) {
      i++
    }
  }

  fmt.Println(i)
}
