package main

import (
	"fmt"

	"../../queue"
)

func main() {
	root := new(node)
	root.lchild = new(node)
	root.rchild = new(node)
	root.rchild.lchild = new(node)
	root.rchild.lchild.rchild = new(node)
	fmt.Println(recursiveDepth(root))
	fmt.Println(nonRecursiveDepth(root))
}

type node struct {
	lchild *node
	rchild *node
}

/*
跟fibonacci楼梯一样，从上向下看，根节点的左右子树，哪个子树的层数高，取哪个，最后加上根节点这一层，即是树高
*/
func recursiveDepth(root *node) int {
	if root == nil {
		return 0
	}

	ldepth := recursiveDepth(root.lchild)
	rdepth := recursiveDepth(root.rchild)

	if ldepth > rdepth {
		return ldepth + 1
	} else {
		return rdepth + 1
	}
}

/*
按层遍历树
上层节点依次出队
出队的时候，其左右子节点依次入队
用一个数记录下层节点数
当前层节点数遍历完时，下一层节点数更新当前层数，下一层节点数清零
初始，将根节点放入队列，当前节点数置1，只要队列中有，则出队
*/
func nonRecursiveDepth(root *node) int {
	que := queue.New()
	if root != nil {
		que.Put(root)
	}
	depth := 0
	currentNum := 1
	nextNum := 0
	for ni := que.Get(); ni != nil; ni = que.Get() {
		n := ni.(*node)
		if n.lchild != nil {
			que.Put(n.lchild)
			nextNum++
		}
		if n.rchild != nil {
			que.Put(n.rchild)
			nextNum++
		}
		currentNum--
		if currentNum == 0 {
			currentNum = nextNum
			nextNum = 0
			depth++
		}
	}
	return depth
}
