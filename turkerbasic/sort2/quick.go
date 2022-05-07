package main

import "fmt"

func main() {
	//arr := []int{10, 30, 27, 5, 6, 4, 1, 90, 60, 3}
	arr := []int{27, 5, 6, 4, 10}
	sortedArr := QuickSort(arr)

	fmt.Println(sortedArr)
}

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	p := len(arr) - 1
	i := 0
	j := p - 1
	for i <= j {
		if arr[i] < arr[p] {
			i++
		} else {
			if arr[j] >= arr[p] {
				j--
			} else {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}
	}
	arr[p], arr[i] = arr[i], arr[p]
	QuickSort(arr[:i])
	QuickSort(arr[i+1:])

	return arr
}
