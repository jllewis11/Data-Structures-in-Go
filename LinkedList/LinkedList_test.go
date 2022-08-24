package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestLinkedList(t *testing.T) {
	list := &List{}
	head := rand.Intn(100)
	tail := rand.Intn(100)
	list.Insert(head)
	list.Insert(rand.Intn(100))
	list.Insert(rand.Intn(100))
	list.Insert(rand.Intn(100))
	list.Insert(tail)

	fmt.Println("==============================")
	fmt.Println("Test #1: Regular Linked List")
	if list.head.key != head {
		t.Errorf("head.key != %v", head)
		if list.tail.key != tail {
			t.Errorf("tail.key != %v", tail)
		}

	} else {
		fmt.Printf("Head: %v\n", list.head.key)
		fmt.Printf("Tail: %v\n", list.tail.key)
		list.Display()
	}
	list.Reverse()

	fmt.Println("==============================")
	fmt.Println("Test #2: Reversed Linked List")
	// Head and tail values will be reversed.
	if list.head.key != head {
		t.Errorf("head.key != %v", tail)
		if list.tail.key != tail {
			t.Errorf("tail.key != %v", head)
		}
		list.Display()
	} else {
		fmt.Printf("Head: %v\n", list.tail.key)
		fmt.Printf("Tail: %v\n", list.head.key)
		list.Display()
	}
	fmt.Println("==============================")
}
