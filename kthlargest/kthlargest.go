// Kth Largest Element in an Array
//
// Given an integer array nums and an integer k, return the kth largest element in the array.
//
// Note that it is the kth largest element in the sorted order, not the kth distinct element.
//
// Can you solve it without sorting?
//
// #array #divideandconquer #sorting #heap(priorty queue) #quickselect
//
// https://leetcode.com/problems/kth-largest-element-in-an-array/
package kthlargest

import "math/rand"

func Solution(nums []int, k int) int {
	return sorted(nums, k)[len(nums)-k]
}

func sorted(nums []int, k int) []int {
	l, r, k := 0, len(nums)-1, len(nums)-k
	for l < r {
		p := partition(nums, l, r)
		if p < k {
			// look right
			l = p + 1
		} else if p > k {
			// look left
			r = p - 1
		} else {
			// we've sorted it out
			break
		}
	}
	return nums
}

func partition(nums []int, l, r int) int {
	pivot := rand.Intn(r-l+1) + l // random pivot
	pivotVal := nums[pivot]
	nums[pivot], nums[r] = nums[r], nums[pivot]
	for i := l; i < r; i++ {
		if nums[i] <= pivotVal {
			nums[l], nums[i] = nums[i], nums[l]
			l++
		}
	}
	nums[l], nums[r] = nums[r], nums[l]
	return l
}
