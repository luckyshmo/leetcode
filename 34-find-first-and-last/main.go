package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{2, 2}, 2))
}

func searchRange(nums []int, target int) (res []int) {
	base := 0
	res = []int{-1, -1}
	for len(nums) != 0 {
		middleInd := len(nums) / 2
		middle := nums[middleInd]
		if middle == target {
			res = []int{base + middleInd, base + middleInd}

			f := nums[:middleInd]
			fBase := 0
			for len(f) != 0 {
				middleInd := len(f) / 2
				middle := f[middleInd]
				if middle == target {
					res[0] = middleInd + base + fBase
					f = f[:middleInd]
				} else {
					f = f[middleInd+1:]
					fBase += middleInd + 1
				}
			}

			s := nums[middleInd+1:]
			sBase := middleInd + 1
			for len(s) != 0 {
				middleInd := len(s) / 2
				middle := s[middleInd]
				if middle == target {
					res[1] = base + middleInd + sBase
					s = s[middleInd+1:]
					sBase += middleInd + 1
				} else {
					s = s[:middleInd]
				}
			}
			return
		}
		if target > middle {
			nums = nums[middleInd+1:]
			base += middleInd + 1
		} else {
			nums = nums[:middleInd]
		}
	}
	return
}
