package main

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

type NodeWidthDepth struct {
	node  *TreeNode
	depth int
}

//queue를 사용 BFS
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := make([]NodeWidthDepth, 0, 1024)
	q = append(q, NodeWidthDepth{root, 1})

	maxDepth := 0
	for len(q) > 0 {
		nd := q[0]
		if nd.depth > maxDepth {
			maxDepth = nd.depth
		}
		if nd.node.Left != nil {
			q = append(q, NodeWidthDepth{nd.node.Left, nd.depth + 1})
		}
		if nd.node.Right != nil {
			q = append(q, NodeWidthDepth{nd.node.Right, nd.depth + 1})
		}
		q = q[1:]
	}
	return maxDepth
}

var md int

//재귀호출방식 DFS
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	md = 1

	if root.Left != nil {
		dfs(root.Left, 2)
	}
	if root.Right != nil {
		dfs(root.Right, 2)
	}
	return md
}

func dfs(node *TreeNode, depth int) {
	if depth > md {
		md = depth
	}
	if node.Left != nil {
		dfs(node.Left, depth+1)
	}
	if node.Right != nil {
		dfs(node.Right, depth+1)
	}
}

//기본
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}
