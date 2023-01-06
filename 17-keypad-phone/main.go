package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("23"))
	// fmt.Println(letterCombinations("23456789"))
}

var rtr = [][]rune{
	nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, {'a', 'b', 'c'}, {'d', 'e', 'f'}, {'g', 'h', 'i'}, {'j', 'k', 'l'}, {'m', 'n', 'o'}, {'p', 'q', 'r', 's'}, {'t', 'u', 'v'}, {'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	return combinations(reverse(digits))
}

func combinations(digits string) (c []string) {
	if len(digits) == 1 {
		runes := rtr[rune(digits[0])]
		resp := make([]string, len(runes))
		for i, r := range runes {
			resp[i] = string(r)
		}
		return resp
	}
	nextComb := combinations(digits[1:])
	runes := rtr[rune(digits[0])]
	resp := make([]string, 0)
	for _, s := range nextComb {
		for _, r := range runes {
			kek := s + string(r)
			resp = append(resp, kek)
		}
	}

	return resp
}

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}
