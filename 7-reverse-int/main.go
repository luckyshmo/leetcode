package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	// fmt.Println(reverse(120))
}

func reverse(x int) int {
	isNegative := false
	if x < 0 {
		x = -x
		isNegative = true
	}
	if x > 2147483647 {
		return 0
	}
	ans := 0
	n := len(fmt.Sprint(x)) - 1
	for factor := int(math.Pow10(n)); x > 0; factor = factor / 10 {
		ans += (x % 10) * factor
		x = x / 10
	}
	if ans > 2147483647 {
		return 0
	}
	if isNegative {
		return -ans
	}
	return ans
}
