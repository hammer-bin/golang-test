package main

import "fmt"

func numIdenticalPairs(nums []int) int {
	a := 0
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				if i != j && i < j {
					a++

					fmt.Println("i=", i, "num:", nums[i], " j=", j, "num:", nums[j])
				}
			}
		}
	}
	return a
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(nums)
	rst := numIdenticalPairs(nums)
	fmt.Println(rst)
}
