package main

/*
Input: num = 2932
Output: 52
Explanation: Some possible pairs [new1, new2] are [29, 23], [223, 9], etc.
The minimum sum can be obtained by the pair [29, 23]: 29 + 23 = 52.
*/
import (
	"fmt"
	"sort"
)

func minimumSum(num int) int {
	a := make([]int, 4)
	for i := 0; i < len(a); i++ {
		a[i] = num % 10
		num /= 10
	}
	fmt.Println(a)
	sort.Ints(a)
	fmt.Println(a)
	a1 := a[0]*10 + a[2]
	b1 := a[1]*10 + a[3]

	return a1 + b1
}

func findPeakElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)-2]
}

func main() {
	rst := minimumSum(2932)
	fmt.Println(rst)
}
