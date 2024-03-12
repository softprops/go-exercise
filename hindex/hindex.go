// H-Index
//
// Given an array of integers citations where citations[i] is the number of citations a researcher received for their ith paper, return the researcher's h-index.
//
// According to the definition of h-index on Wikipedia: The h-index is defined as the maximum value of h such that the given researcher has published at least h papers that have each been cited at least h times.
//
// https://leetcode.com/problems/h-index/description/
package hindex

import (
	"slices"
)

func Solution(citations []int) int {
	slices.Sort(citations)
	h := len(citations)
	for _, c := range citations {
		if c < h {
			h--
		}
	}

	return h
}
