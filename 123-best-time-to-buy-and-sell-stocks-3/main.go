package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 1, 2, 3}))
}

/*
oneBuy keeps track of the lowest price, and oneBuyOneSell keeps track of the biggest profit we could get.
Then the tricky part comes, how to handle the twoBuy? Suppose in real life, you have bought and sold a stock and made $100 dollar profit. When you want to purchase a stock which costs you $300 dollars, how would you think this? You must think, um, I have made $100 profit, so I think this $300 dollar stock is worth $200 FOR ME since I have hold $100 for free.
There we go, you got the idea how we calculate twoBuy!! We just minimize the cost again!! The twoBuyTwoSell is just making as much profit as possible.
*/

func maxProfit(prices []int) int {
	oneBuyOneSell, twoBuyTwoSell := 0, 0
	oneBuy, twoBuy := math.MaxInt, math.MaxInt
	for _, p := range prices {
		oneBuy = min(oneBuy, p)
		oneBuyOneSell = max(oneBuyOneSell, p-oneBuy)
		twoBuy = min(twoBuy, p-oneBuyOneSell)
		twoBuyTwoSell = max(twoBuyTwoSell, p-twoBuy)
		fmt.Println("kek")
	}
	return twoBuyTwoSell
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
