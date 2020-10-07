// Package dsgo linked list implementation
package dsgo

// List container
type List struct {
	head   *Node
	length int
}

// Node of linked list
type Node struct {
	Value  interface{}
	Next *Node
}

// New initializes a linked list
func New() List {
	return List{}
}

// Get the value of the index-th node in the linked list.
// Return -1 if the index is invalid. List is zero-based
func (l *List) Get(index int) interface{} {
	if index < 0 || index > l.length-1 {
		return -1
	}

	currentNode := l.head
	for i := 0; i < index && currentNode != nil; i++ {
		currentNode = currentNode.Next
	}

	if currentNode == nil {
		return -1
	}

	return currentNode.Value
}

// AddAtHead a node before the first element of the linked list.
// After the insertion, the new node will be the first node of the linked list
func (l *List) AddAtHead(val interface{}) {
	var newNode Node
	newNode.Value = val
	newNode.Next = l.head
	l.head = &newNode
	l.length++
}

// AddAtTail a node to the last element of the linked list
func (l *List) AddAtTail(val interface{}) {
	if l.head == nil {
		l.AddAtHead(val)
		return
	}

	currentNode := l.head
	for i := 0; currentNode.Next != nil; i++ {
		currentNode = currentNode.Next
	}

	var newNode Node
	newNode.Value = val
	currentNode.Next = &newNode

	l.length++
}

// AddAtIndex a node before the index-th node in the linked list.
// If index equals to the length of linked list, the node will be appended
// to the end of linked list. If index is greater than the length, the node
// will not be inserted
func (l *List) AddAtIndex(index int, val interface{}) {
	if index == 0 {
		l.AddAtHead(val)
		return
	} else if index == l.length {
		l.AddAtTail(val)
		return
	} else if index < 0 || index > l.length {
		return
	}

	currentNode := l.head
	for i := 0; i < index-1; i++ {
		currentNode = currentNode.Next
	}

	var newNode Node
	newNode.Value = val
	newNode.Next = currentNode.Next
	currentNode.Next = &newNode

	l.length++
}

// DeleteAtIndex a node in the linked list, if the index is valid
func (l *List) DeleteAtIndex(index int) {
	if index < 0 || index >= l.length {
		return
	} else if index == 0 {
		l.head = l.head.Next
		l.length--
		return
	}

	currentNode := l.head
	for i := 0; i < index-1; i++ {
		currentNode = currentNode.Next
	}

	currentNode.Next = currentNode.Next.Next
	l.length--
}
