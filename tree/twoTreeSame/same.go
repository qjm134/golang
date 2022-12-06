package same

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   *int
}

func IsSame(p, q *TreeNode) bool {
	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}
	if p == nil && q == nil {
		return true
	}
	if p.Val != q.Val {
		return false
	}
	if !IsSame(p.Left, q.Left) {
		return false
	}
	if !IsSame(p.Right, q.Right) {
		return false
	}
	return true
}
