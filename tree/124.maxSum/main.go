/*
124. 二叉树中的最大路径和
路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。

路径和 是路径中各节点值的总和。

给你一个二叉树的根节点 root ，返回其最大路径和。


思路：


以下想法不对
对于一个根节点，2个左右节点的简单树，和最大的有2种可能：
左右节点比较出一个最大的
1. 如果根节点大于0，再加上根节点，即是这个树最大的
2. 如果根节点小于0，则左右节点中最大的，即是这个树最大的，不加根节点

对于第二种情况，再往上组一层时，
     5
   -3  2
 8
最大有3种情况：
1、2. 上面第一种加上根节点的情况，重复上面1、2
3. 上面第二种不加根节点的情况
这3种比较出一个最大的即是

*/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var preMax = -10000

func maxSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftMax := maxSum(root.Left)
	rightMax := maxSum(root.Right)

	var childMax int
	if leftMax >= rightMax {
		childMax = leftMax
	} else {
		childMax = rightMax
	}

	var parentMax int
	if root.Val > 0 {
		parentMax = childMax + root.Val
	} else {
		parentMax = childMax
	}

	var max int
	if parentMax >= preMax {
		max = parentMax
		preMax = parentMax
	} else {
		max = preMax
	}

	return max
}
