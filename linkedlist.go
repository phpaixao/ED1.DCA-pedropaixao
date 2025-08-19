package main

import (
	"errors"
	"fmt"
)

type List interface {
	Tamanho() int                      //Feito
	Get(index int) (int, error)        //Feito
	Append(value int)                  //Feito
	Insert(index int, value int) error //Feito
	Remove(index int) error            //Feito
}

/*
type ArrayList struct {
	vetor    []int
	inserted int
}
*/

type Node struct {
	value int
	next * Node
}

type LinkedList struct {
	head * Node
	inserted int
}

func (lista *LinkedList) Tamanho() int {
	return lista.inserted
}

func (lista *LinkedList) Get(index int) (int, error) {
	if index < 0 || index >= lista.inserted {
		return -1, errors.New(fmt.Sprintf("Index out of range: %d", index))
	}
	aux := lista.head
	for i:=0; i<index; i++ {
		aux = aux.next
	}
	return aux.value. nil
}

func (lista *ArrayList) doubleV() {
	newV := make([]int, len(lista.vetor)*2)
	for i := 0; i < len(lista.vetor); i++ {
		newV[i] = lista.vetor[i]
	}
	lista.vetor = newV
}

func (lista *ArrayList) Append(value int) {
	if lista.inserted == len(lista.vetor) {
		lista.doubleV()
	}
	lista.vetor[lista.inserted] = value
	lista.inserted++
}

func (lista *ArrayList) Insert(index int, value int) error {
	if index < 0 || index > lista.inserted {
		return errors.New(fmt.Sprintf("Index out of range: %d", index))
	}
	if lista.inserted == len(lista.vetor) {
		lista.doubleV()
	}
	for i := lista.inserted; i > index; i-- {
		lista.vetor[i] = lista.vetor[i-1]
	}
	lista.vetor[index] = value
	lista.inserted++
	return nil
}

func (lista *ArrayList) Remove(index int) error {
	if index < 0 || index >= lista.inserted {
		return errors.New(fmt.Sprintf("Index out of range: %d", index))
	}
	for i := index; i < lista.inserted-1; i++ {
		lista.vetor[i] = lista.vetor[i+1]
	}
	lista.inserted--
	return nil
}

func (lista *ArrayList) Set(index int, value int) error {
	if index < 0 || index >= lista.Tamanho() {
		return errors.New(fmt.Sprintf("Index out of range: %d", index))
	}
	lista.vetor[index] = value
	return nil
}

func (lista *ArrayList) Pop() error {
	if lista.Tamanho() <= 0 {
		return errors.New(fmt.Sprintf("Lista já possui tamanho igual a 0.\n"))
	}
	lista.inserted--
	return nil
}

func (lista *ArrayList) Print() {
	for i := 0; i < lista.Tamanho(); i++ {
		fmt.Printf("%d ", lista.vetor[i])
	}
	fmt.Printf("\n")
}

func main() {
	l := &ArrayList{}
	l.Init(10)
	for i := 0; i < 50; i++ {
		l.Append(i)
	}
	l.Print()

	l.Insert(2, 10)
	l.Print()

	l.Remove(2)
	l.Print()

	l.Set(0, -10)
	l.Print()

	l.Remove(0)
	l.Print()

	x, _ := l.Get(0)
	fmt.Printf("%d \n", x)

	tam := l.Tamanho()
	for i := 10; i < 40; i++ {
		l.Remove(i)
		i--

		if l.Tamanho() == tam-30 {
			break
		}
	}
	l.Print()

	tam = l.Tamanho()
	fmt.Printf("\nO tamanho do array eh: %d", tam)
}
