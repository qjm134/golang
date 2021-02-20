/*
141. Linked List Cycle
https://leetcode.com/problems/linked-list-cycle/

Difficulty:
Easy

Description
Given a linked list, determine if it has a cycle in it.

To represent a cycle in the given linked list, we use an integer pos which represents the position (0-indexed) in the linked list where tail connects to. If pos is -1, then there is no cycle in the linked list.

Example 1: example-1

Input: head = [3,2,0,-4], pos = 1
Output: true
Explanation: There is a cycle in the linked list, where tail connects to the second node.
Example 2: example-2

Input: head = [1,2], pos = 0
Output: true
Explanation: There is a cycle in the linked list, where tail connects to the first node.
Example 3: example-3

Input: head = [1], pos = -1
Output: false
Explanation: There is no cycle in the linked list.
Follow up:

Can you solve it using O(1) (i.e. constant) memory?


思路：
初始2点相距c，速度分别是v1、v2，2点重合的时候，相遇在d，分别走了m、n圈，设周长是l
v1*t = m*l + d
v2*t = n*l + (d-c)

v1/v2 = (m*l +d)/(n*l + (d-c))
m*l + d = k*n*l + kd - kc
m*l = k*n*l + (k-1)*d - kc
c、k任意固定一个值，如果m!=k*n，即非平行的2条线，则总会相交一点，得到m,n,d的一组解
如果 m == kn，则 d = kc/(k-1)
*/

package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	var fast, slow *ListNode
	if head == nil {
		return false
	}
	if head.Next == nil {
		return false
	}

	slow = head
	fast = head.Next.Next
	for fast != nil {
		if slow == fast {
			return true
		}
		if fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}
