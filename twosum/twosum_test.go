package twosum

import (
	"reflect"
	"slices"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		expect []int
	}{
		{
			name:   "a",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			expect: []int{0, 1},
		},
		{
			name:   "b",
			nums:   []int{3, 2, 4},
			target: 6,
			expect: []int{1, 2},
		},
		{
			name:   "c",
			nums:   []int{3, 3},
			target: 6,
			expect: []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Solution(tt.nums, tt.target)
			slices.Sort(actual)
			if !reflect.DeepEqual(actual, tt.expect) {
				t.Fatalf(`expected %v but got %v`, tt.expect, actual)
			}
		})

	}

}
