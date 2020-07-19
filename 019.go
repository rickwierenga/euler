package main

import "fmt"

func isLeapYear(y int) bool {
  if y % 100 == 0 {
    if y % 400 == 0 {
      return true
    } else {
      return false
    }
  } else if y % 4 == 0 {
    return true
  }
  return false
}

func daysInMonth(m int, y int) int {
  if m == 2 && isLeapYear(y) {
    return 29
  }

  switch m {
  case 1,3,5,7,8,10,12:
    return 31
  case 2:
    return 28
  case 4,6,9,11:
    return 30
  }

  panic(fmt.Sprintf("unkown month, %v", m))
}

func main() {
  numDays := 0
  sundays := 0
  for y := 1901; y <= 2000; y++ {
    for m := 1; m <= 12; m++ {
      if numDays % 7 == 0 {
        sundays++
      }
      numDays += daysInMonth(m, y)
    }
  }
  fmt.Println(sundays)
}
