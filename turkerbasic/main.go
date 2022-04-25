package main

import (
	"fmt"
	"turkerbasic/dataStruct"
)

func main() {
	tree := dataStruct.Tree{}

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

}
