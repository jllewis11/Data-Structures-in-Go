package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestLinkedList(t *testing.T) {
	rand.NewSource(time.Now().UnixNano())
	list := &List{}
	tail := rand.Intn(100)
	head := rand.Intn(100)
	test := rand.Intn(100)
	list.Insert(tail)
	list.Insert(test)
	fmt.Println("==============================")
	fmt.Println("Test #1: Testing Insert()")
	if list.head.key != test {
		t.Errorf("Insert() failed. Expected %d, got %d", test, list.head.key)
	}
	list.Insert(3)
	list.Insert(4)
	list.Insert(head)
	list.Display()
	fmt.Println("==============================")
	fmt.Println("Test #1: Regular Linked List")
	if list.head.key != head {
		t.Errorf("head.key != %v", head)
		fmt.Println("Pass")
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
	rhead := tail
	rtail := head
	if list.head.key != rhead {
		t.Errorf("head.key != %v", rtail)
		if list.tail.key != rtail {
			t.Errorf("tail.key != %v", rhead)
		}
		list.Display()
	} else {
		fmt.Printf("Head: %v\n", list.head.key)
		fmt.Printf("Tail: %v\n", list.tail.key)
		list.Display()
	}
	fmt.Println("==============================")
}
