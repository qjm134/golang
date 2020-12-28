/*
21. Merge Two Sorted Lists
https://leetcode.com/problems/merge-two-sorted-lists/

Description
Merge two sorted linked lists and return it as a new list.

The new list should be made by splicing together the nodes of the first two lists.

Example:

Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
*/
package mergeSortedList

type node struct {
	Val  int
	Next *node
}

func merge(l1, l2 *node) *node {
	var h, p *node
	h = &node{}
	p = h

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}

	if l1 == nil {
		p.Next = l2
	}
	if l2 == nil {
		p.Next = l1
	}
	h = h.Next
	return h
}
