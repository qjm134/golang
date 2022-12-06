/*
101. 对称二叉树
给你一个二叉树的根节点 root ， 检查它是否轴对称。

1
2 2
3 4 4 3
null null 5 null null 5 null null
输入：root = [1,2,2,3,4,4,3]
输出：true

1
2 2
null 3 null 3
输入：root = [1,2,2,null,3,null,3]
输出：false

思路一
一棵树轴对称，2棵镜像对折相等的数，其左子树跟另一个右子树相等，右子树跟左子树相等

思路二
按层遍历，出队列一个，其左右节点依次入队列，记录每层是节点数
判断每层的数组是否是对称的
遇到空，左右节点都放空

思路三
按层，出队列2个，左节点和另一个右节点先入，再右节点和另一个左节点入
*/
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isTwoSymmetric(root.Left, root.Right)
}

func isTwoSymmetric(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	if res := isTwoSymmetric(p.Left, q.Right); !res {
		return false
	}

	if res := isTwoSymmetric(p.Right, q.Left); !res {
		return false
	}

	return true
}

func main() {
	root := &TreeNode{}
	root.Val = 1
	left := &TreeNode{}
	left.Val = 0
	root.Left = left

	fmt.Println(isSymmetric(root))
}
