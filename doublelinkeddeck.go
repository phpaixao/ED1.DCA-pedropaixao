package main

import (
	"errors"
	"fmt"
)

type iDeck interface {
	push_front(val int)
	push_back(val int)
	pull_front() (int, error)
	pull_back() (int, error)
	size() int
	Front() (int, error)
	back() (int, error)
}

type Node struct{
	value int
	next *Node
	prev *Node
}

type DoubleLinkedQueue struct {
	head *Node
	tail *Node
	inserted int
}

func (deck *DoubleLinkedQueue) push_front(val int) {
	new_node := &Node{
		value: val,
		next: nil,
		prev: nil,
	}
	if deck.size() == 0{
		deck.head = new_node
		deck.tail = new_node
	} else {
		new_node.next = deck.head
		deck.head.prev = new_node
		deck.head = new_node
	}
	deck.inserted++
}

func (deck *DoubleLinkedQueue) push_back(val int) {
	new_node := &Node{
		value: val,
		next: nil
		prev: nil,
	}
	if deck.size() == 0{
		deck.head = new_node
		deck.tail = new_node
	} else {
		new_node.prev = deck.tail
		deck.tail.next = new_node
		deck.tail = new_node
	}
	deck.inserted++
}

func (deck *DoubleLinkedQueue) pull_front() (int, error){
	if deck.size() == 0{
		return -1, errors.New("Fila vazia.")
	}
}