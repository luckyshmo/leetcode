package main

func main() {
	println(lengthOfLongestSubstring("abccc123455678"))
}

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	max := 0
	current := 0

	for i := 0; i < len(s); i++ {
		char := s[i]
		if repIndex, ok := m[char]; ok {
			i = repIndex
			m = make(map[byte]int)
			if current > max {
				max = current
			}
			current = 0
			continue
		}

		m[char] = i
		current++
		if current > max {
			max = current
		}
	}

	return max
}
