package helpers

import "math"

func slidingTemplate(st, t string) {
	s := []rune(st)
	m := make(map[rune]int)
	for _, c := range t {
		m[c] = m[c] + 1
	}

	// maintain a counter to check whether match the target string.
	// the length of the substring which match the target string.
	counter, _ := len(m), math.MaxInt
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
			begin++
		}
	}
}
