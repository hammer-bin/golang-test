package main

import "fmt"

func main() {
	var i, j, a int
	//i,j는 두 개의 반복문에 쓰일 변수

	fmt.Scanf("%d", &a)
	i = 0
	for i < a {
		j = 0
		for j < i {
			fmt.Print("o ")
			j += 1
		}
		fmt.Println("* ")
		i += 1
	}
}
