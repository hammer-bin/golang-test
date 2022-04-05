package main

import (
	"fmt"
	"net"
)

func main() {

	/*
		fmt.Println("octet 1: ")
		var input = "00000177.0.0.1/24"
		var addr, mask, err = net.ParseCIDR(input)
		fmt.Println("input: ", input, " result: ", addr, mask, err)
	*/

	fmt.Println("octet 2: ")
	var input1 = "192.168.0.1/24"
	var addr1, mask1, err1 = net.ParseCIDR(input1)
	fmt.Println("input: ", input1, " result: ", addr1, mask1, err1)
}
