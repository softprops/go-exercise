package maxprofit

import "math"

func Solution(prices []int) int {
	low, profit := math.MaxInt, 0
	for _, price := range prices {
		low = min(low, price)
		profit = max(profit, price-low)
	}
	return profit
}
