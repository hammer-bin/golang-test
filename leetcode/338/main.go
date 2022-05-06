package main

import "fmt"

func countBits(n int) []int {
	var rst []int
	for i := 0; i <= n; i++ {
		sum := 0
		d := i
		for d > 0 {
			if d&1 == 1 {
				sum++
			}
			d = d >> 1
		}
		rst = append(rst, sum)
	}

	return rst
}

func main() {
	rst := countBits(10000000)
	fmt.Println(rst)
}
