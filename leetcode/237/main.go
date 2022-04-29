package main

import "fmt"

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

func (l *ListNode) printNode() {
	for l.Next != nil {
		fmt.Printf("%d -> ", l.Val)
	}
	fmt.Printf("%d -> ", l.Val)
}

func (l *LinkedList) printNode() {
	node := l.head
	for node.Next != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Next
	}
	fmt.Printf("%d -> ", node.Val)
}

func (l *LinkedList) printNode2() {
	for l.head.Next != nil {
		fmt.Printf("%d -> ", l.head.Val)
		l.head = l.head.Next
	}
	fmt.Printf("%d -> ", l.head.Val)
}

func deleteNode(node *ListNode) {
	ref := node
	ref.Val = ref.Next.Val
	ref.Next = ref.Next.Next
}

func main() {
	list := &LinkedList{}
	sl := []int{4, 5, 1, 9}
	for _, v := range sl {
		list.addNode(v)
	}
	list.printNode()
	fmt.Println()
	list.printNode2()

	deleteNode(list.head)
}
