package main

func main() {

}

type node struct {
	lchild *node
	rchild *node
}

func depth(root *node) int {
	if root == nil {
		return 0
	}

	ldepth := depth(root.lchild)
	rdepth := depth(root.rchild)

	if ldepth > rdepth {
		return ldepth + 1
	} else {
		return rdepth + 1
	}
}
