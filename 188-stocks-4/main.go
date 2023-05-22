package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println()
}

func maxProfit(k int, prices []int) int {
	buy, sell := make([]int, k+1), make([]int, k+1)
	for i := range buy {
		buy[i] = math.MinInt
	}
	for _, price := range prices {
		for i := 1; i <= k; i++ {
			buy[i] = max(buy[i], sell[i-1]-price)
			sell[i] = max(sell[i], buy[i]+price)
		}
	}
	return sell[k]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
