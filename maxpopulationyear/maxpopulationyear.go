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
// space O(maxYear - minYear + 1)
func Solution(logs [][]int) int {
	// minYear, maxYear = 1950, 2050
	minYear, maxYear := lifeRange(logs)
	return solution(logs, minYear, maxYear)
}

func solution(logs [][]int, firstYear, lastYear int) int {
	// minimize space complexity
	pops := make([]int, lastYear-firstYear+1)
	for _, l := range logs {
		// record "getting on and off the bus"
		pops[l[0]-firstYear]++
		pops[l[1]-firstYear]--
	}
	maxIndex := 0
	for i := 1; i < len(pops); i++ {
		// prefix sum keep track of max
		pops[i] += pops[i-1]
		if pops[i] > pops[maxIndex] {
			maxIndex = i
		}
	}
	return maxIndex + firstYear
}

func lifeRange(logs [][]int) (int, int) {
	minYear, maxYear := 0, 0
	for _, l := range logs {
		minYear = min(minYear, l[0])
		maxYear = max(maxYear, l[1])
	}
	return minYear, maxYear
}
