// Remove Duplicate Letters
//
// Given a string s, remove duplicate letters so that every letter appears once and only once. You must make sure your result is
// the smallest in lexicographical order
// among all possible results.
//
// constraints
// * 1 <= s.length <= 104
// * s consists of lowercase English letters.
//
// https://leetcode.com/problems/remove-duplicate-letters/
package removedupletters

// scales linearly with len(s)
//
// time O(n) iterate over each rune in the string once
//
// space O(n)
func Solution(s string) string {
	seen := make(map[rune]bool)
	stack := NewStack()
	lastIndex := make(map[rune]int)
	for i, r := range s {
		lastIndex[r] = i
	}
	for i, r := range s {
		if _, ok := seen[r]; !ok {
			for !stack.Empty() && r < stack.Peek() && i < lastIndex[stack.Peek()] {
				last := stack.Pop()
				delete(seen, last)
			}
			seen[r] = true
			stack.Push(r)
		}
	}
	return stack.String()
}

type Stack struct {
	data []rune
}

func NewStack() *Stack {
	return &Stack{data: []rune{}}
}

func (s *Stack) String() string {
	return string(s.data)
}

func (s *Stack) Push(r rune) {
	s.data = append(s.data, r)
}

func (s *Stack) Pop() rune {
	lastIndex := len(s.data) - 1
	r := s.data[lastIndex]
	s.data = s.data[:lastIndex]
	return r
}

func (s *Stack) Peek() rune {
	return s.data[len(s.data)-1]
}

func (s *Stack) Empty() bool {
	return len(s.data) < 1
}
