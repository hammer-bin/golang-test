package main

import "fmt"

func longestPalindrome(s string) int {
	l := 0

	m := make(map[rune]int)

	// 각 글자의 카운트
	for _, r := range s {
		m[r]++
	}

	odd := false
	for _, c := range m {
		if c%2 == 0 {
			l += c
		} else {
			l += (c / 2) * 2
			odd = true
		}
	}

	if odd {
		l++
	}

	return l
}

func main() {
	rst := longestPalindrome("abccccdd")
	fmt.Println(rst)
}
