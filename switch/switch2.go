package main

import "fmt"

func main() {
	var num int
	var result string

	fmt.Print("10, 20, 30중 하나를 입력하시오:")
	fmt.Scanln(&num)

	switch num / 10 { //표현식
	case 1:
		result = "A"
	case 2:
		result = "B"
	case 3:
		result = "C"
	default:
		fmt.Println("모르겠어요.")
		return
	}

	fmt.Println(result)
}
