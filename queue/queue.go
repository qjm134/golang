package queue

type Queue struct {
	front *node
}

type node struct {
	data interface{}
	next *node
}

func New() Queue {
	front := new(node)
	return Queue{front}
}

func (que *Queue) Put(data interface{}) {
	p := que.front
	for p.next != nil {
		p = p.next
	}
	n := new(node)
	n.data = data
	p.next = n
}

func (que *Queue) Get() (data interface{}) {
	if que.front.next == nil {
		return nil
	}

	p := que.front.next
	que.front.next = p.next

	return p.data
}
