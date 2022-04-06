package main

import "fmt"

func exampleFunc1() {
	var a int = 10

	a++

	fmt.Println("exampleFunc1의 a는 ", a)
}

func exampleFunc2() {
	var b int = 20
	var c int = 30

	b++
	c++

	fmt.Println("b와 c는 ", b, c)
}

func main() {
	var a int = 28

	exampleFunc1()
	exampleFunc2()
	fmt.Println("main의 a는", a)
}
