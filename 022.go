package main

import (
  "fmt"
  "sort"
	"strings"
  "os"
  "io/ioutil"
)

func score(name string) int {
  s := 0
  for i := 0; i < len(name); i++ {
    s += int(name[i]) - 64
  }
  return s
}

func main() {
  // Load names from file.
	file, err := os.Open("data/022.txt")
	if err != nil {
			panic(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
  if err != nil {
    panic(err)
  }
	var names []string = strings.Split(string(b), "\",\"")

  // Remove " from first/last names.
  lastIdx := len(names)-1
  names[lastIdx] = names[lastIdx][:len(names[lastIdx])-2]
  names[0] = names[0][1:]

  // Sort names.
  sort.Strings(names)

  // Calculate value for each name.
  sum := 0
  for i, name := range names {
    sum += (i+1) * score(name)
  }

  fmt.Println(sum)
}
