package main

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

func FindEntrance(head *Node) (*Node, bool) {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			slow = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow, true
		}
	}
	return nil, false
}

func main() {
	head := &Node{1, nil}
	head.Next = &Node{2, nil}
	head.Next.Next = &Node{3, nil}
	head.Next.Next.Next = head.Next

	fmt.Println(FindEntrance(head))
}
