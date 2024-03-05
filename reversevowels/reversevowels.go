// Reverse Vowels of a String
//
// Given a string s, reverse only all the vowels in the string and return it.
//
// The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.
//
// https://leetcode.com/problems/reverse-vowels-of-a-string/description
package reversevowels

import (
	"bytes"
)

func Solution(s string) string {
	vowels := []byte("aeiouAEIOU")
	b := []rune(s)
	for l, r := 0, len(s)-1; l < r; {
		if !bytes.ContainsRune(vowels, b[l]) {
			l++
		} else if !bytes.ContainsRune(vowels, b[r]) {
			r--
		} else {
			b[l], b[r] = b[r], b[l]
			l++
			r--
		}
	}
	return string(b)
}
