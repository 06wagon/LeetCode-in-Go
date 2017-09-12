package Problem0098

import (
	"github.com/aQuaYi/LeetCode-in-Golang/kit"
)

type TreeNode = kit.TreeNode

func isValidBST(root *TreeNode) bool {
	// Golang int 类型的最小值与最大值
	MIN, MAX := -1<<63, 1<<63-1

	return cur(MIN, MAX, root)
}

// 以递归的方式，检查 root.Val 是否在 (min, max) 范围内。
func cur(min, max int, root *TreeNode) bool {
	if root == nil {
		return true
	}

	return min < root.Val && root.Val < max &&
		cur(min, root.Val, root.Left) &&
		cur(root.Val, max, root.Right)
}
