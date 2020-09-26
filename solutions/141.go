package solutions

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow := head
	fast := &ListNode{head.Val, head.Next}

	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast.Val == slow.Val && fast.Next == slow.Next {
			return true
		}
	}
	return false
}
