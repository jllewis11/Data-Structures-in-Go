package main

//Node Creation
type Node struct {
	prev *Node
	next *Node
	key  interface{}
}

type LinkedList struct {
	head *Node
	tail *Node
}
