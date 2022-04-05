package main

import "fmt"

func main() {
	var arr1 [5]int
	fmt.Println(arr1)

	arr1 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1, arr1[0], arr1[4])

	arr2 := [4]int{4, 5, 6, 7}
	arr2[0] = 32
	fmt.Println(arr2)

	var arr3 = [...]int{9, 8, 7, 6}
	fmt.Println(arr3, len(arr3))

	var a = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(a[1][2])
}
