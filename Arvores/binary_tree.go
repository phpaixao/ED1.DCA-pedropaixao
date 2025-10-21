package main

import(
	"fmt"
	"errors"
)

type Node struct{
	val int
	left *Node
	right *Node
}

type BST struct{
	root *Node
	size int
}

type Tree interface{
	Add(val int) // Feito
	Search(val int) bool // Feito
	Height() int  // Feito
	Min() (int, error) // Feito
	Max() (int, error) // Feito
	preOrder() // Feito
	inOrder() // Feito
	posOrder() // Feito
	Remove() error // Feito
}

func createNode(val int) *Node{
	return &Node{
		val: val,
		left: nil,
		right: nil,
	}
}

func (no *Node) AddNode(val int){
	if val <= no.val{
		if no.left == nil{
			no.left = createNode(val)
		} else {
			no.left.AddNode(val)
		}
	} else {
		if no.right == nil{
			no.right = createNode(val)
		} else {
			no.right.AddNode(val)
		}
	}
}

func (bst *BST) Add(val int){
	if bst.root == nil{
		bst.root = createNode(val)
	} else {
		bst.root.AddNode(val)
	}
	bst.size++
}

func (bst *BST) Search(val int) bool {
	if bst.root == nil {
		return false
	} else {
		return bst.root.SearchNode(val)
	}
}

func (no *Node) SearchNode(val int) bool {
	if val == no.val{
		return true
	} else if val < no.val{
		if no.left == nil{
			return false
		} else {
			return no.left.SearchNode(val)
		}
	} else {
		if no.right == nil{
			return false
		} else {
			return no.right.SearchNode(val)
		}
	}
}

func (bst *BST) Height() int {
	if bst.root == nil {
		return 0 
	} else {
		return bst.root.NodeHeight()
	}
}

func (no *Node) NodeHeight() int {
	if no.left == nil && no.right == nil {
		return 0
	}

	hRelativeLeft := 0
	if no.left != nil {
		hRelativeLeft = 1 + no.left.NodeHeight()
	}
	hRelativeRight := 0
	if no.right != nil {
		hRelativeRight = 1 + no.right.NodeHeight()
	}

	if hRelativeLeft >= hRelativeRight {
		return hRelativeLeft
	} else {
		return hRelativeRight
	}
}

func (bst *BST) preOrder() {
	if bst.root != nil {
		bst.root.preOrder()
	}
}

func (no *Node) preOrder() {
	fmt.Println(no.val)
	if no.left != nil {
		no.left.preOrder()
	}
	if no.right != nil {
		no.right.preOrder()
	}
}

func (bst *BST) inOrder() {
	if bst.root != nil {
		bst.root.inOrder()
	}
}

func (no *Node) inOrder() {
	if no.left != nil {
		no.left.inOrder()
	}
	fmt.Println(no.val)
	if no.right != nil {
		no.right.inOrder()
	}
}

func (bst *BST) posOrder() {
	if bst.root != nil {
		bst.root.posOrder()
	}
}

func (no *Node) posOrder() {
	if no.left != nil {
		no.left.posOrder()
	}
	if no.right != nil {
		no.right.posOrder()
	}
	fmt.Println(no.val)
}


func (bst *BST) Min() (int, error) {
	if bst.root == nil {
		return -1, errors.New("Arvore vazia.")
	}
	no := bst.root
	for no.left != nil {
		no = no.left
	}
	return no.val, nil 
}

func (bst *BST) Max() (int, error) {
	if bst.root == nil {
		return -1, errors.New("Arvore vazia.")
	}
	no := bst.root
	for no.right != nil {
		no = no.right
	}
	return no.val, nil
}

func main(){
	bst := &BST{}
	bst.Add(10)
	bst.Add(5)
	bst.Add(3)
	bst.Add(8)
	bst.Add(9)

	fmt.Println(bst.Search(10))
	fmt.Println(bst.Search(5))
	fmt.Println(bst.Search(3))
	fmt.Println(bst.Search(8))
	fmt.Println(bst.Search(9))
	fmt.Println(bst.Search(20))

	fmt.Println(bst.Min())
	fmt.Println(bst.Max())

	fmt.Println(bst.Height())
	fmt.Println()

	bst.preOrder()
	fmt.Println()

	bst.inOrder()
	fmt.Println()

	bst.posOrder()
	fmt.Println()
}