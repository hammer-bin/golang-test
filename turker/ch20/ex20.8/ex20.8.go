package main

type Stringer interface {
	String() string
}

type Student struct {
}

func main() {
	var stringer Stringer
	stringer.(Student) // 컴파일 타임에러 발생
}
