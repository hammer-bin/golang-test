package main

import "fmt"

func main() {
	var a []int        //슬라이스 변수 선언 아무것도 초기화 되지 않은 상태
	a = []int{1, 2, 3} //슬라이스에 리터럴값 지정

	a[1] = 10 //값이 할당되어 메모리가 생겼기 때문에 이렇게 접근 가능

	fmt.Println(a)
	fmt.Println("용량이", cap(a), "길이가", len(a), " Nil Slice입니다.")

	var b []int //nil slice 선언

	if b == nil {
		fmt.Println("용량이", cap(b), "길이가", len(b), " Nil Slice입니다.")
	}

	fmt.Println("------------------------------------------------------------")

	s := make([]int, 0, 3) // len=0, cap=3 인 슬라이스 선언

	for i := 1; i <= 10; i++ { // 1부터 차례대로 한 요소씩 추가
		s = append(s, i)

		fmt.Println(len(s), cap(s)) // 슬라이스 길이와 용량 확인
	}

	fmt.Println(s) //최종 슬라이스 출력

	fmt.Println("------------------------------------------------------------")

	sliceA := []int{1, 2, 3}
	sliceB := []int{4, 5, 6}

	sliceA = append(sliceA, sliceB...)
	//sliceA = append(sliceA, 4, 5, 6)

	fmt.Println(sliceA)

	fmt.Println("------------------------------------------------------------")

	sliceC := []int{0, 1, 2}
	sliceD := make([]int, len(sliceC), cap(sliceC)*2) //sliceA에 2배 용량인 슬라이스 선언

	copy(sliceD, sliceC)

	fmt.Println(sliceD)
	println(len(sliceD), cap(sliceD))
}
