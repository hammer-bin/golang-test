package main

import "fmt"

func reverse(a []int) []int {
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
	return a
}

func main() {
	fmt.Printf("Reversed = %d", reverse([]int{1, 2, 3, 4}))
}
