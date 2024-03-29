package linkedlists

import "fmt"

// functions for the doubly linked list

func (dll *DoublyLinkedList) Append(n *DoublyNode) {
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

func (dll *DoublyLinkedList) Prepend(n *DoublyNode) {
	// if head is nil this is the first node in the list so tail and head are this node
	if dll.Head == nil {
		dll.Head = n
		dll.Tail = n
	} else {
		// new node is previous node of current head
		dll.Head.Previous = n

		// current head is the next node of new node
		n.Next = dll.Head

		// new node is new head
		dll.Head = n
	}

	dll.Length++
}

func (dll *DoublyLinkedList) Insert(index int, n *DoublyNode) error {
	if index == 0 {
		dll.Prepend(n)
		return nil
	} else if index >= dll.Length {
		dll.Append(n)
		return nil
	} else {

		current := dll.Head

		// get node at given index
		for i := 0; i < index; i++ {
			current = current.Next
		}

		// set Next op new node
		n.Next = current

		// set Previous of new node
		n.Previous = current.Previous

		// point leader of selected node to new node
		current.Previous.Next = n

		dll.Length++

		return nil
	}
}

func (dll *DoublyLinkedList) PrintAll() error {
	fmt.Printf("---------\n")

	if dll.Length == 0 {
		return fmt.Errorf("Doubly List is empty.")
	}

	current := dll.Head

	for i := 0; i <= dll.Length-1; i++ {
		fmt.Println(current.Value)
		current = current.Next
	}

	return nil
}
