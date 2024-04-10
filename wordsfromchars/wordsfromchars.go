// Find Words That Can Be Formed by Characters
//
// You are given an array of strings words and a string chars.
//
// A string is good if it can be formed by characters from chars (each character can only be used once).
//
// Return the sum of lengths of all good strings in words.
//
// #easy
// https://leetcode.com/problems/find-words-that-can-be-formed-by-characters
package wordsfromchars

// time O(n*m)
// space O(1)
func Solution(words []string, chars string) int {
	limits := make([]int, 26) // lowercase english chars
	for _, r := range chars {
		limits[r-'a']++
	}
	sums := 0 // sum of len of all "Good" words
	for _, word := range words {
		if good(word, limits) {
			sums += len(word)
		}
	}
	return sums
}

func good(word string, limits []int) bool {
	wordChars := make([]int, 26) // lowercase english chars
	for _, r := range word {
		index := r - 'a'
		wordChars[index]++
		// if we have more chars at this alphabet index and possible
		if wordChars[index] > limits[index] {
			return false
		}
	}
	return true
}
