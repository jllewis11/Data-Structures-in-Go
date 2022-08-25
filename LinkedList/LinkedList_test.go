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
	list.Insert(0)
	list.Insert(1)
	list.Insert(rand.Intn(100))
	list.Insert(rand.Intn(100))
	list.Insert(6)
	list.Display()
	fmt.Println("==============================")
	fmt.Println("Test #1: Regular Linked List")
	fmt.Println("Tail is ", list.tail.key)
	fmt.Println("Tail", tail)
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
	if list.head.key != head {
		t.Errorf("head.key != %v", tail)
		if list.tail.key != tail {
			t.Errorf("tail.key != %v", head)
		}
		list.Display()
	} else {
		fmt.Printf("Head: %v\n", list.head.key)
		fmt.Printf("Tail: %v\n", list.tail.key)
		list.Display()
	}
	fmt.Println("==============================")
}
