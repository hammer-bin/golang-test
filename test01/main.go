package main

import "fmt"

func main() {
	var num1, num2, num3, result int

	fmt.Scanln(&num1, &num2, &num3)

	result = num1*num2 + num3

	fmt.Printf("num1의 값은:%d num2의 값은:%d\n", num1, num2)

	fmt.Printf("%d x %d + %d = %d", num1, num2, num3, result)
}
