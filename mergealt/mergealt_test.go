package mergealt

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "a",
			args: args{
				word1: "abcd",
				word2: "pq",
			},
			want: "apbqcd",
		},
		{
			name: "b",
			args: args{
				word1: "ab",
				word2: "pqrs",
			},
			want: "apbqrs",
		},
		{
			name: "c",
			args: args{
				word1: "abcd",
				word2: "pq",
			},
			want: "apbqcd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
