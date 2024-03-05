package romanint

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "a",
			args: args{
				s: "III",
			},
			want: 3,
		},
		{
			name: "b",
			args: args{
				s: "LVIII",
			},
			want: 58,
		},
		{
			name: "c",
			args: args{
				s: "MCMXCIV",
			},
			want: 1994,
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
