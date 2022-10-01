package LinkedList

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

//Append
//InsertAfter
//Reverse
//Display
//prepend

func TestLinkedList(t *testing.T) {
	rand.NewSource(time.Now().UnixNano())
	list := &List{}
	tail := rand.Intn(100)
	head := rand.Intn(100)
	test := rand.Intn(100)
	list.append(tail)
	list.append(test)

	fmt.Println("==============================")
	fmt.Println("Test #1: Testing append()")
	if list.head.key != test {
		t.Errorf("append() failed. Expected %d, got %d", test, list.head.key)
	}

	list.append(rand.Intn(100))
	list.append(rand.Intn(100))
	list.append(head)
	list.Display()

	fmt.Println("==============================")
	fmt.Println("Test #2: Regular Linked List")
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

	fmt.Println("==============================")
	fmt.Println("Test #3: Testing prepend()")

	r1 := rand.Intn(100)
	list.prepend(r1)

	if list.head.key != r1 {
		t.Errorf("prepend() failed. Expected %d, got %d", r1, list.head.key)
	}
	list.Display()
	fmt.Println("==============================")
	fmt.Println("Test #4: Testing InsertAfter()")

	rand1 := rand.Intn(100)
	list.InsertAfter(r1, rand1)
	list.Display()
	//Check if the new node is inserted after the specified key
	if list.head.next.key != rand1 {
		t.Errorf("InsertAfter() failed. Expected %d, got %d", head, list.head.next.key)
	}
	fmt.Println("")
	fmt.Println("Test #5: Testing InsertAfter() if key is not found")
	rand2 := rand.Intn(100)
	//If the key is not found, insert at the end of the list. 101 is not in the list.
	list.InsertAfter(101, rand2)
	list.Display()
	//Check if the new node is inserted after the specified key
	if list.tail.key == rand2 {
		t.Errorf("InsertAfter() failed. Expected %d, got %d", head, list.head.next.key)
	}

	fmt.Println("==============================")
	fmt.Println("Test #6: Reversed Linked List")
	list.Reverse()
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
