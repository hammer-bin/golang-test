package main

import "fmt"

type person struct {
	name    string
	age     int
	contact string
}

func main() {
	var p1 = person{}
	fmt.Println(p1)

	p1.name = "kim"
	p1.age = 25
	p1.contact = "01000002423"
	fmt.Println(p1)

	p2 := person{"nam", 31, "01048292223"}
	fmt.Println(p2)

	p3 := person{contact: "01023202984", name: "park", age: 23}
	fmt.Println(p3)

	p3.name = "ryu"
	fmt.Println(p3)

	fmt.Println(p3.contact)
}
