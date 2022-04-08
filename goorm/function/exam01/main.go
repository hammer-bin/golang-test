package main

import "fmt"

func bubbleSort(r []int) {
	var temp int
	for i := len(r) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if r[j] < r[j+1] {
				temp = r[j]
				r[j] = r[j+1]
				r[j+1] = temp
			}
		}
	}
}

func inputNums() []int {
	var count, num int
	var slice []int
	fmt.Scanf("%d", &count)
	for i := 0; i < count; i++ {
		fmt.Scanf("%d\n", &num)
		slice = append(slice, num)
	}
	return slice
}

func outputNums(r []int) {
	for i := 0; i < len(r); i++ {
		fmt.Printf("%d ", r[i])
	}
}

func main() {
	fmt.Println("yuyuyu")
	nums := inputNums()
	bubbleSort(nums)
	outputNums(nums)
}
