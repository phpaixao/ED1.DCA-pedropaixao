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

type ArrayStack struct {
	top int
	vetor [] int
}

func (pilha *ArrayStack) Init(size int) {
	pilha.vetor = make([] int, size)
}

func (pilha *ArrayStack) doubleV(){
	newV := make([] int, len(pilha.vetor)*2)
	for i := 0; i<pilha.Size(); i++{
		newV[i] = pilha.vetor[i]
	}
	pilha.vetor = newV
}

func (pilha *ArrayStack) Push(value int) {
	if pilha.Size() == len(pilha.vetor){
		pilha.doubleV()
	}
	pilha.vetor[pilha.top] = value
	pilha.top++
}

func (pilha *ArrayStack) Pop() (int, error) {
	if pilha.IsEmpty() {
		return -1, errors.New("Lista vazia.")
	}
	pilha.top--
	return pilha.vetor[pilha.top], nil
}

func (pilha *ArrayStack) Top() (int, error) {
	if pilha.IsEmpty() {
		return -1, errors.New("Lista vazia.")
	}
	return pilha.vetor[pilha.top-1], nil	
}

func (pilha *ArrayStack) IsEmpty() bool {
	return pilha.top == 0
}

func (pilha *ArrayStack) Size() int {
	return pilha.top
}

func (pilha *ArrayStack) Print() {
	for i:=0; i<pilha.Size(); i++{
		fmt.Printf("%d ", pilha.vetor[i])
	}
	fmt.Printf("\n")
}

func main() {
	var pilha ArrayStack
	pilha.Init(2) // inicializa com tamanho inicial 2

	fmt.Println("Pilha vazia?", pilha.IsEmpty()) // true

	// Testando Push
	fmt.Println("\n--- Testando Push ---")
	pilha.Push(10)
	pilha.Push(20)
	pilha.Push(30) // aqui já deve dobrar o vetor (de 2 -> 4)
	fmt.Println("Tamanho:", pilha.Size())        // 3
	fmt.Print("Conteúdo da pilha: ")
	pilha.Print() // mostra a pilha

	// Testando Top
	fmt.Println("\n--- Testando Top ---")
	if topo, err := pilha.Top(); err == nil {
		fmt.Println("Topo:", topo) // 30
	} else {
		fmt.Println("Erro:", err)
	}

	// Testando Pop
	fmt.Println("\n--- Testando Pop ---")
	if val, err := pilha.Pop(); err == nil {
		fmt.Println("Pop:", val) // 30
	}
	fmt.Print("Pilha após Pop: ")
	pilha.Print()

	if val, err := pilha.Pop(); err == nil {
		fmt.Println("Pop:", val) // 20
	}
	fmt.Print("Pilha após Pop: ")
	pilha.Print()

	if val, err := pilha.Pop(); err == nil {
		fmt.Println("Pop:", val) // 10
	}
	fmt.Print("Pilha após Pop: ")
	pilha.Print()

	// Tentando estourar Pop em pilha vazia
	if val, err := pilha.Pop(); err == nil {
		fmt.Println("Pop:", val)
	} else {
		fmt.Println("Erro ao dar Pop:", err)
	}

	fmt.Println("\nPilha vazia?", pilha.IsEmpty()) // true
	fmt.Println("Tamanho final:", pilha.Size())    // 0
	fmt.Print("Conteúdo final da pilha: ")
	pilha.Print()
}
