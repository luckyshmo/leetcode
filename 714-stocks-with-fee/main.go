package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxProfit([]int{1, 3, 5, 7}, 1))
}

func maxProfit(prices []int, fee int) (cash int) {
	hold := -prices[0]
	for i := 1; i < len(prices); i++ {
		cash = max(cash, hold+prices[i]-fee)
		hold = max(hold, cash-prices[i])
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
