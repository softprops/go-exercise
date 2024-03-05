package containsdup

func Solution(nums []int) bool {
	lookup := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := lookup[num]; ok {
			return true
		}
		lookup[num] = struct{}{}
	}
	return false
}
