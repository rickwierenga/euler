package main

import (
  "fmt"
  "strconv"
)

func isPalindrome(num string) (bool) {
	for i := 0; i < len(num)/2; i++ {
		if num[i] != num[len(num)-i-1] {
			return false
		}
	}
	return true
}

func main() {
  res := 0

  for i := 1000; i > 100; i-- {
    for j := 1000; j > i; j-- {
      prod := i * j
      if isPalindrome(strconv.Itoa(prod)) && prod > res {
        res = prod
      } else if res > prod {
        break
      }
    }
  }

  fmt.Println(res)
}
