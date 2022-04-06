package main

import "fmt"

func main() {
	for i := 2; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		for j := 1; j < 10; j++ {
			if i < j {
				break
			}
			fmt.Printf("%d x %d = %d\n", i, j, i*j)
		}
		fmt.Println()
	}
}
