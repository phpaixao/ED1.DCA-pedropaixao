package main

import (
	"errors"
	"fmt"
)

type IStack interface {
	Push(value int)    //Feito
	Pop() (int, error) //Feito
	Top() (int, error) //Feito
	IsEmpty() bool	   //Feito
	Size() int         //Feito
}

type Node struct {
	value int
	next * Node
}

type LinkedStack struct {
	top * Node
	inserted int
}

func (pilha *LinkedStack) Push(value int) {
	new_node := &Node{
		value: value,
		next: pilha.top,
	}
	pilha.top = new_node
	pilha.inserted++
}

func (pilha *LinkedStack) Pop() (int, error) {
	if pilha.IsEmpty(){
		return -1, errors.New("Lista vazia.")
	}
	val := pilha.top.value
	pilha.top = pilha.top.next
	pilha.inserted--
	return val, nil
}

func(pilha *LinkedStack) Top() (int, error) {
	if pilha.IsEmpty(){
		return -1, errors.New("Lista vazia.")
	}
	return pilha.top.value, nil
}

func (pilha *LinkedStack) IsEmpty() bool {
	return pilha.Size() == 0
}

func (pilha *LinkedStack) Size() int {
	return pilha.inserted
}

func (pilha *LinkedStack) Invert() error {
	if pilha.IsEmpty() {
		return errors.New("Lista vazia.")
	}
	aux := pilha.top
	prox := aux.next
	aux.next = nil
	ant := aux
	for i:=0; i<pilha.Size()-1; i++{
		aux = prox
		prox = aux.next
		aux.next = ant
		ant = aux
	}
	pilha.top = ant
	return nil
}

func (pilha *LinkedStack) Print() {
	auxV := make([] int, pilha.Size())
	aux := pilha.top
	for i:=0; i<pilha.Size(); i++{
		auxV[i] = aux.value
		aux = aux.next
	}

	for i:=pilha.Size()-1; i>=0; i--{
		fmt.Printf("%d ", auxV[i])
	}
	fmt.Printf("\n")
}

func main() {
	var pilha LinkedStack

	fmt.Println("🔹 Testando pilha encadeada:")

	// Teste inicial
	fmt.Println("A pilha está vazia?", pilha.IsEmpty())
	fmt.Println("Tamanho inicial:", pilha.Size())

	// Inserindo elementos
	fmt.Println("\nInserindo elementos...")
	pilha.Push(10)
	pilha.Push(20)
	pilha.Push(30)

	fmt.Println("Tamanho após inserções:", pilha.Size())
	fmt.Print("Estado da pilha: ")
	pilha.Print() // imprime do fundo para o topo: 10 20 30

	// Topo da pilha
	if val, err := pilha.Top(); err == nil {
		fmt.Println("Elemento no topo:", val)
	} else {
		fmt.Println("Erro:", err)
	}

	// Removendo elementos
	if val, err := pilha.Pop(); err == nil {
		fmt.Println("Elemento removido:", val)
	} else {
		fmt.Println("Erro:", err)
	}

	fmt.Print("Estado da pilha após Pop: ")
	pilha.Print() // imprime do fundo para o topo: 10 20

	if val, err := pilha.Top(); err == nil {
		fmt.Println("Novo topo:", val)
	} else {
		fmt.Println("Erro:", err)
	}

	fmt.Println("Tamanho agora:", pilha.Size())

	// Esvaziando a pilha
	fmt.Println("\nEsvaziando a pilha...")
	for !pilha.IsEmpty() {
		val, _ := pilha.Pop()
		fmt.Println("Removido:", val)
	}

	fmt.Println("A pilha está vazia?", pilha.IsEmpty())
	fmt.Println("Tamanho final:", pilha.Size())

	// Testando erro ao remover de pilha vazia
	if _, err := pilha.Pop(); err != nil {
		fmt.Println("Tentativa de Pop em pilha vazia -> Erro:", err)
	}

	if _, err := pilha.Top(); err != nil {
		fmt.Println("Tentativa de Top em pilha vazia -> Erro:", err)
	}
}
