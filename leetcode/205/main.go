package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	stmap := make([]byte, 256)
	tsmap := make([]byte, 256)

	for i := 0; i < len(s); i++ {
		a, b := s[i], t[i]
		if stmap[a] != 0 {
			if stmap[a] != b {
				return false
			}
		} else if tsmap[b] != 0 {
			if tsmap[b] != a {
				return false
			}
		} else {
			stmap[a] = b
			tsmap[b] = a
		}
	}
	return true
}

func main() {
	rst := isIsomorphic("paper", "biber")

	fmt.Println(rst)
}
