/*
236. 二叉树的最近公共祖先
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

思路一
最近公共祖先的特点总结为2点：
1.这2个节点，分散在祖先节点的2个子树上
2.其中一个节点是公共祖先

思路二
存储父节点
找到某个节点的所有父节点：可以遍历树用map保存所有节点的所有父节点
最近公共父节点：
某个节点开始往回遍历父节点，遍历过的标记。相当从某个点开始画了一条路
从另一个节点开始也往回遍历父节点，第一个遇到已经被标记过的节点，即最近公共父节点

*/
package main

//**
//* Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	parentMap := make(map[*TreeNode]*TreeNode)
	findParent(root, parentMap)

	flagMap := make(map[*TreeNode]bool)
	pParent := p
	for pParent != nil {
		flagMap[pParent] = true
		pParent = parentMap[pParent]
	}
	qParent := q
	for qParent != nil {
		if flagMap[qParent] == true {
			return qParent
		}
		qParent = parentMap[qParent]
	}

	return root
}

func findParent(root *TreeNode, parentMap map[*TreeNode]*TreeNode) {
	if root == nil {
		return
	}

	if root.Left != nil {
		parentMap[root.Left] = root
		findParent(root.Left, parentMap)
	}
	if root.Right != nil {
		parentMap[root.Right] = root
		findParent(root.Right, parentMap)
	}

	return
}

func forLowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	childMap := make(map[*TreeNode][]*TreeNode)
	parentList := make([]*TreeNode, 0)
	findParentList(root, childMap, parentList)

	pParant := childMap[p]
	qParent := childMap[q]

	for i := len(pParant) - 1; i >= 0; i-- {
		for j := len(qParent) - 1; j >= 0; j-- {
			if pParant[i] == qParent[j] {
				return pParant[i]
			}
		}
	}

	return nil
}

func findParentList(root *TreeNode, childMap map[*TreeNode][]*TreeNode, parentList []*TreeNode) {
	if root.Left != nil {
		childParentList := []*TreeNode{root}
		childMap[root.Left] = append(childParentList, parentList...)
		findParentList(root.Left, childMap, childParentList)
	}

	if root.Right != nil {
		childParentList := []*TreeNode{root}
		childMap[root.Right] = append(childParentList, parentList...)
		findParentList(root.Right, childMap, childParentList)
	}

	return
}
