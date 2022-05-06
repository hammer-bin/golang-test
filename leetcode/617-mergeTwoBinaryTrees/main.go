package main

//// missing

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Tree struct {
	root *TreeNode
}

func (t *Tree) addNode(val int) {
	if t.root == nil {
		t.root = &TreeNode{Val: val}
		return
	}

}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	} else if root2 == nil {
		return root1
	} else {
		root1.Val += root2.Val
		root1.Left = mergeTrees(root1.Left, root2.Left)
		root1.Right = mergeTrees(root1.Right, root2.Right)
	}

	return root1
}

func main() {

}
