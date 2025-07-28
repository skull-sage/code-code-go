package adhoc_dp

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func diameterOfBinaryTree(root *TreeNode) int {

	if root == nil {
		return 0
	}

	maxD := 0

	var findHeight func(node *TreeNode) int
	findHeight = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftH := findHeight(node.Left)
		rightH := findHeight(node.Right)

		dm := 1 + leftH + rightH
		if maxD < dm {
			maxD = dm
		}

		fmt.Println(node.Val, ":", "leftH", leftH, "rightH", rightH)
		// return height for passing to root
		return 1 + max(leftH, rightH)
	}

	return maxD

}
