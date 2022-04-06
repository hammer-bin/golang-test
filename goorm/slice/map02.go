package main

import "fmt"

func main() {
	var m = make(map[string]string)

	m["02"] = "서울특별시"
	m["031"] = "경기도"
	m["032"] = "강원도"
	m["041"] = "충청남도"

	fmt.Println(m)

	m["032"] = "인천"

	fmt.Println(m)

	delete(m, "031")

	fmt.Println(m)

}
