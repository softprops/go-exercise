// Remove Duplicate Letters
//
// Given a string s, remove duplicate letters so that every letter appears once and only once. You must make sure your result is
// the smallest in lexicographical order
// among all possible results.
//
// https://leetcode.com/problems/remove-duplicate-letters/
package removedupletters

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
				s: "bcabc",
			},
			want: "abc",
		},
		{
			name: "b",
			args: args{
				s: "cbacdcbc",
			},
			want: "acdb",
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
