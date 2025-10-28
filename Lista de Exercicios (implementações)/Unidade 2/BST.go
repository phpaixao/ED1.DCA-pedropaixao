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
	PrintLevels()				// Feito
	Height() int				// Feito		
	Remove(value int) bool		// Feito
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
	} else {
		bst.root.AddNode(val)
	}
	bst.inserted++
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

func (no *Node) SearchNode(val int) bool {
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

func (bst *BST) PrintLevels() {
	if bst.root == nil {return}

	queue := []*Node{bst.root}
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i<n; i++ {
			no := queue[0]
			queue = queue[1:]
			fmt.Printf("%d ", no.val)

			if no.left != nil {
				queue = append(queue, no.left)
			}
			if no.right != nil {
				queue = append(queue, no.right)
			}
		}
		fmt.Println()
	}
}

func (bst *BST) Height() int {
	if bst.root == nil {
		return -1
	}
	return bst.root.NodeHeight()
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

func (bst *BST) Remove(val int) bool {
	if bst.root == nil {
		return false
	}
	var removed bool
	bst.root, removed = bst.root.RemoveNode(val)
	if removed {
		bst.inserted--
	}
	return removed
}

func (no *Node) RemoveNode(val int) (*Node, bool) {
	if no == nil {
		return nil, false
	}

	var removed bool
	if val < no.val {
		no.left, removed = no.left.RemoveNode(val)
	} else if val > no.val {
		no.right, removed = no.right.RemoveNode(val)
	} else {
		//Encontramos o nó que estavamos buscando
		removed = true
		// Nó não tem filhos
		if no.left == nil && no.right == nil {
			return nil, removed
		} else if no.left != nil && no.right == nil {
			// Nó possui um filho
			return no.left, removed
		} else if no.left == nil && no.right != nil {
			// Nó possui um filho
			return no.right, removed
		} else {
			//Nó possui dois filhos
			min := no.right.MinNode()
			no.val = min
			no.right, _ = no.right.RemoveNode(min)
		}
	}
	return no, removed
}

/*
func (no *Node) IsBst() bool {
	if no == nil {
		return true
	}
	if no.left != nil && no.right == nil {
		if no.val < no.left.val {
			return false
		} else {
			return true && no.left.IsBst()
		}
	} else if no.left == nil && no.right != nil {
		if no.val > no.right.val {
			return false
		} else {
			return true && no.right.IsBst()
		}
	} else if no.left != nil && no.right != nil {

	}
}
*/

func main() {
	bst := BST{}
	bst.Add(10)
	bst.Add(5)
	bst.Add(15)
	bst.Add(3)
	bst.Add(7)
	bst.Add(12)
	bst.Add(18)

	fmt.Println("--- PrintIn (Ordenado) ---")
	bst.PrintIn() // Saída: 3 5 7 10 12 15 18
	fmt.Println("\n\n--- PrintLevels ---")
	bst.PrintLevels()
	// Saída:
	// 10 
	// 5 15 
	// 3 7 12 18 

	fmt.Println("\n--- Altura ---")
	fmt.Printf("Altura da árvore: %d\n", bst.Height()) // Saída: 2

	fmt.Println("\n--- Remoção (15) ---")
	bst.Remove(15) // Remove nó com 2 filhos
	bst.PrintLevels()
	// Saída:
	// 10 
	// 5 18 
	// 3 7 12 
}