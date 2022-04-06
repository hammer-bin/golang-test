package main

import "fmt"

func main() {
	i := 0

TEST1:
	for {
		if i == 0 {
			break TEST1
		}
	}

	fmt.Println("End")
}
