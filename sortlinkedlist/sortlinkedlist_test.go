package sortlinkedlist

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		head *ListNode
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
					4,
					&ListNode{
						2,
						&ListNode{
							1, &ListNode{
								3,
								nil,
							},
						},
					},
				},
			},
			want: &ListNode{
				1,
				&ListNode{
					2,
					&ListNode{
						3, &ListNode{
							4,
							nil,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
