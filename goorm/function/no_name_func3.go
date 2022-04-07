package main

import "fmt"

func add() {
	fmt.Println("선언 함수를 호출했습니다.")
}

func main() {
	add := func() {
		fmt.Println("익명 함수를 호출했습니다.")
	}

	add()
}
