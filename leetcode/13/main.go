package main

import "fmt"

const (
	I int = 1
	V int = 5
	X int = 10
	L int = 50
	C int = 100
	D int = 500
	M int = 1000
)

//I can be placed before V (5) and X (10) to make 4 and 9.
//X can be placed before L (50) and C (100) to make 40 and 90.
//C can be placed before D (500) and M (1000) to make 400 and 900.

func romanToInt(s string) int {
	runes := []rune(s)
	var sum int
	var i2, v, x, l, c, d, m int

	for i := 0; i < len(runes); i++ {
		//fmt.Printf("%c\n", runes[i])
		//fmt.Println(string(runes[i]))
		if string(runes[i]) == "I" {
			sum += I
			i2 += I
		} else if string(runes[i]) == "V" {
			if i != 0 && string(runes[i-1]) == "I" {
				sum += V - I*2
			} else {
				sum += V
				v += V
			}

		} else if string(runes[i]) == "X" {
			if i != 0 && string(runes[i-1]) == "I" {
				sum += X - I*2
			} else {
				sum += X
				x += X
			}

		} else if string(runes[i]) == "L" {
			if i != 0 && string(runes[i-1]) == "X" {
				sum += L - X*2
			} else {
				sum += L
				l += L
			}

		} else if string(runes[i]) == "C" {
			if i != 0 && string(runes[i-1]) == "X" {
				sum += C - X*2
			} else {
				sum += C
				c += C
			}

		} else if string(runes[i]) == "D" {
			if i != 0 && string(runes[i-1]) == "C" {
				sum += D - C*2
			} else {
				sum += D
				d += D
			}

		} else if string(runes[i]) == "M" {
			if i != 0 && string(runes[i-1]) == "C" {
				sum += M - C*2
			} else {
				sum += M
				m += M
			}

		} else {
			fmt.Println("wrong input")
			return 0
		}
	}
	fmt.Printf("I : %d\nV : %d\nX : %d\nL : %d\nC : %d\nD : %d\nM : %d\n", i2, v, x, l, c, d, m)
	return sum
}

func main() {
	var input string

	fmt.Scanf("%s", &input)

	sum := romanToInt(input)

	fmt.Println(sum)
}
