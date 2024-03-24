// Lowest Common Ancestor of Deepest Leaves
//
// Given the root of a binary tree, return the lowest common ancestor of its deepest leaves.
//
// Recall that:
//
// The node of a binary tree is a leaf if and only if it has no children
// The depth of the root of the tree is 0. if the depth of a node is d, the depth of each of its children is d + 1.
// The lowest common ancestor of a set S of nodes, is the node A with the largest depth such that every node in S is in the subtree with root A.
//
// https://leetcode.com/problems/lowest-common-ancestor-of-deepest-leaves
package deepleavelca

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// stringer so we can see our tree
func (n *TreeNode) String() string {
	return fmt.Sprintf("TreeNode{%v,%v,%v}", n.Val, n.Left, n.Right)
}

func Solution(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	left, right := depth(root.Left), depth(root.Right)
	// we found our leaves
	if left == right {
		return root
	}
	// look left
	if left > right {
		return Solution(root.Left)
	}
	// look right
	return Solution(root.Right)
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(depth(root.Left), depth(root.Right))
}
