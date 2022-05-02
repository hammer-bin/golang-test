package main

import (
	"fmt"
	"turkerbasic/dataStruct"
)

func queuemain() {
	// slice로 stack 구현
	stack := []int{}
	for i := 1; i <= 5; i++ {
		stack = append(stack, i)
	}

	fmt.Println(stack)

	for len(stack) > 0 {
		var last int
		last, stack = stack[len(stack)-1], stack[:len(stack)-1]
		fmt.Println(last)
	}

	// slice로 queue 구현
	queue := []int{}
	for i := 1; i <= 5; i++ {
		queue = append(queue, i)
	}

	fmt.Println(queue)

	for len(queue) > 0 {
		var front int
		front, queue = queue[0], queue[1:]
		fmt.Println(front)
	}

	//linkedlist로 stack 구현
	stack2 := dataStruct.NewStack()

	for i := 1; i <= 5; i++ {
		stack2.Push(i)
	}

	fmt.Println("NewStack")

	for !stack2.Empty() {
		val := stack2.Pop()
		fmt.Printf("%d ->", val)
	}

	queue2 := dataStruct.NewQueue()

	for i := 1; i <= 5; i++ {
		queue2.Push(i)
	}

	fmt.Println("\nNewQueue")

	for !queue2.Empty() {
		val := queue2.Pop()
		fmt.Printf("%d ->", val)
	}
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	fmt.Println(s)
}

func Amain() {
	tree := dataStruct.Tree{}

	val := 1
	tree.AddNode(val)
	val++

	// 자식노드 3개 추가
	for i := 0; i < 3; i++ {
		tree.Root.AddNode(val)
		val++
	}

	//자식노드 2개씩 추가
	for i := 0; i < len(tree.Root.Childs); i++ {
		for j := 0; j < 2; j++ {
			tree.Root.Childs[i].AddNode(val)
			val++
		}
	}

	tree.DFS1()
	fmt.Println()

	tree.DFS2()
	fmt.Println()

	tree.BFS()

	fmt.Println()

	tree1 := dataStruct.NewBinaryTree(5)
	tree1.Root.AddNode(3)
	tree1.Root.AddNode(2)
	tree1.Root.AddNode(4)
	tree1.Root.AddNode(8)
	tree1.Root.AddNode(7)
	tree1.Root.AddNode(6)
	tree1.Root.AddNode(10)
	tree1.Root.AddNode(9)

	tree1.Print()
	fmt.Println()

	if found, cnt := tree1.Search(6); found {
		fmt.Println("found 6 cnt", cnt)
	} else {
		fmt.Println("not found 6 cnt ", cnt)
	}

	fmt.Println()

	if found, cnt := tree1.Search(11); found {
		fmt.Println("found 11 cnt", cnt)
	} else {
		fmt.Println("not found 11 cnt ", cnt)
	}

	var s = []int{1, 2, 3, 4, 5}
	reverse(s)

}

func main() {
	h := &dataStruct.Heap{}

	h.Push(9)
	h.Push(8)
	h.Push(7)
	h.Push(6)
	h.Push(5)

	h.Print()
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
}
