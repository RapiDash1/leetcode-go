package solutions

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prevNode *ListNode
	for {
		nextNode := head.Next
		head.Next = prevNode
		prevNode = head
		if nextNode == nil {
			break
		}
		head = nextNode
	}
	return head
}
