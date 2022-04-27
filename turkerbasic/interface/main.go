package main

import (
	"fmt"
	"strconv"
)

type structA struct {
}

func (a *structA) AAA(x int) int {
	return x * x
}

func (a *structA) BBB(x int) string {
	return "X=" + strconv.Itoa(x)
}

type StructB struct {
}

func (b *StructB) AAA(x int) int {
	return x * 2
}

func main() {
	var c InterfaceA
	c = &structA{}

	//var d InterfaceA
	fmt.Println(c.BBB(3))

}
