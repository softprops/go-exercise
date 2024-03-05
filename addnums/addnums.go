package addnums

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(v int, n *ListNode) *ListNode {
	return &ListNode{Val: v, Next: n}
}

func (l *ListNode) String() string {
	return fmt.Sprintf("ListNode{Val:%d,Next:%s}", l.Val, l.Next)
}

func Solution(a, b *ListNode) *ListNode {
	// "loop" through next values building up a result carrying over remainder
	result := &ListNode{}
	for tail, carry := result, 0; a != nil || b != nil || carry != 0; {
		va, vb := 0, 0
		if a != nil {
			va = a.Val
		}
		if b != nil {
			vb = b.Val
		}

		// math
		sum := va + vb + carry
		digit := sum % 10
		carry = sum / 10

		next := NewListNode(digit, nil)
		tail.Next = next
		tail = next

		// drive the loop
		if a != nil {
			a = a.Next
		}
		if b != nil {
			b = b.Next
		}
	}
	return result.Next
}
