package main

import (
	"fmt"
	"math"
	"strconv"
)

func pow(b, e int) int {
	return int(math.Pow(float64(b), float64(e)))
}

// Start index of decimal place with n digits.
func s(n int) int {
	if n <= 0 {
		return 0
	} else {
		return n*9*pow(10, n-1) + s(n-1)
	}
}

func d(n int) int {
	if n == 1 || n == 10 {
		return 1
	}
	logN := int(math.Log10(float64(n)))
	decimalIdx := (n - 1 - s(logN-1)) // since start of log10(n-1)
	number := decimalIdx/logN + n/10 - 1
	digitIdx := decimalIdx % logN
	digit, _ := strconv.Atoi(string(strconv.Itoa(number)[digitIdx]))
	return digit
}

func main() {
	fmt.Println(d(1e0) * d(1e1) * d(1e2) * d(1e3) * d(1e4) * d(1e5) * d(1e6))
}
