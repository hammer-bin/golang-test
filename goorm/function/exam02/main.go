package main

import (
	"fmt"
)

func inputNums() []int {
	var score int
	nums := make([]int, 0)

	for i := 0; i < 5; i++ {
		fmt.Scanf("%d\n", &score)
		nums = append(nums, score)
	}
	return nums
}

func calExam(arr ...int) (sum, high, low int) {
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if arr[i] >= 90 {
			high++
		} else if arr[i] < 70 {
			low++
		}
	}
	return sum, high, low
}

func printResult(sum, high, low int) {
	var result bool = true

	if sum <= 400 {
		result = false
		fmt.Printf("총점이 %d점 미만입니다.\n", 400)
	}
	if high < 2 {
		result = false
		fmt.Printf("%d점 이상 과목 수가 %d개 미만입니다.\n", 90, 2)
	}
	if low >= 1 {
		result = false
		fmt.Printf("%d점 미만 과목이 있습니다.\n", 70)
	}

	if result == true {
		fmt.Println("아이패드를 살 수 있습니다.\n")
	} else {
		fmt.Println("아이패드를 살 수 없습니다.\n")
	}
}

func main() {
	nums := inputNums()
	printResult(calExam(nums...))
}
