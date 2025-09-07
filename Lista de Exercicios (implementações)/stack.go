package main

import (
	"fmt"
	"errors"
)

type IStack interface{
	Push(value int)
	Pop() (int, error)
	Peek() (int, error)
	IsEmpty() bool
	Size() int
}

type ArrayStack struct{
	v [] int
	top int
}

type Node struct{
	value int
	next *Node
}

type LinkedStack struct{
	top *Node
	inserted int
}

func (stack *ArrayStack) Init(size int){
	stack.v = make([] int, size)
}

func (stack *ArrayStack) doubleV(){
	newV := make([] int, stack.Size()*2)
	for i:=0; i<stack.Size(); i++{
		newV[i] = stack.v[i]
	}
	stack.v = newV
}