package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("정수 a와 b를 입력하시오:")
	fmt.Scanln(&a, &b)

	switch {
	case a > b:
		fmt.Println("a가 b보다 큽니다.")
	case a < b:
		fmt.Println("a가 b보다 작습니다.")
	case a == b:
		fmt.Println("a와 b가 같습니다.")
	default:
		fmt.Println("모르겠어요.")
	}
}
