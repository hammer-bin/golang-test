package main

import "fmt"

func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		} else if nums[i] < target {
			if len(nums)-1 <= i {
				return i + 1
			}
		}
	}
	return 0
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 7

	rst := searchInsert(nums, target)

	fmt.Println(rst)
}
