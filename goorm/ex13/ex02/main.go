package main

import "fmt"

type person struct {
	name    string
	age     int
	contact string
}

func addAgeRef(a *person) {
	a.age += 4
}

func addAgeVal(a person) {
	a.age += 4
}

func main() {
	var p1 = new(person) //포인터 구조체 객체 생성
	var p2 = person{}    // 빈 구조체 객체 생성

	fmt.Println(p1, p2)

	p1.age = 25
	p2.age = 25

	fmt.Println(p1, p2)

	addAgeRef(p1) //& 생략 가능
	addAgeVal(p2)

	fmt.Println(p1, p2)

	addAgeRef(&p2) //&을 붙이면 pass by reference 가능

	fmt.Println(p1, p2)
}
