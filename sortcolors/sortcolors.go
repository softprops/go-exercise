//	Sort Colors
//
// https://leetcode.com/problems/sort-colors
package sortcolors

// intuition. dutch national flag algo problem
// time O(n)
// space O(1)
func Solution(nums []int) {
	low, mid, high := 0, 0, len(nums)-1
	for mid <= high {
		switch nums[mid] {
		case 0:
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
		case 1:
			mid++
		case 2:
			nums[high], nums[mid] = nums[mid], nums[high]
			high--
		default:
		}
	}
}
