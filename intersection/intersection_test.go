// Intersection of Two Arrays
//
// https://leetcode.com/problems/intersection-of-two-arrays/
package intersection

import (
	"reflect"
	"slices"
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "a",
			args: args{
				nums1: []int{1, 2, 2, 1},
				nums2: []int{2, 2},
			},
			want: []int{2},
		},
		{
			name: "b",
			args: args{
				nums1: []int{4, 9, 5},
				nums2: []int{9, 4, 9, 8, 4},
			},
			want: []int{9, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(sorted(got), sorted(tt.want)) {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func sorted(nums []int) []int {
	slices.Sort(nums)
	return nums
}
