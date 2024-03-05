package maxprofit

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		prices []int
		expect int
	}{
		{
			name:   "a",
			prices: []int{7, 1, 5, 3, 6, 4},
			expect: 5,
		},
		{
			name:   "b",
			prices: []int{7, 6, 4, 3, 1},
			expect: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Solution(tt.prices)
			if !reflect.DeepEqual(actual, tt.expect) {
				t.Fatalf(`expected %v but got %v`, tt.expect, actual)
			}
		})
	}
}
