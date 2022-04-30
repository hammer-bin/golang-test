package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	fmt.Println(g)
	fmt.Println(s)

	i := 0
	j := 0
	num := 0
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			i++
			j++
			num++
		} else {
			j++
		}
	}

	return num

	/*var rst int
	for _, v := range g {
		for i, a := range s {
			if v == a {
				rst++
				s = s[i+1 : len(s)]
				break

			}
		}
	}
	return rst*/
}

func main() {
	g := []int{10, 9, 8, 7}
	s := []int{1, 2, 5, 6}

	a := findContentChildren(g, s)
	fmt.Println(a)
}
