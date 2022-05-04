package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3}

	plusTwo(s)
}

func plusOne(digits []int) []int {
	n := convInt(digits)

	n++

	return convArr(n)
}

func convInt(digits []int) int {
	no := 0
	for i := 0; i < len(digits); i++ {
		no *= 10
		no += digits[i]
	}
	return no
}

func convArr(n int) []int {
	s := make([]int, 0)

	s1 := 10
	for n != 0 {
		x := n % s1
		s = append([]int{x}, s...)
		n = n / s1

	}

	fmt.Println(s)

	return s
}

func plusTwo(digits []int) []int {
	//[1,2,3] --> [1,2,4]
	//[9,9] --> [1,0,0]
	carry := true
	for i := len(digits) - 1; i >= 0; i-- {
		if carry {
			if digits[i] == 9 {
				digits[i] = 0
				carry = true
			} else {
				digits[i]++
				carry = false
			}
		} else {
			break
		}

		if carry {
			digits = append([]int{1}, digits...)
		}

	}
	fmt.Println(digits)
	return digits
}
