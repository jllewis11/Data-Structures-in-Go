package Queue

import "fmt"

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
	key := q.head.key
	q.head = q.tail
	return key
}
func (q *Queue) Display() {
	node := q.head
	for node != nil {
		fmt.Println(node.key)
		node = node.next
	}
}
