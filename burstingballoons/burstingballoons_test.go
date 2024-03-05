// Minimum Number of Arrows to Burst Balloons
//
// https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/description
//
// #intervals
package burstingballoons

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "a",
			args: args{
				points: [][]int{
					{10, 16}, {2, 8}, {1, 6}, {7, 12},
				},
			},
			want: 2,
		},
		{
			name: "b",
			args: args{
				points: [][]int{
					{1, 2}, {3, 4}, {5, 6}, {7, 8},
				},
			},
			want: 4,
		},
		{
			name: "c",
			args: args{
				points: [][]int{
					{1, 2}, {2, 3}, {3, 4}, {4, 5},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.points); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
