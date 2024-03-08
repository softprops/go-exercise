package topkfreq

import (
	"reflect"
	"slices"
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "a",
			args: args{
				nums: []int{
					1, 1, 1, 2, 2, 3,
				},
				k: 2,
			},
			want: []int{1, 2},
		},
		{
			name: "a",
			args: args{
				nums: []int{
					1,
				},
				k: 1,
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.nums, tt.args.k); !reflect.DeepEqual(sorted(got), sorted(tt.want)) {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func sorted(nums []int) []int {
	slices.Sort(nums)
	return nums

}
