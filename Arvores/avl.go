package main

import(
	"fmt"
	"errors"
)

type Node struct {
	left *Node
	val int
	right *Node
	bf int
	height int
}

func (root *Node) RotRight() *Node {
	newRoot := root.left
	root.left = newRoot.right
	newRoot.right = root
	root.UpdateProperties()
	newRoot.UpdateProperties()
	return newRoot
}

func (root *Node) RotLeft() *Node {
	newRoot := root.right
	root.right = newRoot.left
	newRoot.left = root
	root.UpdateProperties()
	newRoot.UpdateProperties()
	return newRoot
}

func (root *Node) UpdateProperties() {
	hleft, hright := 0, 0
	
	if root.left != nil {
		hleft = root.left.height
	}
	if root.right != nil {
		hright = root.right.height
	}
	root.bf = hright - hleft
	if root.left == nil && root.right == nil {
		root.height = 0
	} else {
		if hleft >= hright {
			root.height = hleft + 1
		} else {
			root.height = hright + 1
		}
	}
}