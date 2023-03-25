package linkedlists

import "fmt"

// Append adds a node to the end of the linked list.
func (l *LinkedList) Append(n *Node) {
	if l.Head == nil {
		l.Head = n
		l.Tail = n
	} else {
		l.Tail.Next = n
		l.Tail = n
	}
	l.Length++
}

// Prepend adds a node to the beginning of the linked list.
func (l *LinkedList) Prepend(n *Node) {
	if l.Head == nil {
		l.Head = n
		l.Tail = n
	} else {
		n.Next = l.Head
		l.Head = n
	}
	l.Length++
}

// Find returns the first node with the given value.
func (l *LinkedList) Find(v int) *Node {
	current := l.Head
	for current != nil {
		if current.Value == v {
			return current
		} else {
			current = current.Next
		}
	}
	return nil
}

// Delete removes the first node with the given value.
func (l *LinkedList) DeleteValue(v int) {
	current := l.Head

	for current != nil {
		if current.Value == v {
			current.Value = current.Next.Value
			current.Next = current.Next.Next
			l.Length--
		}
	}
}


// Remove index removes the node at the given index
func (l *LinkedList) RemoveIndex(index int) {
	// start at beginning
	leading_node := l.Head
	
	// get the node before target index
	for i := 0; i < index - 1; i++ {
		leading_node = leading_node.Next
	}
	
	// get the node after the target index
	following_node := leading_node.Next.Next
	
	// point the leading node to following (instead of to target)
	leading_node.Next = following_node

	l.Length--
}

// Insert insert a node based on givin index.
func (l *LinkedList) Insert(index int, n *Node) error {
	if index < 0 || index > l.Length {
		return fmt.Errorf("index %d out of bounds for linked list of length %d", index, l.Length)
	}

	if index == 0 {
		l.Prepend(n)
		return nil
	}

	if index == l.Length {
		l.Append(n)
		return nil
	}

	currentNode := l.Head
	for i := 0; i < index-1; i++ {
		currentNode = currentNode.Next
	}

	n.Next = currentNode.Next
	currentNode.Next = n
	l.Length++

	return nil
}


func (l *LinkedList) PrintAll() {
	fmt.Printf("---------\n")
	
	// start at first node
	current := l.Head

	for i := 0; i < l.Length -1; i++{
		fmt.Println(current.Value)
		current = current.Next
	}

}

