package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {

	node := head
	var s []int
	for {

		fmt.Printf("%d - > ", node.Val)
		s = append(s, node.Val)
		node = node.Next
		b

	}
	return false
}

func main() {

}
