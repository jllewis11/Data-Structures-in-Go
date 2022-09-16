package main

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

func (L *List) append(key interface{}) {
	list := &Node{
		next: L.head,
		key:  key,
	}
	if L.head != nil {
		L.head.prev = list
	}
	L.head = list
	if L.tail == nil {
		L.tail = list
	}
}

<<<<<<< HEAD
// Insert after a specific key interface{}
=======
func (L *List) prepend(key interface{}) {
	//If the list is not empty, insert at the beginning of the list.
	if L.head != nil {
		list := &Node{
			next: L.head,
			key:  key,
		}
		L.head.prev = list
		L.head = list
	} else {
		L.head = &Node{
			key: key,
		}
		L.tail = L.head
	}
}

//Insert after a specific key interface{}
>>>>>>> 0bd37a5ff4179e391bc946093bb37139e2a799ac
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
	//If the key is not found, insert at the end of the list.
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
