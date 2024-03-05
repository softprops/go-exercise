package productofarray

func Solution(nums []int) []int {
	products := make([]int, len(nums))
	// nothing precedes first element
	products[0] = 1
	// walk forwards
	for i := 1; i < len(nums); i++ {
		// assign element to the prev result * the previous number
		products[i] = products[i-1] * nums[i-1]
	}
	// walk backwards
	for right, i := 1, len(nums)-1; i >= 0; i-- {
		// multiply each element but the rolling sum of elements that precede it
		products[i] *= right
		right *= nums[i]
	}
	return products
}
