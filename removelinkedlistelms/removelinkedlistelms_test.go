// Remove Linked List Elements
//
// Given the head of a linked list and an integer val, remove all the nodes of the linked list that has Node.val == val, and return the new head.
//
// https://leetcode.com/problems/remove-linked-list-elements/
package removelinkedlistelms

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		head *ListNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "a",
			args: args{
				head: &ListNode{
					1,
					&ListNode{
						2,
						&ListNode{
							6,
							&ListNode{
								3,
								&ListNode{
									4,
									&ListNode{
										5,
										&ListNode{
											Val: 6,
										},
									},
								},
							},
						},
					},
				},
				val: 6,
			},
			want: &ListNode{
				1,
				&ListNode{
					2,
					&ListNode{
						3,
						&ListNode{
							4,
							&ListNode{
								Val: 5,
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.head, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
