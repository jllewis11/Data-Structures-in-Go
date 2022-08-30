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

func (L *List) Insert(key interface{}) {
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

//Insert after a specific key interface{}
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
	l.Insert(newKey)
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

func ShowBackwards(list *Node) {
	for list != nil {
		fmt.Printf("%v <-", list.key)
		list = list.prev
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
