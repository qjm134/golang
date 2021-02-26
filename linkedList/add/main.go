/*
2. 两数相加
给你两个非空的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0开头。

输入：2 -> 4 -> 3
     5 -> 6 -> 4
输出：7 -> 0 -> 8

*/

package main

import "fmt"

type node struct {
	data int
	next *node
}

func add(h1, h2 *node) *node {
	head := &node{}
	flag := 0
	p := head
	for h1 != nil || h2 != nil || flag != 0 {
		x, y := 0, 0
		if h1 != nil {
			x = h1.data
			h1 = h1.next
		}
		if h2 != nil {
			y = h2.data
			h2 = h2.next
		}

		tmp := x + y + flag
		q := &node{}
		q.data = tmp % 10
		flag = tmp / 10

		p.next = q
		p = p.next
	}
	head = head.next
	return head
}

func printList(head *node) {
	for head != nil {
		fmt.Printf("%d -> ", head.data)
		head = head.next
	}
}

func main() {
	h1 := &node{2, nil}
	h1.next = &node{4, nil}
	h1.next.next = &node{3, nil}

	h2 := &node{5, nil}
	h2.next = &node{6, nil}
	h2.next.next = &node{4, nil}

	printList(add(h1, h2))
}
