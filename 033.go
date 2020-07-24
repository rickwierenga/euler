package main

import "fmt"

func gcd(a, b int) int {
  if b == 0 {
    return a
  }
  return gcd(b, a % b)
}

func main() {
  num := 1
  den := 1
	for i := 10; i < 100; i++ {
		for j := i + 1; j < 100; j++ {
			if float64(i%10)==float64(j/10) &&
				float64(i/10)/float64(j%10) == float64(i)/float64(j) {
          num *= i
          den *= j
			}
		}
	}
  fmt.Println(den / gcd(num, den))
}
