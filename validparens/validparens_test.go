package validparens

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "a",
			args: args {
				s: "()",
			},
			want: true,
		},
		{
			name: "b",
			args: args {
				s: "()[]{}",
			},
			want: true,
		},
		{
			name: "c",
			args: args {
				s: "(]",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.s); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
