// Remove Linked List Elements
//
// Given the head of a linked list and an integer val, remove all the nodes of the linked list that has Node.val == val, and return the new head.
//
// https://leetcode.com/problems/remove-linked-list-elements/
package removelinkedlistelms

func Solution(head *ListNode, val int) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	// clear this head
	if head.Val == val {
		return Solution(head.Next, val)
	}
	next := head.Next
	if next.Val == val {
		next = head.Next.Next
	}
	head.Next = Solution(next, val)

	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
