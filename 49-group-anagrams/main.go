package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

func groupAnagrams(strs []string) [][]string {
	res := make(map[string][]string)

	for i := 0; i < len(strs); i++ {
		sorted := []rune(strs[i])
		sort.Slice(sorted, func(l1, l2 int) bool {
			return sorted[l1] < sorted[l2]
		})
		sortedString := string(sorted)
		if _, ok := res[sortedString]; !ok {
			res[sortedString] = []string{strs[i]}
			continue
		}
		res[sortedString] = append(res[sortedString], strs[i])
	}

	ans := make([][]string, 0)
	for _, v := range res {
		ans = append(ans, v)
	}

	return ans
}
