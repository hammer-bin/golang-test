package main

import (
	"fmt"
	"sort"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// LinkedList를 사용할 구조체 선언
type LinkedList struct {
	root *ListNode
	tail *ListNode
}

// node 삽입 펑션
func (l *LinkedList) AddNode(val int) {
	if l.root == nil {
		l.root = &ListNode{Val: val}
		l.tail = l.root
		return
	}
	l.tail.Next = &ListNode{Val: val}
	l.tail = l.tail.Next
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var a []int
	if list1 != nil {
		root := list1
		for root.Next != nil {
			a = append(a, root.Val)
			root = root.Next
		}
		a = append(a, root.Val)
	}

	if list2 != nil {
		root := list2
		for root.Next != nil {
			a = append(a, root.Val)
			root = root.Next
		}
		a = append(a, root.Val)
	}

	sort.Ints(a)

	rst := &LinkedList{}
	for _, number := range a {
		rst.AddNode(number)
		fmt.Println(number)
	}

	return rst.root
}

func main() {
	list1 := &LinkedList{}
	list1.AddNode(1)
	list1.AddNode(2)
	list1.AddNode(4)
	list2 := &LinkedList{}
	list2.AddNode(1)
	list2.AddNode(3)
	list2.AddNode(4)
	mergeTwoLists(list1.root, list2.root)
}
