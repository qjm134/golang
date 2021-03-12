/*
输入一棵二叉树的根节点，判断该树是不是平衡二叉树。
如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。
*/

package isBalance

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	_, result := getHight(root)
	return result
}

func getHight(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	leftHight, leftResult := getHight(root.Left)
	if leftResult == false {
		return 0, false
	}
	rightHight, rightResult := getHight(root.Right)
	if rightResult == false {
		return 0, false
	}

	max := 0
	min := 0
	if leftHight > rightHight {
		max = leftHight
		min = rightHight
	} else {
		max = rightHight
		min = leftHight
	}

	if max-min <= 1 {
		return max + 1, true
	} else {
		return max + 1, false
	}
}
