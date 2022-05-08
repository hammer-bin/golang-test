package main

import "fmt"

func sumOddLengthSubarrays(arr []int) int {
	sum := 0

	// size of subarray
	for i := 1; i <= len(arr); i += 2 {

		// iterate over all starting indices for subarrays
		for j := 0; j < len(arr)-i+1; j++ {

			// iterate over all elements in current subarray
			for k := j; k < j+i; k++ {
				sum += arr[k]
			}

		}

	}

	return sum
}

func sumOddLengthSubarrays2(arr []int) int {
	sum := 0
	l := len(arr)
	// Traverse the array
	for a := 0; a < l; a++ {
		// Generate all subarray of odd length
		for b := a; b < l; b += 2 {
			for c := a; c <= b; c++ {
				// Add the element to sum
				sum += arr[c]
			}
		}
	}
	return sum
}

func main() {
	input := []int{1, 4, 2, 5, 3}
	rst := sumOddLengthSubarrays(input)

	fmt.Println(rst)
}
