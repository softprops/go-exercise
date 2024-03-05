// Reverse Vowels of a String
//
// Given a string s, reverse only all the vowels in the string and return it.
//
// The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.
//
// https://leetcode.com/problems/reverse-vowels-of-a-string/description
package reversevowels

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
				s: "hello",
			},
			want: "holle",
		},
		{
			name: "b",
			args: args{
				s: "leetcode",
			},
			want: "leotcede",
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
