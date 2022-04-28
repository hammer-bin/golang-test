package main

import (
	"fmt"
)

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
		return
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
	for head != nil && head.Val == val {
		head = head.Next
	}

	node := head
	for node != nil && node.Next != nil {
		if node.Next.Val == val {
			node.Next = node.Next.Next
			continue
		}
		node = node.Next
		fmt.Println(node.Val)
	}
	return head
}

func removeElements2(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}
	if head == nil {
		return head
	}
	node := head
	for node != nil && node.Next != nil {
		if node.Next.Val == val {
			node.Next = node.Next.Next
			continue
		}
		node = node.Next

	}
	return head
}

func main() {
	var a []int
	a = []int{1, 1, 6, 3, 4, 5, 6}
	list := &LinkedList{}
	for _, value := range a {
		list.addNode(value)
	}

	list.printNode()
	removeElements(list.head, 1)
	list.printNode()

}
