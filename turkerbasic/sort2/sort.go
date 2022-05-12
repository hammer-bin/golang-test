package main

import "fmt"

func main() {
	arr := []int{10, 30, 27, 5, 60, 4, 1, 90, 60, 3}
	//sortedArr := MergeSort2(arr)
	sortedArr := QuickSort2(arr)

	fmt.Println(sortedArr)
}

func MergeSort2(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	left := MergeSort2(arr[:len(arr)/2])
	right := MergeSort2(arr[len(arr)/2:])

	i, j := 0, 0
	rst := make([]int, 0)
	for i < len(left) || j < len(right) {
		if i >= len(left) {
			rst = append(rst, right[j])
			j++
		} else if j >= len(right) {
			rst = append(rst, left[i])
			i++
		} else {
			if left[i] < right[j] {
				rst = append(rst, left[i])
				i++
			} else {
				rst = append(rst, right[j])
				j++
			}
		}
	}
	return rst
}

func QuickSort2(arr []int) []int {
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
	QuickSort2(arr[:i])
	QuickSort2(arr[i+1:])
	return arr
}
