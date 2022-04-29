package main

import (
	"fmt"
)

/**
876. Middle of the Linked List
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
Input: head = [1,2,3,4,5]
Output: [3,4,5]
Explanation: The middle node of the list is node 3.

Input: head = [1,2,3,4,5,6]
Output: [4,5,6]
Explanation: Since the list has two middle nodes with values 3 and 4, we return the second one.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	head *ListNode
	tail *ListNode
}

func (l *LinkedList) addNode(val int) {

	if l.head == nil {
		l.head = &ListNode{Val: val}
		l.tail = l.head
		return
	}
	l.tail.Next = &ListNode{Val: val}
	l.tail = l.tail.Next
}

func (l *LinkedList) printNode() {
	node := l.head
	for node.Next != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Next
	}
	fmt.Printf("%d -> ", node.Val)
	fmt.Println()

}

func middleNode(head *ListNode) *ListNode {
	cnt := 0
	node := head
	for node.Next != nil {
		cnt++
		node = node.Next
	}
	cnt++
	fmt.Println("cnt : ", cnt)

	node1 := head
	for i := 0; i < cnt/2; i++ {
		node1 = node1.Next
	}
	rstNode := node1

	return rstNode
}

func main() {
	sl := []int{1, 2, 3, 4, 5, 6}
	list := &LinkedList{}
	for _, v := range sl {
		list.addNode(v)
	}
	list.printNode()

	node := middleNode(list.head)

	fmt.Println()

	printNode(node)

}

func printNode(node *ListNode) {
	for node.Next != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Next
	}
	fmt.Printf("%d -> ", node.Val)
}
