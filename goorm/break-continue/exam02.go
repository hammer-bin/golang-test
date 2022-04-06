// 어려움.......
package main

import "fmt"

func main() {
	var result int

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i == j {
				continue
			}

			result = (i*10 + j) + (j*10 + i)
			if result != 99 {
				continue
			}
			fmt.Printf("%d%d + %d%d = %d\n", i, j, j, i, result)
		}
	}
}
