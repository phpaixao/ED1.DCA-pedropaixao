package main

import (
	"fmt"
	"errors"
)

type BstNode struct {
	left   *BstNode
	value  int
	height int
	bf     int
	right  *BstNode
}

func (root *BstNode) UpdateProperties() {
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

func (root *BstNode) RotRight() *BstNode {
	newRoot := root.left
	root.left = newRoot.right
	newRoot.right = root
	root.UpdateProperties()
	newRoot.UpdateProperties()
	return newRoot
}

func (root *BstNode) RotLeft() *BstNode {
	newRoot := root.right
	root.right = newRoot.left
	newRoot.left = root
	root.UpdateProperties()
	newRoot.UpdateProperties()
	return newRoot
}

func (root *BstNode) RebalanceLeftLeft() *BstNode {
	return root.RotRight()
}

func (root *BstNode) RebalanceLeftNeutral() *BstNode {
	return root.RotRight()
}

func (root *BstNode) RebalanceLeftRight() *BstNode {
	root.left = root.left.RotLeft()
	return root.RotRight()
}

func (root *BstNode) RebalanceRightRight() *BstNode {
	return root.RotLeft()
}

func (root *BstNode) RebalanceRightNeutral() *BstNode {
	return root.RotLeft()
}

func (root *BstNode) RebalanceRightLeft() *BstNode {
	root.right = root.right.RotRight()
	return root.RotLeft()
}

func (root *BstNode) Rebalance() *BstNode {
	if root == nil {
		return nil
	}
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

func NewNode(val int) *BstNode {
	return &BstNode{
		left: nil,
		value: val,
		height: 0,
		bf: 0,
		right: nil,
	}
}

func (no *BstNode) Min() int {
	for no.left != nil {
		no = no.left
	}
	return no.value
}

func (no *BstNode) Max() int {
	for no.right != nil {
		no = no.right
	}
	return no.value
}

func (root *BstNode) Add(value int) *BstNode {
	if value <= root.value {
		if root.left == nil {
			root.left = NewNode(value)
		} else {
			root.left = root.left.Add(value)
		}
	} else {
		if root.right == nil {
			root.right = NewNode(value)
		} else {
			root.right = root.right.Add(value)
		}
	}
	root.UpdateProperties()
	return root.Rebalance()
}

func (root *BstNode) Remove(value int) *BstNode {
	if root == nil {return nil}
	if value < root.value {
		root.left = root.left.Remove(value)
	} else if value > root.value {
		root.right = root.right.Remove(value)
	} else { //encontramos o nó a ser removido
		if root.left == nil && root.right == nil {//nó folha        
			root = nil
		} else if root.left != nil && root.right == nil { 
                //caso 2: nó com 1 filho (à esquerda)
			root = root.left
		} else if root.left == nil && root.right != nil { 
                //caso 2: nó com 1 filho (à direita)
			root = root.right
		} else { 
                //caso 3: nó com 2 filhos
			maxEsq := root.left.Max()
			root.value = maxEsq
			root.left = root.left.Remove(maxEsq)
		}
	}
	if root == nil {return nil}
	root.UpdateProperties()
	return root.Rebalance()
}

func (root *BstNode) IsAVL(min int, max int) bool {
	if root == nil {
		return true
	}
	
	if root.value < min || root.value > max || root.bf < -1 || root.bf > 1 {
		return false
	}

	return root.left.IsAVL(min, root.value) && root.right.IsAVL(root.value, max)
}

func printTree(root *BstNode, indent string, last bool) {
	if root != nil {
		fmt.Print(indent)
		if last {
			fmt.Print("R----")
			indent += "   "
		} else {
			fmt.Print("L----")
			indent += "|  "
		}
		// Mostra Valor, Fator de Balanceamento (bf) e Altura (h)
		fmt.Printf("%d (bf:%d, h:%d)\n", root.value, root.bf, root.height)
		
		// Recursão para os filhos
		printTree(root.left, indent, false)
		printTree(root.right, indent, true)
	}
}

func main() {
	var root *BstNode

	fmt.Println("=== 1. TESTE DE INSERÇÃO (Balanceamento) ===")
	// Inserindo valores que forçam rotações simples e duplas
	// Sequência: 10, 20, 30 (Rotação Esq), 40, 50 (Rot Esq), 25 (Rot Dupla)
	vals := []int{10, 20, 30, 40, 50, 25}
	
	for _, v := range vals {
		fmt.Printf("-> Inserindo %d\n", v)
		if root == nil {
			root = NewNode(v)
		} else {
			root = root.Add(v)
		}
	}

	fmt.Println("\n--- Árvore Após Inserções ---")
	printTree(root, "", true)

	// Validação automática
	if root.IsAVL(math.MinInt64, math.MaxInt64) {
		fmt.Println("\n[✔] A árvore é uma AVL válida!")
	} else {
		fmt.Println("\n[✖] ERRO: A árvore perdeu a propriedade AVL.")
	}

	fmt.Println("\n=== 2. TESTE DE REMOÇÃO (Caso Complexo) ===")
	// O nó 30 deve ser a raiz atual e ter 2 filhos. 
	// Sua lógica pega o MAX da ESQUERDA para substituir.
	fmt.Println("-> Removendo 30 (Raiz com 2 filhos)...")
	root = root.Remove(30)

	fmt.Println("\n--- Árvore Após Remoção ---")
	printTree(root, "", true)

	if root.IsAVL(math.MinInt64, math.MaxInt64) {
		fmt.Println("\n[✔] A árvore continua sendo uma AVL válida!")
	} else {
		fmt.Println("\n[✖] ERRO: O balanceamento quebrou após a remoção.")
	}

	fmt.Println("\n=== 3. TESTE DE REMOÇÃO (Caso Folha/Desbalanceamento) ===")
	// Remover o 50 pode causar desbalanceamento no pai e forçar rotação
	fmt.Println("-> Removendo 50...")
	root = root.Remove(50)
	printTree(root, "", true)
}