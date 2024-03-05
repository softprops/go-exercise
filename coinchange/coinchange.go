// Coin Change
//
// https://leetcode.com/problems/coin-change/description/
//
// https://dev.to/shivams136/leetcode-322-coin-change-solution-4kmd
//
// #array #dp #dfs
package coinchange

import (
	"math"
)

func Solution(coins []int, amount int) int {
	unknown := math.MaxInt
	cache := make([]int, amount+1)
	for value := 1; value <= amount; value++ {
		count := unknown
		for _, c := range coins {
			diff := value - c
			if diff >= 0 {
				cached := cache[diff]
				if cached != unknown {
					// add a coin
					count = min(1+cached, count)
				}
			}
		}
		cache[value] = count
	}
	if cache[amount] == unknown {
		return -1
	}
	return cache[amount]
}
