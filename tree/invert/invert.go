package invert

type TreeNode struct {
	data  int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) {
	if root == nil {
		return
	}

	if root.Left != nil || root.Right != nil {
		root.Left, root.Right = root.Right, root.Left
	}

	invertTree(root.Left)
	invertTree(root.Right)

	return
}
