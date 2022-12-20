package main

import "fmt"

func main() {
	fmt.Print(smallestValue(42))
}

func smallestValue(n int) int {
	for ; ; n = getSumOfPrimeFactors(n) {
		newN := getSumOfPrimeFactors(n)
		if newN == n {
			return n
		}
	}
}

func getSumOfPrimeFactors(n int) int {
	ans := 0

	var rec func()
	rec = func() {
		for i := 2; i < n/2; i++ {
			if n%i == 0 {
				ans += i
				n = n / i
				rec()
				break
			}
		}
	}
	rec()

	return ans + n
}
