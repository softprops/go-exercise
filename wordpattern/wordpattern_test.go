// Word Pattern
//
// Given a pattern and a string s, find if s follows the same pattern.
//
// Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s.
//
// https://leetcode.com/problems/word-pattern/
package wordpattern

import (
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		s       string
		pattern string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "a",
			args: args{
				s:       "dog cat cat dog",
				pattern: "abba",
			},
			want: true,
		},
		{
			name: "b",
			args: args{
				s:       "dog cat cat fish",
				pattern: "abba",
			},
			want: false,
		},
		{
			name: "c",
			args: args{
				s:       "dog cat cat dog",
				pattern: "aaaa",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.s, tt.args.pattern); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
