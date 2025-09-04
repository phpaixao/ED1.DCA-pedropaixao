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

type ArrayQueue struct{
	v [] int
	front int
	rear int
}

func (fila *ArrayQueue) init(size int){
	fila.v = make([] int, size)
	fila.front, fila.rear = -1, -1
}

func (fila *ArrayQueue) doubleV(){
	newV := make([] int, len(fila.v)*2)
	for i:=fila.front; i<(fila.Size() + fila.front); i++{
		newV[i-fila.front] = fila.v[i % len(fila.v)]
	}
	fila.v = newV
	fila.front = 0
	fila.rear = fila.Size() - 1
}

func (fila *ArrayQueue) Enqueue(val int){
	if fila.Size() == len(fila.v){
		fila.doubleV()
	}
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

func (fila *ArrayQueue) IsEmpty() bool {
	return fila.Size() == 0
}

// Print corrigido para fila circular
func (fila *ArrayQueue) Print() {
	for i:=fila.front; i<(fila.Size() + fila.front); i++{
		fmt.Printf("%d ", fila.v[i % len(fila.v)])
	}
	fmt.Println()
}

func main() {
	// cria fila de tamanho 2
	fila := ArrayQueue{}
	fila.init(2)

	// A lista está vazia?
	fmt.Println("Fila está vazia?", fila.IsEmpty()) // Esperado: true

	// insere elementos
	fila.Enqueue(10)
	fila.Enqueue(20)
	fila.Enqueue(30)
	fmt.Print("Fila após inserir 10, 20, 30: ")
	fila.Print()

	// A lista está vazia?
	fmt.Println("Fila está vazia?", fila.IsEmpty()) // Esperado: false

	// mostra tamanho e front
	fmt.Println("Tamanho:", fila.Size()) // esperado: 3
	val, _ := fila.Front()
	fmt.Println("Front:", val) // esperado: 10

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
	for !fila.IsEmpty() {
		val, _ = fila.Dequeue()
		fmt.Println("Dequeued:", val)
	}
	fmt.Print("Fila após remover tudo: ")
	fila.Print()

	// A lista está vazia?
	fmt.Println("Fila está vazia?", fila.IsEmpty()) // Esperado: true

	// tenta remover de lista vazia
	_, err := fila.Dequeue()
	if err != nil {
		fmt.Println("Erro:", err)
	}
}