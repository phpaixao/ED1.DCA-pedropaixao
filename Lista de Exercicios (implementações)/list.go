package main

import (
	"fmt"
	"errors"
)

type IList interface {
	Add(value int)							//Feito
	AddOnIndex(value int, index int) error	
	RemoveOnIndex(index int) error
	Get(index int) (int, error)
	Set(value int, index int) error
	Size() int
}

type ArrayList struct{
	v [] int
	inserted int
}

type Node1 struct{
	value int
	next *Node1
}

type Node2 struct{
	value int
	next *Node2
	prev *Node2
}

type LinkedList struct{
	head *Node1
	inserted int
}

type DoubleLinkedList struct{
	head *Node2
	tail *Node2
	inserted int
}

func (list *ArrayList) Init(size int){
	list.v = make([] int, size)
}

func (list *ArrayList) doubleV(){
	newV := make([] int, len(list.v)*2)
	for i:=0; i<len(list.v); i++{
		newV[i] = list.v[i]
	}
	list.v = newV
}

func (list *ArrayList) Add(value int){
	if list.Size() == len(list.v){
		list.doubleV()
	}
	list.v[list.inserted] = value
	list.inserted++
}

func (list *LinkedList) Add(value int){
	new_node := &Node1{
		value: value,
		next: nil,
	}
	if list.Size() == 0{
		list.head = new_node
	} else {
		aux := list.head
		for i:=0; i<list.Size()-1; i++{
			aux = aux.next
		}
		aux.next = new_node
	}
	list.inserte
	d++
}

func (list *DoubleLinkedList) Add(value int){
	new_node := &Node2{
		value: value,
		next: nil,
		prev: nil,
	}
	if list.Size() == 0{
		list.head, list.tail = new_node, new_node
	} else {
		new_node.prev = list.tail
		list.tail.next = new_node
		list.tail = new_node
	}
	list.inserted++
}