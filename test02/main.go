package main

import "fmt"

func main() {
	//var arr [6]int = [6]int{1, 2, 3, 4, 5, 6}
	//
	//for index, num := range arr {
	//	fmt.Printf("arr[%d]의 값은 %d입니다.\n", index, num)
	//}

	//var temp []aa
	temp := make([]aa, 1)

	//temp[1] = aa{}
	temp[0].a1 = "string"

	fmt.Println(temp)
}

type aa struct {
	a1 string
	a2 string
}
