// Find K Pairs with Smallest Sums
//
// You are given two integer arrays nums1 and nums2 sorted in non-decreasing order and an integer k.
//
// Define a pair (u, v) which consists of one element from the first array and one element from the second array.
//
// Return the k pairs (u1, v1), (u2, v2), ..., (uk, vk) with the smallest sums.
//
// https://leetcode.com/problems/find-k-pairs-with-smallest-sums/
package ksmallsumpairs

import (
	"container/heap"
)

type Pair struct{ index1, index2, sum int }

type PriorityQueue []Pair

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].sum < pq[j].sum }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq PriorityQueue) Empty() bool        { return len(pq) == 0 }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Pair))
}

func (h *PriorityQueue) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func Solution(nums1, nums2 []int, k int) [][]int {
	res := [][]int{}

	pq := &PriorityQueue{}
	heap.Init(pq)
	for i := 0; i < len(nums1); i++ {
		heap.Push(pq, Pair{i, 0, nums1[i] + nums2[0]})
	}
	for !pq.Empty() && len(res) < k {
		small := heap.Pop(pq).(Pair)
		res = append(res, []int{nums1[small.index1], nums2[small.index2]})
		if small.index2 < len(nums2)-1 {
			heap.Push(pq, Pair{small.index1, small.index2 + 1, nums1[small.index1] + nums2[small.index2+1]})
		}
	}
	return res
}
