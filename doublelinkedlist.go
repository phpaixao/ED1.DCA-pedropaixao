package main

import (
	"errors"
	"fmt"
)

type List interface {
	Tamanho() int                      //Feito
	Get(index int) (int, error)        
	Append(value int)                  
	Insert(index int, value int) error 
	Remove(index int) error            
}

type Node struct {
	value int
	next * Node
	prev * Node
}

type DoubleLinkedList struct {
	head * Node
	tail * Node
	inserted int
}

// Retorna o tamanho da doublelinkedlist
func (lista *DoubleLinkedList) Tamanho() int {
	return lista.inserted
}

// Retorna o elemento guardado em um indice especifico da doublelinkedlist
func (lista *DoubleLinkedList) Get(index int) (int, error) {
	if index < 0 || index >= lista.inserted {
		return -1, errors.New(fmt.Sprintf("Index out of range: %d", index))
	}
	var aux *Node
	if index <= lista.Tamanho()-1-index{
		aux = lista.head
		for i := 0; i<index; i++{
			aux = aux.next
		}
	} else {
		aux = lista.tail
		for i := lista.Tamanho()-1; i>index; i--{
			aux = aux.prev
		}
	}
	return aux.value, nil
}

// Adiciona um elemento ao final da doublelinkedlist
func (lista *DoubleLinkedList) Append(value int) {
	novo_node := &Node{
		value: value,
		next: nil,
		prev: lista.tail,
	}
	if lista.Tamanho() == 0{
		novo_node.prev = nil
		lista.head, lista.tail = novo_node, novo_node
	} else {
		lista.tail.next = novo_node
		lista.tail = novo_node	
	}
	lista.inserted++
}

//Adiciona um elemento no inicio da doublelinkedlist
func (lista *DoubleLinkedList) PushFront(value int){
	if lista.Tamanho() == 0{
		lista.Append(value)
	} else {
		novo_node := &Node{
			value: value,
			next: lista.head,
			prev: nil,
		}
		lista.head.prev = novo_node
		lista.head = novo_node
		lista.inserted++ 
	}
}

// Adiciona um elemento a um indice em especifico da doublelinkedlist
func (lista *DoubleLinkedList) Insert(index int, value int) error {
	if index < 0 || index > lista.Tamanho() {
		return errors.New(fmt.Sprintf("Index out of range: %d", index))
	}
	if index == 0{
		lista.PushFront(value)
		return nil
	}
	if index == lista.Tamanho(){
		lista.Append(value)
		return nil
	}
	
	novo_node := &Node{
		value: value,
		next: nil,
		prev: nil,
	}

	var aux *Node 
	if index <= lista.Tamanho()-1-index{
		aux = lista.head
		for i := 0; i<index-1; i++{
			aux = aux.next
		}
		novo_node.next = aux.next
		aux.next.prev = novo_node
		aux.next = novo_node
	} else {
		aux = lista.tail
		for i := lista.Tamanho()-1; i>index; i--{
			aux = aux.prev
		}
		novo_node.prev = aux.prev
		aux.prev.next = novo_node
		aux.prev = novo_node
	}
	lista.inserted++
	return nil
}

// Remove um indice da doublelinkedlist
func (lista *DoubleLinkedList) Remove(index int) error {
	if index < 0 || index >= lista.inserted {
		return errors.New(fmt.Sprintf("Index out of range: %d", index))
	} 
	if lista.Tamanho() == 1{
		lista.head = nil
		lista.tail = nil
		lista.inserted--
		return nil
	}
	if index == 0{
		lista.head = lista.head.next
		lista.head.prev = nil
	} else if index == lista.Tamanho()-1{
		lista.tail = lista.tail.prev
		lista.tail.next = nil
	} else {
		var aux *Node
		if index <= lista.Tamanho()-1-index{
			aux = lista.head
			for i := 0; i<index-1; i++{
				aux = aux.next
			}
			aux.next = aux.next.next
			aux.next.prev = aux
		} else {
			aux = lista.tail
			for i := lista.Tamanho()-1; i > index; i--{
				aux = aux.prev
			}
			aux.prev = aux.prev.prev
			aux.prev.next = aux
		}
	} 
	lista.inserted--
	return nil	
}

// Troca o valor guardado na posicao index por value
func (lista *DoubleLinkedList) Set(index int, value int) error {
	if index < 0 || index >= lista.Tamanho() {
		return errors.New(fmt.Sprintf("Index out of range: %d", index))
	}
	var aux *Node
	if index <= lista.Tamanho()-1-index{
		aux = lista.head
		for i := 0; i<index; i++{
			aux = aux.next
		}
	} else {
		aux = lista.tail
		for i := lista.Tamanho()-1; i>index; i--{
			aux = aux.prev
		}
	}
	aux.value = value
	return nil
}

// Remove o ultimo elemento da doublelinkedlist
func (lista *DoubleLinkedList) Pop() error {
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
	for i := 0; i<lista.Tamanho()-2; i++{
		aux = aux.next
	}
	aux.next = nil
	lista.inserted--
	return nil
}

// Inverte uma doublelinkedlist
func (lista *DoubleLinkedList) Invert() error{
	if lista.Tamanho() == 0{
		return errors.New(fmt.Sprintf("Lista possui tamanho igual a 0.\n"))
	}
	aux := lista.head // Node atual
	prox := aux.next  // Salva o proximo node para nao perder a referencia do resto da lista
	aux.next = nil	  // Inverte o ponteiro do node atual
	ant := aux        // Salva o node atual para que o ponteiro do proximo node aponte para este
	for i := 0; i<lista.Tamanho()-1; i++{
		aux = prox       // Atualiza o node atual
		prox = aux.next  // Salva o proximo node para nao perder a referencia do resto da lista
		aux.next = ant   // Inverte o ponteiro do node atual
		ant = aux        // Salva o node atual para que o ponteiro do proximo node aponte para este
	}
	lista.head = aux  // O node atual vira o head da doublelinkedlist
	return nil
}

// Escreve a doublelinkedlist na tela
func (lista *DoubleLinkedList) Print() {
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
	l := &DoubleLinkedList{}

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

	// Invert (inverte a doublelinkedlist)
	l.Invert()
	l.Print()

    // Tamanho final
    fmt.Printf("Tamanho final da lista: %d\n", l.Tamanho())
}
