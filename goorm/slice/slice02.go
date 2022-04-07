package main

import "fmt"

func main() {
	c := make([]int, 0, 3) //용량이 3이고 길이가 0인 정수형 슬라이스 선언
	c = append(c, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(c)
	fmt.Println(len(c), cap(c))

	l := c[1:3]
	fmt.Println(l)

	l = c[2:]
	fmt.Println(l)

	l[0] = 6

	fmt.Println(c)
	fmt.Println(l)
}
