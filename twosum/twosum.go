package twosum

func Solution(nums []int, target int) []int {
	lookup := make(map[int]int)
	for i, num := range nums {
		if diff, ok := lookup[target-num]; ok && diff != i {
			return []int{i, diff}
		}
		lookup[num] = i
	}
	return []int{}
}
