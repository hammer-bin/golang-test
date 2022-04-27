package main

import (
	"fmt"
	"time"
)

func pop(c chan int) {
	fmt.Println("pop func")
	v := <-c
	fmt.Println(v)
}

func main() {
	var c chan int
	c = make(chan int)

	go pop(c)
	time.Sleep(1)
	c <- 10

	fmt.Println("end of program")
}
