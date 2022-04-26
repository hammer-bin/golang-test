package dataStruct

import "fmt"

type Node struct {
	Next *Node
	Prev *Node
	Val  int
}

type LinkedList struct {
	Root *Node
	Tail *Node
}

func (l *LinkedList) AddNode(val int) {
	if l.Root == nil {
		l.Root = &Node{Val: val}
		l.Tail = l.Root
		return
	}
	l.Tail.Next = &Node{Val: val}
	prev := l.Tail
	l.Tail = l.Tail.Next
	l.Tail.Prev = prev

}

func (l *LinkedList) RemoveNode(node *Node) {
	if node == l.Root {
		l.Root = l.Root.Next
		l.Root.Prev = nil
		node.Next = nil
		return
	}
	Prev := node.Prev

	if node == l.Tail {
		Prev.Next = nil
		l.Tail.Prev = nil
		l.Tail = Prev
	} else {
		node.Prev = nil
		Prev.Next = Prev.Next.Next
		Prev.Next.Prev = Prev
	}
	node.Next = nil
}

func (l *LinkedList) PrintNodes() {
	node := l.Root
	for node.Next != nil {
		fmt.Printf("%d ->", node.Val)
		node = node.Next
	}
	fmt.Printf("%d\n", node.Val)
}

func (l *LinkedList) PrintReverse() {
	node := l.Tail
	for node.Prev != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Prev
	}
	fmt.Printf("%d\n", node.Val)
}
