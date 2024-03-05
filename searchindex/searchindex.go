package searchindex

// https://leetcode.com/problems/search-insert-position/description/?envType=study-plan-v2&envId=top-interview-150
func Solution(nums []int, target int) int {
	l := 0
	for r := len(nums) - 1; l <= r; {
		mid := l + (r-l)/2
		val := nums[mid]
		if val == target {
			return mid
		}
		if val > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}
