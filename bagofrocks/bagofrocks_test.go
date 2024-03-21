// Maximum Bags With Full Capacity of Rocks
//
// You have n bags numbered from 0 to n - 1. You are given two 0-indexed integer arrays capacity and rocks. The ith bag can hold a maximum of capacity[i] rocks and currently contains rocks[i] rocks. You are also given an integer additionalRocks, the number of additional rocks you can place in any of the bags.
//
// Return the maximum number of bags that could have full capacity after placing the additional rocks in some bags.
//
// https://leetcode.com/problems/maximum-bags-with-full-capacity-of-rocks/description/
package bagofrocks

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		capacity        []int
		rocks           []int
		additionalRocks int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "a",
			args: args{
				capacity: []int{
					2, 3, 4, 5,
				},
				rocks: []int{
					1, 2, 4, 4,
				},
				additionalRocks: 2,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.capacity, tt.args.rocks, tt.args.additionalRocks); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
