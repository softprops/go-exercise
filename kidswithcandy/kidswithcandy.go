// Kids With the Greatest Number of Candies
//
// There are n kids with candies. You are given an integer array candies, where each candies[i] represents the number of candies the ith kid has, and an integer extraCandies, denoting the number of extra candies that you have.
//
// Return a boolean array result of length n, where result[i] is true if, after giving the ith kid all the extraCandies, they will have the greatest number of candies among all the kids, or false otherwise.
//
// Note that multiple kids can have the greatest number of candies.
// https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/description
package kidswithcandy

import "math"

func Solution(candies []int, extraCandies int) []bool {
	max := math.MinInt
	for _, n := range candies {
		if n > max {
			max = n
		}
	}
	results := make([]bool, len(candies))
	for i, n := range candies {
		results[i] = n+extraCandies >= max
	}
	return results
}
