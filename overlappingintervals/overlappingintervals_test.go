package overlappingintervals

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "a",
			args: args{
				intervals: [][]int{
					{1, 2},
					{2, 3},
					{3, 4},
					{1, 3},
				},
			},
			want: 1,
		},
		{
			name: "b",
			args: args{
				intervals: [][]int{
					{1, 2},
					{1, 2},
					{1, 2},
				},
			},
			want: 2,
		},
		{
			name: "c",
			args: args{
				intervals: [][]int{
					{1, 2},
					{2, 3},
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.intervals); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
