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
}

type ArrayQueue struct{
	v [] int
	front int
	rear int
}

func (fila *ArrayQueue) Enqueue(val int){
	if fila.front == -1 && fila.rear == -1{
		fila.front++
		fila.rear++
	} else {
		fila.rear = (fila.rear+1) % len(fila.v)
	}
	fila.v[fila.rear] = val
}

func (fila *ArrayQueue) Dequeue() (int, error){
	if fila.front == -1 && fila.rear == -1{
		return -1, errors.New("Lista vazia.")
	}
	val := fila.v[fila.front]
	if fila.front == fila.rear {
		fila.rear, fila.front = -1, -1
	} else {
		fila.front = (fila.front+1) % len(fila.v)
	}
	return val, nil
}

func (fila *ArrayQueue) Size() int{
	if fila.rear == -1 && fila.front == -1{
		return 0
	}
	return (fila.rear - fila.front + len(fila.v)) % len(fila.v) + 1
}

func (fila *ArrayQueue) Front() (int, error){
	if fila.front == -1 && fila.rear == -1 {
		return -1, errors.New("A lista está vazia.")
	}
	return fila.v[fila.front], nil
}

func (fila *ArrayQueue) Print() {
	for i:=0; i<fila.Size(); i++{
		fmt.Printf("%d ", fila.v[i])
	}
}

func main() {
	// cria fila de tamanho 5
	fila := ArrayQueue{
		v:     make([]int, 5),
		front: -1,
		rear:  -1,
	}

	// insere elementos
	fila.Enqueue(10)
	fila.Enqueue(20)
	fila.Enqueue(30)
	fmt.Print("Fila após inserir 10, 20, 30: ")
	fila.Print()

	// mostra tamanho e front
	fmt.Println("Tamanho:", fila.Size()) // esperado: 3
	val, _ := fila.Front()
	fmt.Println("Front:", val)           // esperado: 10

	// remove 2 elementos
	val, _ = fila.Dequeue()
	fmt.Println("Dequeued:", val) 
	val, _ = fila.Dequeue()
	fmt.Println("Dequeued:", val) 
	fmt.Print("Fila após remover 2: ")
	fila.Print()

	// insere mais alguns
	fila.Enqueue(40)
	fila.Enqueue(50)
	fmt.Print("Fila após inserir 40, 50: ")
	fila.Print()

	// mostra o estado final
	fmt.Println("Tamanho:", fila.Size()) 
	val, _ = fila.Front()
	fmt.Println("Front:", val)           

	// remove tudo
	for i := 0; i < 3; i++ {
		val, _ = fila.Dequeue()
		fmt.Println("Dequeued:", val)
	}
	fmt.Print("Fila após remover tudo: ")
	fila.Print()

	// tenta remover de lista vazia
	_, err := fila.Dequeue()
	if err != nil {
		fmt.Println("Erro:", err)
	}
}
