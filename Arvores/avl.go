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

type AVL struct {
	root *Node

}

type AVLTree interface {
	Add(val int) //*Node
	Search(val int) bool 
	Height() int
	Min() int
	Max() int
	preOrder()
	inOrder()
	posOrder()
	levelOrder()
	Remove(val int) bool
	// UpdateProperties()				// Feito
	// RotRight() *Node					// Feito
	// RotLeft() *Node					// Feito
	// RebalanceLeftLeft() *Node		// Feito
	// RebalanceLeftNeutral() *Node		// Feito
	// RebalanceLeftRight() *Node		// Feito
	// RebalanceRightRight() *Node		// Feito
	// RebalanceRightNeutral() *Node	// Feito
	// RebalanceRightLeft() *Node		// Feito
	// Rebalance() *Node				// Feito
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

func (root *Node) RebalanceLeftLeft() *Node {return root.RotRight()}

func (root *Node) RebalanceLeftNeutral() *Node {return root.RotRight()}

func (root *Node) RebalanceLeftRight() *Node {
	root.left = root.left.RotLeft()
	return root.RotRight()
}

func (root *Node) RebalanceRightRight() *Node {return root.RotLeft()}

func (root *Node) RebalanceRightNeutral() *Node {return root.RotLeft()}

func (root *Node) RebalanceRightLeft() *Node {
	root.right = root.right.RotRight()
	return root.RotLeft()
}

func (root *Node) Rebalance() *Node {
	if root == nil {return nil}
	if root.bf < -1 {
		if root.left.bf == -1 {
			root = root.RebalanceLeftLeft()
		} else if root.left.bf == 0 {
			root = root.RebalanceLeftNeutral()
		} else {
			root = root.RebalanceLeftRight()
		}
	} else if root.bf > 1 {
		if root.right.bf == 1 {
			root = root.RebalanceRightRight()
		} else if root.right.bf == 0 {
			root = root.RebalanceRightNeutral()
		} else {
			root = root.RebalanceRightLeft()
		}
	}
	return root
}

func (avl *AVL) Add(val int) {

}

func (no *Node) AddNode(val int) *Node {
	
}