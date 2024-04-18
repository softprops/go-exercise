package firstbadversion

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		n   int
		bad func(int) bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "a",
			args: args{
				n: 5,
				bad: func(i int) bool {
					return i < 5 && i > 3
				},
			},
			want: 4,
		},
		{
			name: "b",
			args: args{
				n: 1,
				bad: func(i int) bool {
					return i > 0
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.n, tt.args.bad); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
