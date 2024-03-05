package firstbadversion

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		n   int
		bad int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "a",
			args: args{
				n:   5,
				bad: 4,
			},
			want: 4,
		},
		{
			name: "b",
			args: args{
				n:   1,
				bad: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.n); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
