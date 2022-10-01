package Queue

type Node struct {
	next *Node
	key  interface{}
}

type Queue struct {
	head *Node
	tail *Node
}

func (q *Queue) push(key interface{}) {
	node := &Node{key: key}
	if q.head == nil {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		q.tail = node
	}
}

func (q *Queue) pop() interface{} {
	if q.head == nil {
		return nil
	}
	node := q.head
	q.head = q.head.next
	return node.key
}

func (q *Queue) Display() {
	curr := q.head
	for curr != nil {
		println(curr.key)
		curr = curr.next
	}
}
