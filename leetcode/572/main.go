package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// s의 모든 노드를 돌면서 compareTree 를 호출해서 true가 반환되면 true를 반환한다.
func isSubtree(s *TreeNode, t *TreeNode) bool {
	return DFSFunc(s, t, compareTree)
}

func DFSFunc(s, t *TreeNode, f func(*TreeNode, *TreeNode) bool) bool {
	if s == nil {
		if t == nil {
			return true
		}
		return false
	}
	if f(s, t) == true {
		return true
	}
	if DFSFunc(s.Left, t, f) == true {
		return true
	}
	return DFSFunc(s.Right, t, f)
}

func compareTree(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil {
		if t2 == nil {
			return true
		}
		return false
	} else if t2 == nil {
		return false
	}
	if t1.Val != t2.Val {
		return false
	}
	if !compareTree(t1.Left, t2.Left) {
		return false
	}
	return compareTree(t1.Right, t2.Right)
}
