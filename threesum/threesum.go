// 3Sum
//
// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
//
// Notice that the solution set must not contain duplicate triplets.
//
// https://leetcode.com/problems/3sum
package threesum

import (
	"slices"
)

// intution. if we fix pick one num, this becomes the two sum problem
// time O(n^2)
// space O(1)
func Solution(nums []int) [][]int {
	target := 0
	slices.Sort(nums) // O(n logn) time
	cache := make(map[[3]int]struct{})
	for i := range nums {
		j := i + 1
		k := len(nums) - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				cache[[3]int{nums[i], nums[j], nums[k]}] = struct{}{}
				j++
				k--
			} else if sum < target {
				j++
			} else {
				k--
			}
		}
	}
	keys := make([][]int, 0, len(cache))
	for k := range cache {
		keys = append(keys, k[:])
	}
	return keys
}
