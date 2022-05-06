package main

import (
	"fmt"
	"strings"
)

func repeatedSubstringPattern(s string) bool {
	if len(s) == 0 {
		return false
	}

	size := len(s)
	fmt.Println(s + s)
	fmt.Println((s + s)[1 : size*2-1])
	ss := (s + s)[1 : size*2-1]

	return strings.Contains(ss, s)
}

func main() {
	rst := repeatedSubstringPattern("sf")
	fmt.Println(rst)
}
