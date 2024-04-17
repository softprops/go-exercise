// Maximum Population Year
//
// You are given a 2D integer array logs where each logs[i] = [birthi, deathi] indicates the birth and death years of the ith person.
//
// The population of some year x is the number of people alive during that year. The ith person is counted in year x's population if x is in the inclusive range [birthi, deathi - 1]. Note that the person is not counted in the year that they die.
//
// Return the earliest year with the maximum population.
//
// constraints
// 1 <= logs.length <= 100
// 1950 <= birthi < deathi <= 2050
//
// https://leetcode.com/problems/maximum-population-year/
package maxpopulationyear

// time O(n)
// space O(n)
func Solution(logs [][]int) int {
	pops, max := make([]int, 2051), 1950
	for _, l := range logs {
		// record "getting on and off the bus"
		pops[l[0]]++
		pops[l[1]]--
	}
	for i := 1950; i < 2050; i++ {
		// prefix sum keep track of max
		pops[i] += pops[i-1]
		if pops[i] > pops[max] {
			max = i
		}
	}
	return max
}
