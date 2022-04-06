package main

import "fmt"

func main() {
	var a = make(map[string]int)
	var subject string
	var point int
	var avg float32
	var totalpoint int

	for {
		fmt.Scanf("%s %d", &subject, &point)
		if subject == "0" {
			break
		}
		a[subject] = point
	}

	for subject, point := range a {
		fmt.Printf("%s %d\n", subject, point)
		totalpoint = totalpoint + point
	}

	avg = float32(totalpoint) / float32(len(a))
	fmt.Printf("%.2f", avg)
}
