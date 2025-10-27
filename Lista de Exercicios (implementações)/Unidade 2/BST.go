package main

import(
	"fmt"
	"errors"
)

type ITree interface {
	Add(value int)				// Feito
	Search(value int) bool		// Feito
	Min() (int, error)			// Feito
	Max() (int, error)			// Feito
	PrintPre()					// Feito
	PrintIn()					// Feito
	PrintPos()					// Feito
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

func createNode(val int) *Node {
	return &Node{
		val: val,
		left: nil,
		right: nil,
	}
}

func (bst *BST) Add(val int){
	if bst.root == nil {
		bst.root = createNode(val)
	}
	bst.root.AddNode(val)
}

func (no *Node) AddNode(val int) {
	if val < no.val {
		if no.left == nil {
			no.left = createNode(val)
		} else {
			no.left.AddNode(val)
		}
	} else {
		if no.right == nil {
			no.right = createNode(val)
		} else {
			no.right.AddNode(val)
		}
	}
}

func (bst *BST) Search(val int) bool {
	if bst.root == nil {
		return false
	}
	return bst.root.SearchNode(val)
}

func (no *Node) SearchNode(val) bool {
	if val == no.val {
		return true
	}
	if val < no.val {
		if no.left == nil {
			return false
		} else {
			return no.left.SearchNode(val)
		}
	} else {
		if no.right == nil {
			return false
		} else {
			return no.right.SearchNode(val)
		}
	}
}

func (bst *BST) Min() (int, error) {
	if bst.root == nil {
		return -1, errors.New("Empty BST")
	}
	return bst.root.MinNode(), nil
}

func (no *Node) MinNode() int {
	for no.left != nil {
		no = no.left
	}
	return no.val
}

func (bst *BST) Max() (int, error) {
	if bst.root == nil {
		return -1, errors.New("Empty BST")
	}
	return bst.root.MaxNode(), nil
}

func (no *Node) MaxNode() int {
	for no.right != nil {
		no = no.right
	}
	return no.val
}

func (bst *BST) PrintPre() {
	if bst.root != nil {
		bst.root.PreOrder()
	}
}

func (no *Node) PreOrder() {
	fmt.Println(no.val)
	if no.left != nil {
		no.left.PreOrder()
	}
	if no.right != nil {
		no.right.PreOrder()
	}
}

func (bst *BST) PrintIn() {
	if bst.root != nil {
		bst.root.InOrder()
	}
}

func (no *Node) InOrder() {
	if no.left != nil {
		no.left.InOrder()
	}
	fmt.Println(no.val)
	if no.right != nil {
		no.right.InOrder()
	}
}

func (bst *BST) PrintPos() {
	if bst.root != nil {
		bst.root.PosOrder()
	}
}

func (no *Node) PosOrder() {
	if no.left != nil {
		no.left.PosOrder()
	}
	if no.right != nil {
		no.right.PosOrder()
	}
	fmt.Println(no.val)
}