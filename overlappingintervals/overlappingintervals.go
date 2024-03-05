// Non-overlapping Intervals
//
// Given an array of intervals intervals where intervals[i] = [starti, endi], return the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.
//
// https://leetcode.com/problems/non-overlapping-intervals/description
package overlappingintervals

import "slices"

func Solution(intervals [][]int) int {
	// sort according to interval end
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[1] - b[1]
	})
	del := 0
	// keep a "stack" of intervals that are non overlapping
	stack := [][]int{}
	for _, interval := range intervals {
		// count the number of intervals that start with a value less than the ending
		// of the previous interval end value
		if len(stack) != 0 && interval[0] < stack[len(stack)-1][1] {
			del++
		} else {
			stack = append(stack, interval)
		}
	}
	return del
}
