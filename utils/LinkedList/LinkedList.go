package LinkedList

import "fmt"

type Node struct {
	prev *Node
	next *Node
	key  interface{}
}

type List struct {
	head *Node
	tail *Node
}

func (l *List) append(key interface{}) {
	list := &Node{
		next: l.head,
		key:  key,
	}
	if l.head != nil {
		l.head.prev = list
	}
	l.head = list
	if l.tail == nil {
		l.tail = list
	}
}

func (l *List) prepend(key interface{}) {
	list := &Node{
		next: l.head,
		key:  key,
	}
	if l.head == nil {
		l.head = list
		l.tail = list
	} else {
		list.next = l.head
		l.head = list
	}
}

func (l *List) InsertAfter(key interface{}, newKey interface{}) {
	curr := l.head
	for curr != nil {
		if curr.key == key {
			list := &Node{
				prev: curr,
				next: curr.next,
				key:  newKey,
			}
			curr.next = list
			if list.next != nil {
				list.next.prev = list
			} else {
				l.tail = list
			}
			return
		}
		curr = curr.next
	}
	l.append(newKey)
}

func (l *List) Display() {
	list := l.head
	for list != nil {
		fmt.Printf("%+v ->", list.key)
		list = list.next
	}
	fmt.Println()
}

func Display(list *Node) {
	for list != nil {
		fmt.Printf("%v ->", list.key)
		list = list.next
	}
	fmt.Println()
}

func (l *List) Reverse() {
	curr := l.head
	var prev *Node
	l.tail = l.head

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head = prev
}
