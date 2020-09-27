package solutions

func checEquivalentTree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil {
		return false
	}
	if t == nil {
		return false
	}

	if s.Val != t.Val {
		return false
	}

	return checEquivalentTree(s.Left, t.Left) && checEquivalentTree(s.Right, t.Right)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	if s.Val == t.Val {
		if checEquivalentTree(s, t) {
			return true
		}
	}
	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}
