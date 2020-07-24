package main

import (
  "fmt"
  "strconv"
)

func isPalindrome(num string) bool {
	for i := 0; i < len(num)/2; i++ {
		if num[i] != num[len(num)-i-1] {
			return false
		}
	}
	return true
}

func main() {
  sum := 0
  for i := 0; i < 1e6; i++ {
    if isPalindrome(strconv.Itoa(i)) && isPalindrome(strconv.FormatInt(int64(i), 2)) {
      sum += i
    }
  }
  fmt.Println(sum)
}
