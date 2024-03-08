// Intersection of Two Arrays
//
// https://leetcode.com/problems/intersection-of-two-arrays/
package intersection

func Solution(nums1, nums2 []int) []int {
	// decide which is smaller and which is larger
	// choosing the smaller is a more effective use of memory
	// i.e. you don't want a allocate a set of 1000 distinct nums if you're comparison
	// only has 1
	s, l := nums1, nums2
	if len(nums1) > len(nums2) {
		s, l = nums2, nums1
	}

	// small set
	ss := NewSet()
	for _, n := range s {
		ss.Add(n)
	}

	// intersection set
	is := NewSet()
	for _, n := range l {
		if ss.Contains(n) {
			is.Add(n)
		}
	}

	return is.Values()
}

type Set struct {
	data map[int]struct{}
}

func NewSet() *Set {
	return &Set{data: make(map[int]struct{})}
}

func (s *Set) Add(n int) {
	s.data[n] = struct{}{}
}

func (s *Set) Contains(n int) bool {
	_, ok := s.data[n]
	return ok
}

func (s *Set) Values() []int {
	keys := make([]int, 0, len(s.data))
	for k := range s.data {
		keys = append(keys, k)
	}
	return keys
}
