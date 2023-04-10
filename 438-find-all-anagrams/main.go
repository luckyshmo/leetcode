package main

import (
	"fmt"
)

func main() {
	fmt.Println(findAnagrams("abcababccbaab", "ab"))
}

func findAnagrams(st, t string) (res []int) {
	s := []rune(st)
	m := make(map[rune]int)
	for _, c := range t {
		m[c] = m[c] + 1
	}

	// maintain a counter to check whether match the target string.
	// the length of the substring which match the target string.
	counter := len(m)
	begin, end := 0, 0

	for end < len(s) {
		c := s[end]
		if v, ok := m[c]; ok {
			m[c] = v - 1
			if v-1 == 0 {
				counter--
			}
		}
		end++

		//increase begin pointer to make it invalid/valid again
		for counter == 0 { // counter condition. different question may have different condition
			tempc := s[begin]
			if v, ok := m[tempc]; ok {
				m[tempc] = v + 1
				if v+1 > 0 {
					counter++
				}
			}
			/* save / update(min/max) the result if find a target*/
			// result collections or result int value
			if end-begin == len(t) {
				res = append(res, begin)
			}
			begin++
		}
	}
	return
}

// SLOW
func findAnagramsSlow(s string, p string) (res []int) {
	h := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		h[p[i]]++
	}

	for i := 0; i <= len(s)-len(p); i++ {
		j, ok := isAnagramSlow(s[i:i+len(p)], h)
		if ok {
			res = append(res, j+i)
		}
		i += j
	}
	return
}

func isAnagramSlow(s string, h map[byte]int) (i int, isA bool) {
	for i := 0; i < len(s); i++ {
		if _, ok := h[s[i]]; ok {
			h[s[i]]--
			defer func(j int) {
				h[s[j]]++
			}(i)
			if h[s[i]] < 0 {
				return 0, false
			}
			continue
		}
		return i, false
	}
	return 0, true
}
