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

import (
	"regexp"
	"strings"
)

func Solution(s string) string {
	words := regexp.MustCompile(`\s+`).Split(strings.Trim(s, " "), -1)
	for l, r := 0, len(words)-1; l < r; {
		words[l], words[r] = words[r], words[l]
		l++
		r--
	}
	return strings.Join(words, " ")
}
