package main

import (
	"errors"
	"fmt"
)

type IArrayQueue interface {
	Enqueue(val int)
	Dequeue() (int, error)
	Size() int
	Front() (int, error)
	IsEmpty() bool
}

type Node struct {
	value int
	next *Node
}

type LinkedQueue struct {
	front *Node
	rear *Node
	inserted int
}

func (fila *LinkedQueue) Enqueue(val int){
	new_node := &Node{
		value: val,
		next: nil,
	}
	if fila.Size() == 0 {
		fila.front = new_node
		fila.rear = new_node
	} else {
		fila.rear.next = new_node
		fila.rear = new_node
	}
	fila.inserted++
}

func (fila *LinkedQueue) Dequeue() (int, error){
	if fila.Size() == 0{
		return -1, errors.New("Lista vazia.")
	}
	val := fila.front.value
	if fila.Size() == 1{
		fila.front = nil
		fila.rear = nil
	} else {
		fila.front = fila.front.next
	}
	fila.inserted--
	return val, nil
}

func (fila *LinkedQueue) Size() int {
	return fila.inserted
}

func (fila *LinkedQueue) Front() (int, error) {
	if fila.Size() == 0 {
		return -1, errors.New("Lista vazia.")
	}
	return fila.front.value, nil
}

func (fila *LinkedQueue) IsEmpty() bool {
	return fila.Size() == 0
}

func (fila *LinkedQueue) Print(){
	aux := fila.front
	for i:=0; i<fila.Size(); i++{
		fmt.Printf("%d ", aux.value)
		aux = aux.next
	}
	fmt.Printf("\n")
}

func main() {
	fila := &LinkedQueue{}

	fmt.Println("Fila está vazia?", fila.IsEmpty())


	//Insere elementos
	fmt.Printf("Enqueue: 10\n")
	fila.Enqueue(10)
	fila.Print()
	
	fmt.Printf("Enqueue: 20\n")
	fila.Enqueue(20)
	fila.Print()
	
	fmt.Printf("Enqueue: 30\n")
	fila.Enqueue(30)
	fila.Print()


	// Imprime o front
	front, _ := fila.Front()
	fmt.Println("Front:", front) // Esperado: 10

	// Remove um elemento da fila
	val, _ := fila.Dequeue()
	fmt.Println("Dequeue:", val) // Esperado: 10
	fila.Print()


	// Adiciona elementos ao final da fila
	fmt.Printf("Enqueue: 40\n")
	fila.Enqueue(40)
	fila.Print()

	fmt.Printf("Enqueue: 50\n")
	fila.Enqueue(50)
	fila.Print()

	// Imprime o tamanho da fila e se ela esta vazia
	fmt.Println("Size:", fila.Size()) // Esperado: 4
	fmt.Println("Fila está vazia?", fila.IsEmpty()) // Esperado: false

	// Desenfileira todos os elementos da fila até que a mesma fique vazia
	for !fila.IsEmpty() {
		val, _ := fila.Dequeue()
		fmt.Println("Removido:", val)
		fila.Print()
	}
	fmt.Println("Fila está vazia?", fila.IsEmpty()) // Esperado: true
}