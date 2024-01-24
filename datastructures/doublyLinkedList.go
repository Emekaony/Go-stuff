package datastructures

import "errors"

type Node[T any] struct {
	value T
	prev  *Node[T]
	next  *Node[T]
}

type DoublyLinkedList[T any] struct {
	size       int
	head, tail *Node[T]
}

// add an element into the LinkedList. Index is zero-based!
func (l *DoublyLinkedList[T]) Add(value T, index int) error {
	currentSize := l.size
	/*
		there are 3 edge cases:
			1). Adding into an empty list.
			2). Adding to the head.
			3). Adding to the tail.
	*/
	newNode := &Node[T]{
		value: value,
	}
	l.size = currentSize + 1 // just increment the size of the list since we will add anyway
	// handle obvious case when index > size of list
	if index > currentSize {
		return errors.New("Index exceeds list size.")
	}

	// handle empty list case
	if l.head == nil {
		// I for sure forgot to handle the tail scenario. NEVER make this mistake again
		l.head, l.tail = newNode, newNode
		return nil
	}

	// handle adding to the head.
	if index == 0 {
		newNode.next = l.head
		l.head.prev, l.head = newNode, newNode
		return nil
	}

	// handle addition to the tail. This is when the zer-baded index is equal to the currentSize.
	if index == currentSize {
		newNode.prev = l.tail
		l.tail.next, l.tail = newNode, newNode
		return nil
	}

	// handle the trivial case of adding in the middle
	current := l.head
	for i := 1; i < index; i++ {
		current = current.next
	}

	newNode.next = current.next
	newNode.prev = current
	current.next.prev, current.next = newNode, newNode

	return nil
}
