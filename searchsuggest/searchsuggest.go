// Search Suggestions System
//
// You are given an array of strings products and a string searchWord.
//
// Design a system that suggests at most three product names from products after each character of searchWord is typed. Suggested products should have common prefix with searchWord. If there are more than three products with a common prefix return the three lexicographically minimums products.
//
// Return a list of lists of the suggested products after each character of searchWord is typed.
//
// https://leetcode.com/problems/search-suggestions-system
// #trie
package searchsuggest

import (
	"sort"
	"strings"
)

func Solution(products []string, searchWord string) [][]string {
	sort.Strings(products)
	result := make([][]string, len(searchWord))
	// for each character typed
	for i := range searchWord {
		// for each product
		for _, p := range products {
			if strings.HasPrefix(p, searchWord[:i+1]) {
				result[i] = append(result[i], p)
				// ... up to 3 items
				if len(result[i]) == 3 {
					break
				}
			}
		}
	}
	return result
}
