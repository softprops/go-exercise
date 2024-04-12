package sortlinkedlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n ListNode) String() string {
	return fmt.Sprintf("ListNode{%v,%v}", n.Val, n.Next)
}

// time O(n logn) merge sort
// space O(n logn) (stack space)
func Solution(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// split list in half
	var prev *ListNode = nil
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil // detach

	return merge(Solution(head), Solution(slow))
}

func merge(left, right *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for left != nil && right != nil {
		if left.Val < right.Val {
			current.Next = left
			left = left.Next
		} else {
			current.Next = right
			right = right.Next
		}
		current = current.Next
	}

	if left != nil {
		current.Next = left
	}
	if right != nil {
		current.Next = right
	}

	return dummy.Next
}
