// https://leetcode.com/problems/first-bad-version/description/
package firstbadversion

// intuition: versions are ordered (binary search)
func Solution(n int, badVersion func(int) bool) int {
	l, r := 1, n
	for l <= r {
		mid := l + (r-l)/2
		if badVersion(mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}
