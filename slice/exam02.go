package main

import "fmt"

func main() {
	var names []string

	var name string

	for {
		fmt.Scan(&name)
		if name == "1" {
			break
		}
		names = append(names, name)
	}

	var result string = names[0]

	for i := 0; i < len(names); i++ {
		if len(result) < len(names[i]) {
			result = names[i]
		}
	}

	fmt.Println(result, len(result))
}
