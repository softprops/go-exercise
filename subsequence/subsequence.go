// Is Subsequence
//
// https://leetcode.com/problems/is-subsequence/
package subsequence

func Solution(s, t string) bool {
	if len(s) > len(t) {
		return false
	}
	i := 0
	for _, r := range s {
		if i == len(t) {
			return false
		}
		for t[i] != byte(r) {
			i++
			if i == len(t) {
				return false
			}
		}
	}
	return true
}
