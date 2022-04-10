package main

import "fmt"

type triangle struct {
	width, height float32
}

func triArea(s *triangle) float32 {
	return s.width * s.height / 2
}

func main() {
	tri1 := new(triangle)
	tri1.width = 12.5
	tri1.height = 5.2

	triarea := triArea(tri1)
	fmt.Printf("삼각형 너비:%.2f, 높이:%.2f 일때, 넓이:%.2f ", tri1.width, tri1.height, triarea)
}
