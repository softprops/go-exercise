// Sum Root to Leaf Numbers
//
// https://leetcode.com/problems/sum-root-to-leaf-numbers/description/
//
// #tree #dfs #binarytree
package treesum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Solution(root *TreeNode) int {
	return recurse(root, 0)
}

func recurse(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	sum = (sum * 10) + root.Val
	// when we've reached the end, return the sum
	if root.Left == nil && root.Right == nil {
		return sum
	}

	return recurse(root.Left, sum) + recurse(root.Right, sum)
}
