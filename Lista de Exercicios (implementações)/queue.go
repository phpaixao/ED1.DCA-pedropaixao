package main

import (
	"errors"
	"fmt"
)

type IQueue interface {
	Enqueue(value int)
	Dequeue() (int, error)
	Front() (int, error)
	IsEmpty() bool
	Size() int
}

type ArrayQueue struct{
	front int
	rear int
	v[] int
}

type Node struct{
	value int
	next *Node
}

type LinkedListQueue struct{
	front *Node
	rear *Node
	inserted int
}