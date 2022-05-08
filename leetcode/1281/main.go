package main

import "fmt"

func subtractProductAndSum(n int) int {
	d := 10
	sum, pro := 0, 0

	first := true
	for n > 0 {
		a := n % d
		sum += a
		if first {
			pro += a
			first = false
		} else {
			pro *= a
		}

		n /= d
	}

	return pro - sum
}

func main() {
	rst := subtractProductAndSum(234)
	fmt.Println(rst)
}
