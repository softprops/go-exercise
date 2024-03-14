// Top K Frequent Elements
//
// Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.
//
// https://leetcode.com/problems/top-k-frequent-elements/
package topkfreq

func Solution(nums []int, k int) []int {
	// freq count
	freq := make(map[int]int)
	for _, n := range nums {
		freq[n]++
	}

	// return slice of top k (quickselect) sorted values
	return sorted(freq, k)[len(freq)-k:]
}

func sorted(freq map[int]int, k int) []int {
	// slice of keys
	keys := make([]int, 0, len(freq))
	for key := range freq {
		keys = append(keys, key)
	}

	// quick select
	l, r, k := 0, len(keys)-1, len(keys)-k
	for l < r {
		p := partition(freq, keys, l, r)
		if p < k {
			// look right
			l = p + 1
		} else if p > k {
			// look left
			r = p - 1
		} else {
			// we've "sorted" it out
			break
		}
	}

	return keys
}

func partition(freq map[int]int, keys []int, l, r int) int {
	pivot := keys[r]
	index := l
	// shift everthing on the left that <= our pivot to the new pivot
	// and incr the pivot position
	for i := l; i < r; i++ {
		if freq[keys[i]] <= freq[pivot] {
			keys[i], keys[index] = keys[index], keys[i]
			index++
		}
	}
	// swap the last r value with pivot
	keys[index], keys[r] = keys[r], keys[index]
	return index
}
