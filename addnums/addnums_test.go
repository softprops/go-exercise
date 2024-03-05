package addnums

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		a *ListNode
		b *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "a",
			args: args{
				// 243
				a: NewListNode(2, NewListNode(4, NewListNode(3, nil))),
				//564
				b: NewListNode(5, NewListNode(6, NewListNode(4, nil))),
			},
			//708
			want: NewListNode(7, NewListNode(0, NewListNode(8, nil))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution() = %s, want %s", got, tt.want)
			}
		})
	}
}
