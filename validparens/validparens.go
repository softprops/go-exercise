package validparens

func Solution(s string) bool {
	stack := NewStack[rune]()
	for _, r := range s {
		switch r {
		case '(', '{', '[':
			stack.Push(r)
		default:
			switch stack.Pop() {
			case '(':
				return r == ')'
			case '[':
				return r == ']'
			case '{':
				return r == '}'
			default:
				return false
			}
		}
	}
	return stack.Len() == 0
}

type Stack[T any] struct {
	items []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(a T) T {
	s.items = append(s.items, a)
	return a
}

func (s *Stack[T]) Pop() T {
	l := s.Len()
	item := s.items[l-1]
	s.items = s.items[:l-1]
	return item
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}
