package main

import(
	"fmt"
	"errors"
)

type ITree interface {
	Add(value int)
	Search(value int) bool
	Min() int
	Max() int
	PrintPre()
	PrintIn()
	PrintPos()
	PrintLevels()
	Height() int
	Remove(value int) *BstNode
}

type Node struct {
	val int
	left *Node
	right *Node
}

type BST struct {
	root *Node
	inserted int
}