package main

import "fmt"

func main() {
	atm := Constructor()
	atm.Deposit([]int{0, 0, 1, 2, 1})
	fmt.Println(atm.Withdraw(600))
}

var indexToBanknote []int = []int{20, 50, 100, 200, 500}

type ATM struct {
	count []int
}

func Constructor() ATM {
	return ATM{
		count: make([]int, 5),
	}
}

func (this *ATM) Deposit(banknotesCount []int) {
	for i, count := range banknotesCount {
		this.count[i] += count
	}
}

func (this *ATM) Withdraw(amount int) []int {
	res := make([]int, 5)

	for i := 4; i >= 0 && amount != 0; i-- {
		takeThisMany := min(amount/indexToBanknote[i], this.count[i])
		this.count[i] -= takeThisMany
		res[i] = takeThisMany
		amount -= takeThisMany * indexToBanknote[i]
	}

	if amount == 0 {
		return res
	} else {
		for i, count := range res {
			this.count[i] += count
		}
		return []int{-1}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
