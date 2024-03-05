package kidswithcandy

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		candies      []int
		extraCandies int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "a",
			args: args{
				candies:      []int{2, 3, 5, 1, 3},
				extraCandies: 3,
			},
			want: []bool{true, true, true, false, true},
		},
		{
			name: "b",
			args: args{
				candies:      []int{4, 2, 1, 1, 2},
				extraCandies: 1,
			},
			want: []bool{true, false, false, false, false},
		},
		{
			name: "c",
			args: args{
				candies:      []int{12, 1, 12},
				extraCandies: 10,
			},
			want: []bool{true, false, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.candies, tt.args.extraCandies); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
