/*
一个链表，奇数位升序，偶数位降序，全部升序重排
*/
package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

func reSort(h *node) *node {
	if h == nil || h.next == nil {
		return h
	}

	odd := h
	even := h.next
	o := odd
	e := even
	isOdd := true

	h = h.next.next
	for h != nil {
		if isOdd {
			o.next = h
			o = o.next
			isOdd = false
		} else {
			e.next = h
			e = e.next
			isOdd = true
		}
		h = h.next
	}
	if isOdd {
		o.next = nil
	} else {
		e.next = nil
	}

	re := reverse(even)
	return merge(odd, re)
}

func reverse(h *node) *node {
	if h == nil {
		return h
	}

	p := h
	q := p.next
	h.next = nil
	for q != nil {
		r := q.next
		q.next = p
		p = q
		q = r
	}
	return p
}

func merge(h1, h2 *node) *node {
	h := &node{}
	p := h

	for h1 != nil && h2 != nil {
		if h1.data <= h2.data {
			p.next = h1
			h1 = h1.next
		} else {
			p.next = h2
			h2 = h2.next
		}
		p = p.next
	}

	if h1 == nil {
		p.next = h2
	}
	if h2 == nil {
		p.next = h1
	}

	return h.next
}

func main() {
	h := &node{2, nil}
	p := h
	p.next = &node{5, nil}
	p = p.next
	p.next = &node{3, nil}
	p = p.next
	p.next = &node{1, nil}
	p = p.next
	p.next = &node{4, nil}

	printl(h)
	printl(reSort(h))
}

func printl(h *node) {
	for h != nil {
		fmt.Printf("%v->", h.data)
		h = h.next
	}
	fmt.Println()
}
