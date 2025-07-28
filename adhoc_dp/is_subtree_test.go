package adhoc_dp

import (
	"fmt"
	"strings"
)

// tree hack approach
func serializeTree(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	return "," + fmt.Sprintf("%d", root.Val) + "-" + serializeTree(root.Left) + "-" + serializeTree(root.Right)
}

func isSubtree_SerializationHack(root *TreeNode, subRoot *TreeNode) bool {

	rootStr := serializeTree(root)
	subRootStr := serializeTree(subRoot)

	return strings.Contains(rootStr, subRootStr)
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {

	if root == nil {
		return false
	}
	if isSame(root, subRoot) {
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSame(root1 *TreeNode, root2 *TreeNode) bool {

	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}

	return root1.Val == root2.Val && isSame(root1.Left, root2.Left) && isSame(root1.Right, root2.Right)
}

// another is to use reursion
