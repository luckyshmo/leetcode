package main

func dailyTemperaturesSlow(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		for j := i + 1; j < len(temperatures); j++ {
			if temperatures[i] < temperatures[j] {
				ans[i] = j - i
				break
			}
		}
	}

	return ans
}
