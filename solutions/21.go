package solutions

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	iterNode := &ListNode{Val: 0}
	retNode := iterNode

	for l1 != nil && l2 != nil {
		nodeVal := 0
		if l1.Val < l2.Val {
			nodeVal = l1.Val
			l1 = l1.Next
		} else {
			nodeVal = l2.Val
			l2 = l2.Next
		}
		iterNode.Next = &ListNode{Val: nodeVal}
		iterNode = iterNode.Next
	}

	if l1 != nil {
		iterNode.Next = l1
	}
	if l2 != nil {
		iterNode.Next = l2
	}

	return retNode.Next
}
