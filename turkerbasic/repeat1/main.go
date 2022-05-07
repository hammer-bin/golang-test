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
	/*var stack []int

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

	fmt.Println()

	//tree
	tree := dataStruct1.Tree{}

	val := 1
	tree.AddNode(val)
	val++

	for i := 0; i < 3; i++ {
		tree.Root.AddNode(val)
		val++
	}

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
	*/

	/*tree := dataStruct1.NewBinaryTree(5)
	tree.Root.AddNode(3)
	tree.Root.AddNode(2)
	tree.Root.AddNode(4)
	tree.Root.AddNode(8)
	tree.Root.AddNode(7)
	tree.Root.AddNode(6)
	tree.Root.AddNode(10)
	tree.Root.AddNode(9)

	tree.Print()
	fmt.Println()
	if found, cnt := tree.Search(6); found {
		fmt.Println("found 6 cnt: ", cnt)
	} else {
		fmt.Println("not found 6 cnt: ", cnt)
	}

	if found, cnt := tree.Search(11); found {
		fmt.Println("found 11 cnt: ", cnt)
	} else {
		fmt.Println("not found 11 cnt: ", cnt)
	}*/

	//h := dataStruct1.Heap{}

	/*h.Push(2)
	h.Push(6)
	h.Push(9)
	h.Push(6)
	h.Push(7)
	h.Push(8)
	h.Print()

	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())*/

	/*//[-1,3,-1,5,4], 2번째
	nums := []int{-1, 3, -1, 5, 4}
	for i := 0; i < len(nums); i++ {
		h.Push(nums[i])
		if h.Count() > 2 {
			h.Pop()
		}
	}
	fmt.Println(h.Pop())

	nums = []int{2, 4, -2, -3, 8}
	for i := 0; i < len(nums); i++ {
		h.Push(nums[i])
		if h.Count() > 1 {
			h.Pop()
		}
	}
	fmt.Println(h.Pop())

	nums = []int{-5, -3, 1}
	for i := 0; i < len(nums); i++ {
		h.Push(nums[i])
		if h.Count() > 3 {
			h.Pop()
		}
	}
	fmt.Println(h.Pop())*/

	fmt.Println("abcde = ", dataStruct1.Hash("abcde"))
	fmt.Println("abcde = ", dataStruct1.Hash("abcde"))
	fmt.Println("abcde = ", dataStruct1.Hash("abcdf"))
	fmt.Println("abcde = ", dataStruct1.Hash("6bcdf"))

	m := dataStruct1.CreateMap()
	m.Add("AAA", "0107777777")
	m.Add("BBB", "0107555577")
	m.Add("CCC", "0102424577")
	m.Add("CDG", "0702423982")

	fmt.Println("AAA = ", m.Get("AAA"))
	fmt.Println("CCC = ", m.Get("CCC"))
	fmt.Println("DDD = ", m.Get("DDD"))
}
