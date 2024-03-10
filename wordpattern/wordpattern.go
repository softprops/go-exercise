// Word Pattern
//
// Given a pattern and a string s, find if s follows the same pattern.
//
// Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s.
//
// https://leetcode.com/problems/word-pattern/
package wordpattern

import "strings"

func Solution(s, pattern string) bool {
	words := strings.Split(s, " ")
	if len(words) != len(pattern) {
		return false
	}
	// bijection means bi-directional mapping so two maps are needed
	smap := make(map[string]rune)
	pmap := make(map[rune]string)
	for i, r := range pattern {
		w := words[i]
		if rw, ok := pmap[r]; ok {
			// there existing a pattern mapping but not to this word
			if rw != w {
				return false
			}
		} else {
			// there existed a word mapping but not to this rune
			if _, ok := smap[w]; ok {
				return false
			} else {
				pmap[r], smap[w] = w, r
			}
		}
	}
	return true
}
