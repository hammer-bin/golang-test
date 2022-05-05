package main

import (
	"fmt"
	"turkerbasic/repeat1/dataStruct1"
)

func mainLink() {
	list := &dataStruct1.LinkedList{}
	list.AddNode(0)

	for i := 1; i < 10; i++ {
		list.AddNode(i)
	}

	list.PrintNode()

	list.RemoveNode(list.Root.Next)

	list.PrintNode()

	list.RemoveNode(list.Root)

	list.PrintNode()

	list.RemoveNode(list.Tail)

	list.PrintNode()

	fmt.Printf("tail:%d\n", list.Tail.Val)

	list.PrintReverse()

}

func main() {
	var stack []int

	for i := 1; i <= 5; i++ {
		stack = append(stack, i)
	}

	fmt.Println(stack)

	for len(stack) > 0 {
		var last int
		last, stack = stack[len(stack)-1], stack[:len(stack)-1]
		fmt.Println(last)
	}

	var queue []int

	for i := 1; i <= 5; i++ {
		queue = append(queue, i)
	}

	fmt.Println(queue)

	for len(queue) > 0 {
		var front int
		front, queue = queue[0], queue[1:]
		fmt.Println(front)
	}

	stack2 := dataStruct1.NewStack()

	for i := 1; i <= 5; i++ {
		stack2.Push(i)
	}

	fmt.Println("NewStack")

	for !stack2.Empty() {
		val := stack2.Pop()
		fmt.Printf("%d ->", val)
	}

	fmt.Println("\nNewQueue")
	queue2 := dataStruct1.NewQueue()

	for i := 1; i <= 5; i++ {
		queue2.Push(i)
	}

	for !queue2.Empty() {
		val := queue2.Pop()
		fmt.Printf("%d ->", val)
	}

}
