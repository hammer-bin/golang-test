package main

import (
	"fmt"
)

/**
 * 206. Reverse Linked List
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
Input: head = [1,2]
Output: [2,1]
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
	for node != nil && node.Next != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Next
	}
	fmt.Printf("%d -> ", node.Val)
}

func reverseList(head *ListNode) *ListNode {
	prev := &ListNode{}
	for head != nil {
		temp := head.Next // temp 234  34  4   nil
		head.Next = prev  // cur  1    21  321 4321
		prev = head       // prev 1    21  321 4321
		head = temp       // cur  234  34  4   nil
	}
	return prev
}

func reverseList2(head *ListNode) *ListNode {
	var front *ListNode
	for head != nil {
		temp := head.Next
		head.Next = front
		front, head = head, temp
	}
	return front
}

func main() {
	var slice []int
	slice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	list := &LinkedList{}

	for _, val := range slice {
		list.addNode(val)
	}

	list.printNode()
	reverseList2(list.head)

	fmt.Println()
	list.printNode()
}
