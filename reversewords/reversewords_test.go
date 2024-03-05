// Reverse Words in a String
//
// A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.
//
// Return a string of the words in reverse order concatenated by a single space.
//
// Note that s may contain leading or trailing spaces or multiple spaces between two words. The returned string should only have a single space separating the words. Do not include any extra spaces.
//
// https://leetcode.com/problems/reverse-words-in-a-string/description
package reversewords

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "a",
			args: args{
				s: "the sky is blue",
			},
			want: "blue is sky the",
		},
		{
			name: "b",
			args: args{
				s: "  hello world  ",
			},
			want: "world hello",
		},
		{
			name: "c",
			args: args{
				s: "a good   example",
			},
			want: "example good a",
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
