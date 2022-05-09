package main

import (
	"fmt"
	"turkerbasic/repeat2/dataStruct"
)

func main() {
	h := dataStruct.Heap{}

	h.Push(2)
	h.Push(9)
	h.Push(11)
	h.Push(5)
	h.Push(14)
	h.Push(8)
	h.Push(9)
	h.Push(6)

	h.Print()

	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())

}
