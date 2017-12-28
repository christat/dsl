// Package dsl is a (minimal) data structures library for Go.
// Implements several classic data structures such as single-linked lists,
// stacks and queues (both node and slice based versions).
package dsl

import (
	"errors"
	"math"
)

// Basic Node struct, basis of any single-linked list structure.
type Node struct {
	Data interface{}
	Next *Node
}

/*
NodeList is a an implementation of a singly linked list. It takes any interface{} and
allows:

- Retrieving: obtaining the value contained at any given index within the list.

- Appending: adding a new value at the last position of the list.

- Adding: adding a new value, specifying the index to be inserted at.

- Removing: deleting a value from the list, obtaining it if needed.

Note that the implementation is NOT thread-safe.
*/
type NodeList struct {
	Head *Node
	Tail *Node
	Size uint
}

// Internal function used to iterate through the list and retrieve a Node at the index value. Doesn't check for errors.
func (list *NodeList) getNode(index int) (node *Node) {
	node = list.Head
	for i := 0; i < index; i++ {
		node = node.Next
	}
	return node
}

// Internal function used to allow reverse search by subtracting the size of the list with the negative offset.
func (list *NodeList) reverseIndex(index int) (bool, int) {
	if index < 0 {
		return true, int(list.Size) - int(math.Abs(float64(index)))
	}
	return false, 0
}

// Retrieve obtains data stored at position index within the list. Returns the data or an error if failed.
func (list *NodeList) Retrieve(index int) (interface{}, error) {
	isReverse, value := list.reverseIndex(index)
	if isReverse {
		index = value
	}
	if index >= int(list.Size) || index < 0 {
		return nil, errors.New("cannot Retrieve() index out of bounds")
	}
	node := list.getNode(index)
	return node.Data, nil
}

// Append the data passed as parameter to the end of the list.
func (list *NodeList) Append(data interface{}) {
	node := &Node{Data: data, Next: nil}
	if list.Size > 1 {
		list.Tail.Next = node
	} else if list.Size == 0 {
		list.Head = node
	} else {
		list.Head.Next = node
	}
	list.Tail = node
	list.Size++
}

// Add the data passed as parameter at the position designed by index. Returns an error if out of bounds.
func (list *NodeList) Add(index int, data interface{}) error {
	isReverse, value := list.reverseIndex(index)
	if isReverse {
		index = value
	}
	if index > int(list.Size) || index < 0 {
		return errors.New("cannot Add() index out of bounds")
	}
	if index == int(list.Size) {
		list.Append(data)
		return nil
	}
	node := &Node{Data: data, Next: nil}
	if index == 0 {
		node.Next = list.Head
		list.Head = node
	} else {
		prev := list.getNode(index - 1)
		next := prev.Next
		prev.Next = node
		node.Next = next
	}
	list.Size++
	return nil
}

// Remove the item stored at position index in the list. Returns the extracted data or an error if out of bounds.
func (list *NodeList) Remove(index int) (interface{}, error) {
	isReverse, value := list.reverseIndex(index)
	if isReverse {
		index = value
	}
	if index >= int(list.Size) || index < 0 {
		return nil, errors.New("cannot Remove() index out of bounds")
	}
	var data interface{}
	if index == int(list.Size)-1 {
		data = list.Tail.Data
		prev := list.getNode(int(list.Size) - 2)
		prev.Next = nil
		list.Tail = prev
	} else if index > 0 {
		prev := list.getNode(index - 1)
		node := prev.Next
		data = node.Data
		prev.Next = node.Next
		node.Next = nil
	} else {
		data = list.Head.Data
		next := list.Head.Next
		list.Head.Next = nil
		list.Head = next
	}
	list.Size--
	return data, nil
}
