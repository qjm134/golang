package solution

type Node struct {
	Data int
	Next *Node
}

func FindEntrance(head *Node) (*Node, bool) {
	if head == nil {
		return nil, false
	}
	slow := head
	fast := head
	for slow.Next != nil {
		slow = slow.Next
		fast = slow.Next.Next
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
