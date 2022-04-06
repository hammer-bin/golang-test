package main

import "fmt"

func main() {
	var m = make(map[string]string)

	m["02"] = "서울특별시"
	m["031"] = "경기도"
	m["032"] = "인천"
	m["053"] = "대구광역시"

	fmt.Println(m["032"])
	fmt.Println(m["042"], "빈 칸입니다.")

	val, exist := m["02"]
	fmt.Println(val, exist)

	val, exist = m["042"]
	fmt.Println(val, exist)

	val = m["053"]
	fmt.Println(val)

	_, exist = m["053"]
	fmt.Println(exist)

	fmt.Println(len(m))
}
