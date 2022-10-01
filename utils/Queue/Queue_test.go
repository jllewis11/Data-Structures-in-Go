package Queue

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestQueue(t *testing.T) {
	rand.NewSource(time.Now().UnixNano())
	queue := &Queue{}
	tail := rand.Intn(100)
	head := rand.Intn(100)
	test := rand.Intn(100)
	queue.push(tail)
	queue.push(test)
	fmt.Println("==============================")
	fmt.Println("Test #1: Testing push()")
	if queue.head.key != test {
		t.Errorf("push() failed. Expected %d, got %d", test, queue.head.key)
	}
	queue.push(rand.Intn(100))
	queue.push(rand.Intn(100))
	queue.push(head)
	queue.Display()
	fmt.Println("==============================")
	fmt.Println("Test #2: Regular Queue")
	if queue.head.key != test {
		t.Errorf("head.key != %v", test)
		if queue.tail.key != head {
			t.Errorf("tail.key != %v", head)
		}
	} else {
		fmt.Printf("Head: %v", queue.head.key)
		fmt.Printf("Tail: %v", queue.tail.key)
		queue.Display()
	}
	fmt.Println("==============================")
	fmt.Println("Test #3: Testing pop()")
	queue.pop()
	queue.Display()
	//Check if the new node is inserted after the specified key
	if queue.head.key != tail {
		t.Errorf("pop() failed. Expected %d, got %d", tail, queue.head.key)
	}
}
