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

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "a",
			args: args{
				&TreeNode{
					3,
					&TreeNode{
						5,
						&TreeNode{
							Val: 6,
						},
						&TreeNode{
							2,
							&TreeNode{Val: 7},
							&TreeNode{Val: 4},
						},
					},
					&TreeNode{
						1,
						&TreeNode{Val: 0},
						&TreeNode{Val: 8},
					},
				},
			},
			want: &TreeNode{
				2,
				&TreeNode{Val: 4},
				&TreeNode{Val: 7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.root); !reflect.DeepEqual(got.Val, tt.want.Val) {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
