// Can Place Flowers
//
// You have a long flowerbed in which some of the plots are planted, and some are not. However, flowers cannot be planted in adjacent plots.
//
// Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1 means not empty, and an integer n, return true if n new flowers can be planted in the flowerbed without violating the no-adjacent-flowers rule and false otherwise.
//
// https://leetcode.com/problems/can-place-flowers/descriptio
package canplaceflowers

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		flowerbed []int
		n         int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "a",
			args: args{
				flowerbed: []int{1, 0, 0, 0, 1},
				n:         1,
			},
			want: true,
		},
		{
			name: "b",
			args: args{
				flowerbed: []int{1, 0, 0, 0, 1},
				n:         2,
			},
			want: false,
		},
		{
			name: "c",
			args: args{
				flowerbed: []int{1, 0, 0, 0, 0, 1},
				n:         2,
			},
			want: false,
		},
		{
			name: "d",
			args: args{
				flowerbed: []int{1, 0, 0, 0, 1, 0, 0},
				n:         2,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.flowerbed, tt.args.n); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
