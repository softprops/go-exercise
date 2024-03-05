package validbinarytree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Solution(root *TreeNode) bool {
	return valid(root, math.MinInt64, math.MaxInt64)
}

func valid(node *TreeNode, min int, max int) bool {
	if node == nil {
		return true
	}
	if node.Val <= min || node.Val >= max {
		return false
	}
	return valid(node.Left, node.Val, max) && valid(node.Right, min, node.Val)
}
