// H-Index II
//
// Given an array of integers citations where citations[i] is the number of citations a researcher received for their ith paper and citations is sorted in ascending order, return the researcher's h-index.
//
// According to the definition of h-index on Wikipedia: The h-index is defined as the maximum value of h such that the given researcher has published at least h papers that have each been cited at least h times.
//
// hint: You must write an algorithm that runs in logarithmic time.
//
// #array #binarysearch
//
// https://leetcode.com/problems/h-index-ii/description/
package hindex2

// citations are sorted in ascending order
func Solution(citations []int) int {
	l, r := 0, len(citations)-1
	for l <= r {
		mid := l + (r-l)/2
		if citations[mid] < len(citations)-mid {
			// look right
			l = mid + 1
		} else {
			// look left
			r = mid - 1
		}
	}
	return len(citations) - l
}
