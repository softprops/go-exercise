// Minimum Number of Arrows to Burst Balloons
//
// https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/description
//
// #intervals
package burstingballoons

import (
	"log"
	"slices"
)

func Solution(points [][]int) int {
	if len(points) < 2 {
		return len(points)
	}
	// sort by end
	slices.SortFunc(points, func(a, b []int) int {
		return a[1] - b[1]
	})
	log.Printf("sorted %v", points)
	// 1 arrow for first point and right edge
	arrows, r := 1, points[0][1]
	for _, point := range points[1:] {
		if point[0] > r {
			arrows++
			r = point[1]
		}
	}
	return arrows
}
