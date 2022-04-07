package main

import "fmt"

var a int = 1

func localVar() int {
	var a int = 10

	a += 3

	return a
}

func globalVar() int {
	a += 3

	return a
}

func main() {
	fmt.Println("지역변수 a의 연산: ", localVar())
	fmt.Println("전역변수 a의 연산: ", globalVar())
}
