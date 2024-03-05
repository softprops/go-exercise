package kthlargest

import "log"

// note (l + r) >>> 1 is the same as divide in half with no remainder
func Solution(nums []int, k int) int {
	l, r := 0, len(nums)-1
	target := len(nums) - k
	// loop until we find our target
	for {
		mid := l + (r-l)/2
		pivot := hoaresPartition(nums, l, r, mid)
		log.Printf("target %v l %v r %v pivot %v nums %#v", target, l, r, pivot, nums)
		if pivot == target {
			return nums[pivot]
		} else if pivot > target {
			r = pivot - 1
		} else {
			l = pivot + 1
		}
	}
}

// https://www.geeksforgeeks.org/hoares-vs-lomuto-partition-scheme-quicksort/?ref=ml_lbp

func hoaresPartition(nums []int, l, r, mid int) int {
	pivot := nums[mid]
	for l <= r {
		//log.Printf("l %v r %v pivot %v", l, r, pivot)
		for nums[l] < pivot {
			l++
		}
		for pivot < nums[r] {
			r--
		}
		if l <= r {
			//log.Printf("swapping %v and %v", nums[l], nums[r])
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}

	return l
}

func lumutoPartition(nums []int, l, r, mid int) int {
	pivot := nums[mid]
	// swap
	nums[mid], nums[r] = nums[r], nums[mid]
	result := l
	for i := l; i < r; i++ {
		if nums[i] < pivot {
			nums[l], nums[mid] = nums[mid], nums[l]
			result++
		}
	}
	// swap last
	nums[r], nums[mid] = nums[mid], nums[r]
	return result
}
