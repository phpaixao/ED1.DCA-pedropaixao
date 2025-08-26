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

func main() {
	var pilha IStack = &LinkedStack{}

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