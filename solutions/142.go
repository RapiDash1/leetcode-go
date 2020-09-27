package solutions

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}

	slow := head
	fast := &ListNode{head.Val, head.Next}

	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast.Val == slow.Val && fast.Next == slow.Next {
			break
		}
	}

	for head != nil && slow != nil && head != slow {
		head = head.Next
		slow = slow.Next
	}

	if head == nil || slow == nil {
		return nil
	}

	return head
}
