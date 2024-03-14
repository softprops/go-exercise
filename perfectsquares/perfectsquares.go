// Perfect Squares
//
// Given an integer n, return the least number of perfect square numbers that sum to n.
//
// A perfect square is an integer that is the square of an integer; in other words, it is the product (multiplication) of some integer with itself. For example, 1, 4, 9, and 16 are perfect squares while 3 and 11 are not.
//
// https://leetcode.com/problems/perfect-squares/description/
package perfectsquares

import (
	"math"
)

func Solution(n int) int {
	cache := make([]int, n+1)
	for i := range cache {
		cache[i] = math.MaxInt
	}
	cache[0] = 0

	for i := 1; i <= n; i++ {
		val := math.MaxInt
		for j := 1; j*j <= i; j++ {
			val = min(val, 1+cache[i-j*j])
		}
		cache[i] = val
	}
	return cache[n]
}
