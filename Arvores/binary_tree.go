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
	levelOrder()
	Remove(val int) bool
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
	} else {
		return bst.root.Min(), nil
	}
}

func (no *Node) Min() int {
	val := no
	for val.left != nil {val = val.left}
	return val.val
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

func (bst *BST) Remove(val int) bool {
	if bst.root == nil {
		return false
	}
	var removed bool
	bst.root, removed = bst.root.RemoveNode(val)
	if removed {
		bst.size--
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
		// Achamos o nó
		removed = true
		if no.left == nil && no.right == nil {
			// Ramo folha
			return nil, removed
		} else if no.left != nil && no.right == nil {
			// Temos um filho a esquerda
			return no.left, removed
		} else if no.left == nil && no.right != nil {
			// Temos um filho a direita
			return no.right, removed
		} else {
			// Nó tem dois filhos
			min := no.right.Min()
			no.val = min
			no.right, _ = no.right.RemoveNode(min)
		}
	}
	return no, removed
}

func main(){
	bst := &BST{}
	bst.Add(10)
	bst.Add(5)
	bst.Add(3)
	bst.Add(8)
	bst.Add(9)

    // Tamanho inicial
    fmt.Println("--- Teste de Tamanho ---")
    fmt.Println("Tamanho inicial:", bst.size)
    fmt.Println()

    // Testes de Search
    fmt.Println("--- Testes de Search ---")
	fmt.Println("Search(10):", bst.Search(10)) // true
	fmt.Println("Search(5):", bst.Search(5))   // true
	fmt.Println("Search(3):", bst.Search(3))   // true
	fmt.Println("Search(8):", bst.Search(8))   // true
	fmt.Println("Search(9):", bst.Search(9))   // true
	fmt.Println("Search(20):", bst.Search(20)) // false
    fmt.Println()

    // Testes de Min/Max
    fmt.Println("--- Testes de Min/Max ---")
	fmt.Println(bst.Min())
	fmt.Println(bst.Max())
    fmt.Println()

    // Teste de Altura
    fmt.Println("--- Teste de Altura ---")
	fmt.Println("Altura:", bst.Height())
	fmt.Println()

    // Testes de Travessia
    fmt.Println("--- Teste preOrder ---")
	bst.preOrder()
	fmt.Println()

    fmt.Println("--- Teste inOrder ---")
	bst.inOrder()
	fmt.Println()

    fmt.Println("--- Teste posOrder ---")
	bst.posOrder()
	fmt.Println()

    // --- Testes de Remoção ---
    fmt.Println("--- Testes de Remoção ---")

    fmt.Println("\nRemovendo 999 (não existe)")
    // Use 'removed :=' se você implementou a Opção 2 (com retorno bool)
    bst.Remove(999) 
    fmt.Println("Tamanho atual (deve ser 5):", bst.size) 
    fmt.Println()

    fmt.Println("Removendo 3 (nó folha)")
    bst.Remove(3)
    fmt.Println("Tamanho atual (deve ser 4):", bst.size)
    fmt.Println("Nova travessia inOrder (sem o 3):")
    bst.inOrder() // Deve mostrar 5, 8, 9, 10
    fmt.Println()

    fmt.Println("Removendo 10 (raiz com 2 filhos)")
    bst.Remove(10)
    fmt.Println("Tamanho atual (deve ser 3):", bst.size)
    fmt.Println("Nova travessia inOrder (sem o 10):")
    bst.inOrder() // Deve mostrar 5, 8, 9
    fmt.Println()
    
    fmt.Println("Nova travessia preOrder (raiz deve ser 8 ou 9*):")
    // *Seu Add(val <= no.val) põe iguais à esquerda.
    // Ao remover 10, o sucessor é 8 (Min() da direita). 
    // Ah, não, 10 não tem sub-árvore direita no seu exemplo.
    // Deixe-me corrigir seu exemplo.
    // 10 -> 5 -> 3
    //      -> 8 -> 9
    // O sucessor de 10 (Min da direita) não existe.
    // O antecessor de 10 (Max da esquerda) é 9.
    // Ah, não, o Max da esquerda é 9.
    // Vamos ver: 10 tem 5 à esquerda. 5 tem 3(L) e 8(R). 8 tem 9(R).
    // O sucessor de 10 é o Mínimo da sub-árvore da DIREITA. Mas 10 não tem filho à direita.
    // O sucessor de 5 é o Mínimo da sub-árvore da DIREITA, que é 8.
    // O sucessor de 8 é 9.
    // O antecessor de 10 (Max da esquerda) é 9.
    // Se sua lógica de remoção (dois filhos) usa o sucessor (Min da direita), 
    // seu exemplo de 'Remove(10)' vai falhar, pois 10 não tem filho à direita.
    // Ah, eu li errado. 10 é a raiz. 5 é filho da esquerda. Não há filho da direita.
    // Então, 'Remove(10)' é um caso de "1 filho (esquerda)".
    // A nova raiz deve ser 5.
    
    fmt.Println("Nova travessia preOrder (nova raiz deve ser 5):")
    bst.preOrder() // Deve mostrar 5, 3, 8, 9
    fmt.Println()
}