package main

import(
	"fmt"
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
	Add(val int)
	Search(val int) bool
	Height() int
	Min() int
	Max() int
	preOrder()
	inOrder() bool
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

func main(){
	bst := &BST{}
	bst.Add(10)
	bst.Add(5)
	bst.Add(3)
	bst.Add(8)

	fmt.Println(bst.Search(10))
	fmt.Println(bst.Search(5))
	fmt.Println(bst.Search(3))
	fmt.Println(bst.Search(8))
	fmt.Println(bst.Search(20))
}