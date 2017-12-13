package Problem0572

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}

	if s == nil || t == nil {
		return false
	}

	if s.Val == t.Val {
		return isSubtree(s.Left, t.Left) && isSubtree(s.Right, t.Right)
	}

	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}
