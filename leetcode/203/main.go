package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 * Input: head = [1,2,6,3,4,5,6], val = 6
 * Output: [1,2,3,4,5]
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	head *ListNode
	tail *ListNode
}

func (link *LinkedList) addNode(num int) {
	if link.head == nil {
		link.head = &ListNode{Val: num}
		link.tail = link.head
	}
	link.tail.Next = &ListNode{Val: num}
	link.tail = link.tail.Next

}

func (link *LinkedList) printNode() {
	node := link.head
	for node.Next != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Next
	}
	fmt.Println(node.Val)

}

func removeElements(head *ListNode, val int) *ListNode {
	prev := head
	for prev.Next != nil {

	}
	return nil
}

func main() {
	var a []int
	a = []int{1, 2, 6, 3, 4, 5, 6}
	list := &LinkedList{}
	for _, value := range a {
		list.addNode(value)
	}

	list.printNode()
	removeElements(list.head, 6)

}
