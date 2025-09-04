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

type ArrayList struct {
	vetor    []int
	inserted int
}

// Inicializa o arraylist com um tamanho especifico
func (lista *ArrayList) Init(size int) {
	lista.vetor = make([]int, size)
}

// Retorna o tamanho atual do arraylist
func (lista *ArrayList) Tamanho() int {
	return lista.inserted
}

// Informa qual valor está armazenado no índice consultado do arraylist
func (lista *ArrayList) Get(index int) (int, error) {
	if index < 0 || index >= lista.inserted {
		return -1, fmt.Errorf("Index out of range: %d", index)
	}
	return lista.vetor[index], nil
}

// Dobra a capacidade máxima do arraylist
func (lista *ArrayList) doubleV() {
	newV := make([]int, len(lista.vetor)*2)
	for i := 0; i < len(lista.vetor); i++ {
		newV[i] = lista.vetor[i]
	}
	lista.vetor = newV
}

// Adiciona elementos ao final do arraylist
func (lista *ArrayList) Append(value int) {
	if lista.inserted == len(lista.vetor) {
		lista.doubleV()
	}
	lista.vetor[lista.inserted] = value
	lista.inserted++
}

// Adiciona elementos ao arraylist no índice informado
func (lista *ArrayList) Insert(index int, value int) error {
	if index < 0 || index > lista.inserted {
		return fmt.Errorf("Index out of range: %d", index)
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

// "Apaga" um índice do arraylist
func (lista *ArrayList) Remove(index int) error {
	if index < 0 || index >= lista.inserted {
		return fmt.Errorf("Index out of range: %d", index)
	}
	for i := index; i < lista.inserted-1; i++ {
		lista.vetor[i] = lista.vetor[i+1]
	}
	lista.inserted--
	return nil
}

// Troca o valor armazenado em um indice pelo valor fornecido
func (lista *ArrayList) Set(index int, value int) error {
	if index < 0 || index >= lista.Tamanho() {
		return fmt.Errorf("Index out of range: %d", index)
	}
	lista.vetor[index] = value
	return nil
}

// Apaga o ultimo elemento do arraylist
func (lista *ArrayList) Pop() (int, error) {
	if lista.Tamanho() <= 0 {
		return -1, fmt.Errorf("Lista já possui tamanho igual a 0.\n")
	}
	lista.inserted--
	return lista.vetor[lista.inserted], nil
}

// Escreve o arraylist na tela
func (lista *ArrayList) Print() {
	for i := 0; i < lista.Tamanho(); i++ {
		fmt.Printf("%d ", lista.vetor[i])
	}
	fmt.Printf("\n")
}

// Testes com as funções criadas
func main() {
	l := &ArrayList{}
	l.Init(10)

	// Append
	for i := 0; i < 50; i++ {
		l.Append(i)
	}
	fmt.Println("Depois do Append:")
	l.Print()

	// Insert
	l.Insert(2, 10)
	fmt.Println("Depois do Insert(2, 10):")
	l.Print()

	// Remove
	l.Remove(2)
	fmt.Println("Depois do Remove(2):")
	l.Print()

	// Set
	l.Set(0, -10)
	fmt.Println("Depois do Set(0, -10):")
	l.Print()

	// Remove índice 0
	l.Remove(0)
	fmt.Println("Depois do Remove(0):")
	l.Print()

	// Get índice 0
	x, _ := l.Get(0)
	fmt.Printf("Valor do índice 0: %d\n", x)

	// Pop
	val, err := l.Pop()
	if err != nil {
		fmt.Println("Erro no Pop():", err)
	} else {
		fmt.Printf("Pop() removeu: %d\n", val)
	}
	fmt.Println("Depois do Pop():")
	l.Print()

	// Remove vários elementos
	tam := l.Tamanho()
	for i := 10; i < 40; i++ {
		l.Remove(i)
		i--
		if l.Tamanho() == tam-30 {
			break
		}
	}
	fmt.Println("Depois de remover elementos 10 a 39:")
	l.Print()

	// Tamanho final
	fmt.Printf("Tamanho final do array: %d\n", l.Tamanho())
}
