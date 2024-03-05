// Given a non-negative integer x, return the square root of x rounded down to the nearest integer. The returned integer should be non-negative as well.
//
// You must not use any built-in exponent function or operator.
//
// https://leetcode.com/problems/sqrtx/description/
//
// #binarysearch
package sqrt

func Solution(x int) int {
	// identity multiplication
	if x < 2 {
		return x
	}

	// define range, starting at 1
	l, r := 1, x
	for l <= r {
		mid := l + (r-l)/2
		square := mid * mid

		if square == x {
			// we've found our square root
			return mid
		}
		if square > x {
			// decease upper bound
			r = mid - 1
		} else {
			// increase lower bound
			l = mid + 1
		}
	}

	return r
}
