package main

import "fmt"

func main() {
  table := make([]int, 201)
  table[0] = 1
  coins := []int{1,2,5,10,20,50,100,200}

  for i := 0; i < len(coins); i++ {
    for j := coins[i]; j <= 200; j++ {
      table[j] += table[j - coins[i]]
    }
  }

  fmt.Println(table[200])
}
