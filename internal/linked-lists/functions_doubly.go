package linkedlists

import "fmt"

// functions for the doubly linked list

func (dll *DoublyLinkedList) Append (n *DoublyNode){
	// if head is nil, list is empty, set node as head & tail
	if dll.Head == nil {
		dll.Head = n
		dll.Tail = n
	} else {
		// new tail node has the previous tail as previous node
		n.Previous = dll.Tail
		
		// the previous tail has the new tail as next node
		n.Previous.Next = n
		
		// the new tail of the dll is the new node
		dll.Tail = n
	}

	dll.Length++
}

func (dll *DoublyLinkedList) PrintAll() error {
	fmt.Printf("---------\n")
	if (dll.Length == 0){
		return fmt.Errorf("Doubly List is empty.")
	}
	current := dll.Head
	for i := 0; i < dll.Length; i++ {
		fmt.Println(current.Value)
		current = current.Next
	}

	return nil
}