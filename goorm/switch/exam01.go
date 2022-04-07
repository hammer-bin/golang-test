package main

import "fmt"

func main() {
	var sel int
	var num1, num2, result float64

	fmt.Scan(&sel)
	fmt.Scanln(&num1, &num2)

	switch sel {
	case 1:
		result = num1 + num2
	case 2:
		result = num1 - num2
	case 3:
		result = num1 * num2
	case 4:
		result = num1 / num2
	default:
		fmt.Print("잘못된 입력입니다.")
		return
	}

	if sel > 0 && sel < 5 {
		fmt.Printf("%.1f\n", result)
	}
}
