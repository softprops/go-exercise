package searchindex

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "a",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 5,
			},
			want: 2,
		},
		{
			name: "b",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 2,
			},
			want: 1,
		},
		{
			name: "c",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 7,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
