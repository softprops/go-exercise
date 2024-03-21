// Maximum Bags With Full Capacity of Rocks
//
// You have n bags numbered from 0 to n - 1. You are given two 0-indexed integer arrays capacity and rocks. The ith bag can hold a maximum of capacity[i] rocks and currently contains rocks[i] rocks. You are also given an integer additionalRocks, the number of additional rocks you can place in any of the bags.
//
// Return the maximum number of bags that could have full capacity after placing the additional rocks in some bags.
//
// https://leetcode.com/problems/maximum-bags-with-full-capacity-of-rocks/description/
package bagofrocks

import "slices"

// time O(n log(n)) for sort
// space O(1)
func Solution(capacity, rocks []int, additionalRocks int) int {
	for i := 0; i < len(capacity); i++ {
		capacity[i] -= rocks[i]
	}
	slices.Sort(capacity)
	for i := 0; i < len(capacity); i++ {
		additionalRocks -= capacity[i]
		if additionalRocks < 0 {
			return i
		}
	}
	return len(capacity)
}
