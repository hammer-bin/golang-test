package main

import "fmt"

func main() {
	var a map[int]string

	if a == nil {
		fmt.Println("nil map")
	}

	var m = map[string]string{
		"apple":  "red",
		"grape":  "purple",
		"banana": "yellow",
	}

	fmt.Println(m, "\nm의 길이는", len(m))
}
