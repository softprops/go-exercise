package wordsfromchars

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		words []string
		chars string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "a",
			args: args{
				words: []string{"cat", "bt", "hat", "tree"},
				chars: "atach",
			},
			want: 6,
		},
		{
			name: "b",
			args: args{
				words: []string{"hello", "world", "leetcode"},
				chars: "welldonehoneyr",
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.words, tt.args.chars); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
