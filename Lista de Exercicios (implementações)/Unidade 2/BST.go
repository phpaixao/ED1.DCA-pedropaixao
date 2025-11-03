package main

import(
	"fmt"
	"errors"
	"math"
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
	fmt.Printf("%d ", no.val)
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
	fmt.Printf("%d ", no.val)
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
	fmt.Printf("%d ", no.val)
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

func (bst *BST) IsBst() bool {
	return bst.root.IsBst(math.MinInt, math.MaxInt)
}


// Essa versão funciona também, só é mais simplificada
func (no *Node) IsBst(min int, max int) bool {
	if no == nil {
		return true
	}

	if no.val >= max || no.val < min {
		return false
	}

	return no.left.IsBst(min, no.val) && no.right.IsBst(no.val, max)
}

/*
func (no *Node) IsBst(min int, max int) bool {
	if no == nil {
		return true
	}

	if no.val >= max || no.val < min {
		return false
	}
	if no.left == nil && no.right == nil {
		return true
	}

	if no.left != nil && no.right == nil {
		return no.left.IsBst(min, no.val)
	}
	if no.left == nil && no.right != nil {
		return no.right.IsBst(no.val, max)
	}

	return no.left.IsBst(min, no.val) && no.right.IsBst(no.val, max)
}
*/

func convertToBalancedBst(v []int, ini int, fim int) *Node {
	if ini > fim {return nil}

	mid := (ini+fim)/2

	no := &Node{
		val: v[mid],
		left: convertToBalancedBst(v, ini, mid-1),
		right: convertToBalancedBst(v, mid+1, fim),
	}
	return no
}

func (bst *BST) Par() int {
	return bst.root.Par()
}

func (no *Node) Par() int {
	if no == nil {
		return 0
	}

	count := 0
	queue := []*Node{no}
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i<n; i++ {
			no := queue[0]
			queue = queue[1:]
			
			if no.val % 2 == 0 {
				count++
			}
			if no.left != nil {
				queue = append(queue, no.left)
			}
			if no.right != nil {
				queue = append(queue, no.right)
			}
		}
	}
	return count
}

// Função para mudar o valor de um Nó da árvore e 
// fazer com que ela deixe de ser uma BST
func (bst *BST) SetValue(val int, newVal int) {
	bst.root.SetValueNode(val, newVal)
}

func (no *Node) SetValueNode(val int, newVal int) {
	if no == nil {
		return
	}

	if no.val == val {
		no.val = newVal
		return
	} 

	if val < no.val {
		no.left.SetValueNode(val, newVal)
	} else {
		no.right.SetValueNode(val, newVal)
	}
}

func (bst *BST) FixValue(val int, newVal int) {
	bst.root.FixValueNode(val, newVal)
}

func (no *Node) FixValueNode(val int, newValue int) {
	if no == nil {
		return
	}

	if val == no.val {
		no.val = newValue
		return
	}

	no.left.FixValueNode(val, newValue)
	no.right.FixValueNode(val, newValue)
}

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
 	fmt.Println() // <--- AJUSTE DE FORMATAÇÃO
 	
 	fmt.Println("\n--- PrintLevels (Original) ---")
 	bst.PrintLevels()
 	// Saída:
 	// 10 
 	// 5 15 
 	// 3 7 12 18 

 	fmt.Println("\n--- Altura, Min, Max, Pares ---")
 	fmt.Printf("Altura da árvore: %d\n", bst.Height()) // Saída: 2
 	
 	if min, err := bst.Min(); err == nil {
 		fmt.Printf("Valor Mínimo: %d\n", min) // Saída: 3
 	}
 	if max, err := bst.Max(); err == nil {
 		fmt.Printf("Valor Máximo: %d\n", max) // Saída: 18
 	}
 	fmt.Printf("Qtd Pares: %d\n", bst.Par()) // Saída: 3 (10, 12, 18)


 	fmt.Println("\n--- Teste de Busca (Search) ---")
 	fmt.Printf("Buscando 7: %t\n", bst.Search(7)) 	 // Saída: true
 	fmt.Printf("Buscando 99: %t\n", bst.Search(99)) // Saída: false


 	fmt.Println("\n--- Teste de Percursos (Pre e Pos) ---")
 	fmt.Println("PreOrder:")
 	bst.PrintPre() // Saída: 10 5 3 7 15 12 18
 	fmt.Println() // <--- AJUSTE DE FORMATAÇÃO

 	fmt.Println("PosOrder:") // Removido o \n para ficar junto com o PrintPre
 	bst.PrintPos() // Saída: 3 7 5 12 18 15 10
 	fmt.Println() // <--- AJUSTE DE FORMATAÇÃO

 	
 	fmt.Println("\n\n--- Teste IsBst (Adulterando Árvore) ---")
 	fmt.Printf("Árvore é BST? (Antes): %t\n", bst.IsBst()) // Saída: true

 	fmt.Println("...Adulterando a árvore: Trocando 7 por 99...")
 	bst.SetValue(7, 99) // Função para quebrar a BST
 	
 	fmt.Printf("Árvore é BST? (Depois): %t\n", bst.IsBst()) // Saída: false

 	fmt.Println("...Consertando a árvore: Trocando 99 por 7...")
 	bst.FixValue(99, 7) // <--- AJUSTE DE LÓGICA (Usando a função correta)
 	
 	fmt.Printf("Árvore é BST? (Consertada): %t\n", bst.IsBst()) // Saída: true


 	fmt.Println("\n\n--- Teste de Remoção (Nó com 2 filhos: 15) ---")
 	bst.Remove(15) 
 	bst.PrintLevels()
 	// Saída: (Agora correta, sem o 99)
 	// 10 
 	// 5 18 
 	// 3 7 12 


 	fmt.Println("\n--- Teste de Remoção (Folha: 7) ---")
 	bst.Remove(7)
 	bst.PrintLevels()
 	// Saída:
 	// 10 
 	// 5 18 
 	// 3 12 


 	fmt.Println("\n--- Teste de Remoção (Nó com 1 filho: 5) ---")
 	bst.Remove(5)
 	bst.PrintLevels()
 	// Saída:
 	// 10 
 	// 3 18 
 	// 12 


 	fmt.Println("\n\n--- Teste convertToBalancedBst ---")
 	sortedArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
 	balancedRoot := convertToBalancedBst(sortedArr, 0, len(sortedArr)-1)
 	
 	// Criando uma nova BST com a raiz balanceada
 	bst2 := BST{root: balancedRoot}

 	fmt.Println("Árvore Balanceada (PrintLevels):")
 	bst2.PrintLevels()
 	// Saída:
 	// 5 
 	// 2 7 
 	// 1 3 6 8 
 	// 4 9 
 	
 	fmt.Printf("Altura da árvore balanceada: %d\n", bst2.Height()) // Saída: 3
 	fmt.Printf("Árvore balanceada é BST? %t\n", bst2.IsBst()) // Saída: true
}