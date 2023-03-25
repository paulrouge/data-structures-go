package linkedlists

// Node represents a node in a linked list.
type Node struct {
	Value int
	Next  *Node
}

// DoublyNode represent a node in a doubly linked list
type DoublyNode struct {
	Value int
	Next *DoublyNode
	Previous *DoublyNode
}