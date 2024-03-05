// Longest Common Prefix
//
// Write a function to find the longest common prefix string amongst an array of strings.
//
// If there is no common prefix, return an empty string "".
//
// #trie
// https://leetcode.com/problems/longest-common-prefix/description/
package commonprefix

import "slices"

func Solution(strs []string) string {
	// enforce invariant first will always shorter than last
	slices.Sort(strs)
	l, r := strs[0], strs[len(strs)-1]
	i := 0
	for i < len(l) && i < len(r) {
		if l[i] == r[i] {
			i++
		} else {
			break
		}
	}
	return l[:i]
}
