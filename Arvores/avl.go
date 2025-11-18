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
	inserted int
}

type AVLTree interface {
	Add(val int) //*Node
	Search(val int) bool 
	Height() int
	Min() (int, error)
	Max() (int, error)
	preOrder()
	inOrder()
	posOrder()
	levelOrder()
	Remove(val int) bool
}

type AVLNode interface {
	UpdateProperties()				// Feito
	RotRight() *Node				// Feito
	RotLeft() *Node					// Feito
	RebalanceLeftLeft() *Node		// Feito
	RebalanceLeftNeutral() *Node	// Feito
	RebalanceLeftRight() *Node		// Feito
	RebalanceRightRight() *Node		// Feito
	RebalanceRightNeutral() *Node	// Feito
	RebalanceRightLeft() *Node		// Feito
	Rebalance() *Node				// Feito
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

func createNode(val int) *Node {
	return &Node{
		left: nil,
		val: val,
		right: nil,
		bf: 0,
		height: 0,
	}
}

func (avl *AVL) Add(val int) {
	if avl.root == nil {
		avl.root = createNode(val)
	} else {
		avl.root = avl.root.AddNode(val)	
	}
	avl.inserted++
}

func (no *Node) AddNode(val int) *Node {
	if val < no.val {
		if no.left == nil {
			no.left = createNode(val)
		} else {
			no.left = no.left.AddNode(val)
		}
	} else {
		if no.right == nil {
			no.right = createNode(val)
		} else {
			no.right = no.right.AddNode(val)
		}
	}
	no.UpdateProperties()
	return no.Rebalance()
}

func (avl *AVL) Search(val int) bool {
	if avl.root == nil {return false}
	return avl.root.SearchNode(val)
}

func (no *Node) SearchNode(val int) bool {
	if no == nil {return false}
	if val == no.val {return true}
	if val < no.val {
		return no.left.SearchNode(val)
	} else {
		return no.right.SearchNode(val)
	}
}

func (avl *AVL) Height() int {
	if avl.root == nil {return -1}
	return avl.root.NodeHeight()
}

func (no *Node) NodeHeight() int {return no.height}

func (avl *AVL) Min() (int, error) {
	if avl.root == nil {
		return -1, errors.New("Empty AVL Tree.")
	}
	return avl.root.MinNode(), nil
}

func (no *Node) MinNode() int {
	for no.left != nil {
		no = no.left
	}
	return no.val
}

func (avl *AVL) Max() (int, error) {
	if avl.root == nil {
		return -1, errors.New("Empty AVL Tree.")
	}
	return avl.root.MaxNode(), nil
}

func (no *Node) MaxNode() int {
	for no.right != nil {
		no = no.right
	}
	return no.val
}

func (avl *AVL) preOrder() {
	if avl.root != nil {
		avl.root.preOrder()
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

func (avl *AVL) inOrder() {
	if avl.root != nil {
		avl.root.inOrder()
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

func (avl *AVL) posOrder() {
	if avl.root != nil {
		avl.root.posOrder()
	}
}

func (no *Node) posOrder() {
	if no.left != nil {
		no.left.posOrder()
	}
	fmt.Println(no.val)
	if no.right != nil {
		no.right.posOrder()
	}
}

func (avl *AVL) levelOrder() {
	
}

func (avl *AVL) Remove(val int) bool {
	if avl.root == nil {return false}
	var removido bool
	avl.root, removido = avl.root.RemoveNode(val)
	if removido {avl.inserted--}
	return removido 
}

func (no *Node) RemoveNode(val int) (*Node, bool) {
	if no == nil {return nil, false}
	
	var removed bool
	if val < no.val {
		no.left, removed = no.left.RemoveNode(val)
	} else if val > no.val {
		no.right, removed = no.right.RemoveNode(val)
	} else {
		// Encontramos o nó que queremos remover
		removed = true
		if no.left == nil && no.right == nil {
			no = nil
		} else if no.left != nil && no.right == nil {
			no = no.left
		} else if no.left == nil && no.right != nil {
			no = no.right
		} else {
			min := no.right.MinNode()
			no.val = min
			no.right, _ = no.right.RemoveNode(min) 
		}
	}
	if no == nil {return nil, removed}
	no.UpdateProperties()
	return no.Rebalance(), removed
}

func main(){

}