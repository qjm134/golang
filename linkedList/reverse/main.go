package main

import "fmt"

func main() {
	list := new(linkedList)
	list.Put(1)
	list.Put(2)
	list.Put(3)

	list.Print()
	list.ReverseLinkedList()
	list.Print()
}

type linkedList struct {
	head *node
}

type node struct {
	data interface{}
	next *node
}

func (list *linkedList) ReverseLinkedList() {
	head := list.head
	if head == nil {
		return
	}

	var p, q, r *node
	if head.next != nil {
		p = head
		q = p.next
		p.next = nil
	} else {
		return
	}

	for q.next != nil {
		r = q.next
		q.next = p
		p = q
		q = r
	}
	q.next = p

	list.head = q
}

func (list *linkedList) Put(data interface{}) {
	n := new(node)
	n.data = data

	if list.head == nil {
		list.head = n
		return
	}

	p := list.head
	for p.next != nil {
		p = p.next
	}

	p.next = n
}

func (list *linkedList) Print() {
	p := list.head
	for p != nil {
		fmt.Println(p.data)
		p = p.next
	}
}
