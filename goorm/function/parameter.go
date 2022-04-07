package main

import "fmt"

func addOne(num ...int) int {
	var result int

	for i := 0; i < len(num); i++ {
		result += num[i]
	}
	return result
}

func addTwo(num ...int) int {
	var result int

	for _, val := range num {
		result += val
	}

	return result
}

func main() {
	num1, num2, num3, num4, num5 := 1, 2, 3, 4, 5
	nums := []int{10, 20, 30, 40}

	fmt.Println(addOne(num1, num2))
	fmt.Println(addOne(num1, num2, num3))
	fmt.Println(addOne(nums...))
	fmt.Println(addTwo(num3, num4, num5))
	fmt.Println(addTwo(num1, num3, num4, num5))
	fmt.Println(addTwo(nums...))

}
