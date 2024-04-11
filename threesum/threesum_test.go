package threesum

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "a",
			args: args{
				nums: []int{
					-1, 0, 1, 2, -1, -4,
				},
			},
			want: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			name: "b",
			args: args{
				nums: []int{
					0, 1, 1,
				},
			},
			want: [][]int{},
		},
		{
			name: "c",
			args: args{
				nums: []int{
					0, 0, 0,
				},
			},
			want: [][]int{
				{0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
