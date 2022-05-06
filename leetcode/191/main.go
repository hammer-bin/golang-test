package main

import "fmt"

func hammingWeight(num uint32) int {
	var sum int = 0
	for num > 0 {
		if num&1 == 1 {
			sum++
		}
		num = num >> 1
	}
	return sum
}

func hammingWeight2(num uint32) int {
	result := 0
	for i := 0; i < 32; i++ {
		if (num&(1<<i))>>i == 1 {
			result++
		}
	}
	return result
}

func hammingWeight3(num uint32) int {
	res := 0
	for num > 0 {
		fmt.Println(num & 1)
		if num&1 == 1 {
			res++
		}
		num = num >> 1
	}
	return res
}

func main() {
	var p uint32 = 00000000000000000000000000001011
	ret := hammingWeight(p)
	fmt.Println(ret)
}
