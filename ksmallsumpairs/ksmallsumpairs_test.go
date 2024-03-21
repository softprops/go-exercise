// Find K Pairs with Smallest Sums
//
// You are given two integer arrays nums1 and nums2 sorted in non-decreasing order and an integer k.
//
// Define a pair (u, v) which consists of one element from the first array and one element from the second array.
//
// Return the k pairs (u1, v1), (u2, v2), ..., (uk, vk) with the smallest sums.
//
// https://leetcode.com/problems/find-k-pairs-with-smallest-sums/
package ksmallsumpairs

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "a",
			args: args{
				nums1: []int{1, 7, 11},
				nums2: []int{2, 4, 6},
				k:     3,
			},
			want: [][]int{
				{1, 2}, {1, 4}, {1, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.nums1, tt.args.nums2, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
