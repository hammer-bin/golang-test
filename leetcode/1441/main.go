package main

import "fmt"

func buildArray(target []int, n int) []string {
	var orderSlice []string

	order := 1
	end := target[len(target)-1]
	for i := 0; i < n; i++ {

		if end >= order {
			if order == target[i] {
				orderSlice = append(orderSlice, "Push")
			} else {
				orderSlice = append(orderSlice, "Push")
				orderSlice = append(orderSlice, "Pop")
				i--
				n--
			}
			order++
		}

	}
	return orderSlice

}

func main() {
	aa := []int{1, 2}
	bb := 4
	rst := buildArray(aa, bb)
	fmt.Println(rst)
}
