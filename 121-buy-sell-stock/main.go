package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxProfit([]int{1, 2, 4}))
	fmt.Println(maxProfitRange([]int{7, 1, 5, 3, 6, 4}))
}

// For some reason twice as slow as iteration....
func maxProfitRange(prices []int) int {
	diff := 0
	min := prices[0]

	for _, p := range prices {
		if p < min {
			min = p
		}

		if p-min > diff {
			diff = p - min
		}
	}

	return diff
}

func maxProfit(prices []int) int {
	diff := 0
	min := prices[0]

	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}

		if prices[i]-min > diff {
			diff = prices[i] - min
		}
	}

	return diff
}
