package main

import (
	"errors"
	"fmt"
)

type List interface {
	Tamanho() int                      //Feito
	Get(index int) (int, error)        //Feito
	Append(value int)                  //Feito
	Insert(index int, value int) error 
	Remove(index int) error            
}

type Node struct {
	value int
	next * Node
}

type LinkedList struct {
	head * Node
	inserted int
}


// Retorna o tamanho da linkedlist
func (lista *LinkedList) Tamanho() int {
	return lista.inserted
}

// Retorna o elemento guardado em um indice especifico da linkedlist
func (lista *LinkedList) Get(index int) (int, error) {
	if index < 0 || index >= lista.inserted {
		return -1, errors.New(fmt.Sprintf("Index out of range: %d", index))
	}
	aux := lista.head
	for i:=0; i<index; i++ {
		aux = aux.next
	}
	return aux.value, nil
}

/* ## NAO EXISTE ESPAÇO DE MEMORIA OCIOSA EM UMA LINKEDLIST ##
func (lista *LinkedList) doubleV() {
	newV := make([]int, len(lista.vetor)*2)
	for i := 0; i < len(lista.vetor); i++ {
		newV[i] = lista.vetor[i]
	}
	lista.vetor = newV
}
*/

// Adiciona um elemento ao final da linkedlist
func (lista *LinkedList) Append(value int) {
	novo_node := &Node{
		value: value,
		next: nil,
	}
	if lista.Tamanho() == 0{
		lista.head = novo_node
	} else {
		aux := lista.head
		for aux.next != nil {
			aux = aux.next
		}
		aux.next = novo_node
		}
	lista.inserted++
}

//Adiciona um elemento no inicio da linkedlist
func (lista *LinkedList) PushFront(value int){
	novo_node := &Node{
		value: value,
		next: nil,
	}
	if lista.head == nil{
		lista.head = novo_node
	} else {
		aux := lista.head
		lista.head = novo_node
		lista.head.next = aux
	}
	lista.inserted++
}

// Adiciona um elemento a um indice em especifico da linkedlist
func (lista *LinkedList) Insert(index int, value int) error {
	if index < 0 || index > lista.Tamanho() {
		return errors.New(fmt.Sprintf("Index out of range: %d", index))
	}
	if index == 0{
		lista.PushFront(value)
		return nil
	} 
	aux := lista.head
	// Aux da posição i sempre aponta pro node de i+1
	for i := 0; i < index-1; i++{ 
		aux = aux.next
	}
	novo_node := &Node{
		value: value,
		next: aux.next,
	}
	aux.next = novo_node
	lista.inserted++
	return nil
}

// Remove um indice da linkedlist
func (lista *LinkedList) Remove(index int) error {
	if index < 0 || index >= lista.inserted {
		return errors.New(fmt.Sprintf("Index out of range: %d", index))
	} 

	if index == 0{
		lista.head = lista.head.next
		lista.inserted--
		return nil
	}
	aux := lista.head
	// Aux da posição i sempre aponta pro node de i+1
	for i := 0; i<index-1; i++{
		aux = aux.next
	}
	// O next do aux passa a ser o next do node atual que aux aponta
	aux.next = aux.next.next
	lista.inserted--
	return nil
}
// Troca o valor guardado na posicao index por value
func (lista *LinkedList) Set(index int, value int) error {
	if index < 0 || index >= lista.Tamanho() {
		return errors.New(fmt.Sprintf("Index out of range: %d", index))
	}
	if index == 0{
		lista.head.value = value
		return nil
	}
	aux := lista.head
	// Aux da posição i sempre aponta pro node de i+1
	for i := 0; i<index; i++{
		aux = aux.next
	}
	aux.value = value
	return nil
}

// Remove o ultimo elemento da linkedlist
func (lista *LinkedList) Pop() error {
	if lista.Tamanho() <= 0 {
		return errors.New(fmt.Sprintf("Lista já possui tamanho igual a 0.\n"))
	}
	if lista.Tamanho() == 1{
		lista.head = nil
		lista.inserted--
		return nil
	}
	aux := lista.head
	// Aux da posição i sempre aponta pro node de i+1 
	for i := 0; i<lista.Tamanho() - 2; i++{
		aux = aux.next
	}
	aux.next = nil
	lista.inserted--
	return nil
}

// Escreve a linkedlist na tela
func (lista *LinkedList) Print() {
	aux := lista.head
	for i := 0; i<lista.Tamanho(); i++{
		fmt.Printf("%d ", aux.value)
		aux = aux.next
	}
	fmt.Printf("\n")
}

// Função main para testar as funções implementadas
func main() {
    
	// inicializa a lista vazia
	l := &LinkedList{}

    // Adiciona elementos no final (Append)
    for i := 0; i < 10; i++ {
        l.Append(i)
    }
    fmt.Println("Depois do Append:")
    l.Print()

    // Insere no índice 2
    l.Insert(2, 99)
    fmt.Println("Depois do Insert(2, 99):")
    l.Print()

    // Remove do índice 2
    l.Remove(2)
    fmt.Println("Depois do Remove(2):")
    l.Print()

    // Altera valor do índice 0
    l.Set(0, -10)
    fmt.Println("Depois do Set(0, -10):")
    l.Print()

    // Remove o primeiro elemento
    l.Remove(0)
    fmt.Println("Depois do Remove(0):")
    l.Print()

    // Pega o valor do índice 0
    x, _ := l.Get(0)
    fmt.Printf("Valor do índice 0: %d\n", x)

    // Pop (remove último elemento)
    l.Pop()
    fmt.Println("Depois do Pop():")
    l.Print()

    // PushFront (insere no início)
    l.PushFront(555)
    fmt.Println("Depois do PushFront(555):")
    l.Print()

    // Tamanho final
    fmt.Printf("Tamanho final da lista: %d\n", l.Tamanho())
}
