package kthlargest

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "a",
			args: args{
				//            👇
				// 1, 2, 3, 4, 5, 6
				nums: []int{3, 2, 1, 4, 6, 4},
				k:    2,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
