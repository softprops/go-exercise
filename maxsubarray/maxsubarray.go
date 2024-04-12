// Maximum Subarray
//
// Given an integer array nums, find the
// subarray with the largest sum, and return its sum.
//
// https://leetcode.com/problems/maximum-subarray
package maxsubarray

// Kadanes algorithm
// time O(n)
// space O(1)
func Solution(nums []int) int {
	current, overall := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		current += nums[i]
		if current < nums[i] {
			current = nums[i]
		}
		if overall < current {
			overall = current
		}
	}
	return overall
}
