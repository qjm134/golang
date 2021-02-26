package main

import "fmt"

type node struct {
	data int
	next *node
}

func reverseLinkList(head *node) *node {
	if head == nil {
		return nil
	}

	pre := head
	cur := pre.next
	for cur != nil {
		next := cur.next
		cur.next = pre
		pre = cur
		cur = next
	}
	head.next = nil

	return pre
}

//循环链表逆转
//in: 1->2->3->1
//out: 1->3->2->1
func reverseCycleList(head *node) *node {
	if head == nil {
		return head
	}

	pre := head
	cur := pre.next
	for cur != head {
		next := cur.next
		cur.next = pre
		pre = cur
		cur = next
	}
	head.next = pre
	return head
}

func main() {
	headl := &node{1, nil}
	headl.next = &node{2, nil}
	headl.next.next = &node{3, nil}
	printList(headl)
	printList(reverseLinkList(headl))

	head := &node{1, nil}
	head.next = &node{2, nil}
	head.next.next = &node{3, head}
	printCycleList(head)
	printCycleList(reverseCycleList(head))
}

func printList(head *node) {
	for head != nil {
		fmt.Printf("%d -> ", head.data)
		head = head.next
	}
	fmt.Println()
}

func printCycleList(head *node) {
	if head == nil {
		return
	}
	fmt.Printf("%d -> ", head.data)
	p := head.next
	for p != head {
		fmt.Printf("%d -> ", p.data)
		p = p.next
	}
	fmt.Println(p.data)
}
