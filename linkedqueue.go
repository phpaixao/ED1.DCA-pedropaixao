package main

import (
	"errors"
	"fmt"
)

type IArrayQueue interface {
	Enqueue(val int)
	Dequeue() (int, error)
	Size() int
	Front() (int, error)
}

type Node struct {
	value int
	next *Node
}

type LinkedQueue struct {
	front *Node
	rear *Node
	inserted int
}