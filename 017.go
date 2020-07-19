package main

import (
  "fmt"
  "strings"
)

func numberToEnglish(n int) string {
  ones := [20]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
  tens := [10]string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

  if 0 <= n && n < 20 {
    return ones[n]
  } else if 20 <= n && n < 100 && n % 10 == 0 {
    return tens[n/10]
  } else if 20 < n && n < 100 {
    return tens[n / 10] + "-" + ones[n % 10]
  } else if 100 <= n && n < 1000 && n % 100 == 0 {
    return ones[n / 100] + " hundred"
  } else if 100 < n && n < 1000 {
    return ones[n / 100] + " hundred and " + numberToEnglish(n % 100)
  } else if n == 1000 {
    return "one thousand"
  }

  return ""
}

func main() {
  n := 0
  for i := 1; i <= 1000; i++ {
    english := numberToEnglish(i)
    english = strings.Replace(english, " ", "", -1)
    english = strings.Replace(english, "-", "", -1)
    n += len(english)
  }
  fmt.Println(n)
}
