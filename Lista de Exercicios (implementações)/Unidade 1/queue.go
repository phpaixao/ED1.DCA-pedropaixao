package main

import (
	"errors"
	"fmt"
)

type IQueue interface {
	Enqueue(value int)
	Dequeue() (int, error)
	Front() (int, error)
	IsEmpty() bool
	Size() int
}

type ArrayQueue struct{
	front int
	rear int
	v[] int
}

type Node struct{
	value int
	next *Node
}

type LinkedListQueue struct{
	front *Node
	rear *Node
	inserted int
}

func (queue *ArrayQueue) Init(size int){
	queue.v = make([] int, size)
	queue.rear = -1
	queue.front = -1
}

func (queue *ArrayQueue) doubleV(){
	newV := make([] int, len(queue.v)*2)
	for i := queue.front; i<(len(queue.v) + queue.front); i++{
		newV[i - queue.front] = queue.v[i % len(queue.v)]
	}
	queue.v = newV
	queue.front = 0
	queue.rear = queue.Size() - 1
}

func (queue *ArrayQueue) Enqueue(value int){
	if len(queue.v) == queue.Size(){
		queue.doubleV()
	}
	if queue.Size() == 0{
		queue.front++
		queue.rear++
	} else if queue.rear + 1 < len(queue.v){
		queue.rear++
	} else if queue.rear + 1 == len(queue.v){
		queue.rear = 0
	}
	queue.v[queue.rear] = value
}

func (queue *LinkedListQueue) Enqueue(value int){
	new_node := &Node{
		value: value,
		next: nil,
	}
	if queue.Size() == 0{
		queue.front = new_node
		queue.rear = new_node
	} else {
		queue.rear.next = new_node
		queue.rear = new_node
	}
	queue.inserted++
}

func (queue *ArrayQueue) Dequeue() (int, error){
	if queue.Size() == 0{
		return -1, errors.New("Fila vazia.")
	}
	val := queue.v[queue.front]
	if queue.Size() == 1{
		queue.front = -1
		queue.rear = -1
	} else if queue.front + 1 < len(queue.v){
		queue.front++
	} else if queue.front + 1 == len(queue.v){
		queue.front = 0
	}
	return val, nil	
}

func (queue *LinkedListQueue) Dequeue() (int, error){
	if queue.Size() == 0{
		return -1, errors.New("Fila vazia.")
	}
	val := queue.front.value
	if queue.Size() == 1{
		queue.front = nil
		queue.rear = nil
	} else {
		queue.front = queue.front.next
	}
	queue.inserted--
	return val, nil
}

func (queue *ArrayQueue) Front() (int, error){
	if queue.Size() == 0{
		return -1, errors.New("Fila vazia.")
	}
	return queue.v[queue.front], nil
}

func (queue *LinkedListQueue) Front() (int, error){
	if queue.Size() == 0{
		return -1, errors.New("Fila vazia.")
	}
	return queue.front.value, nil
}

func (queue *ArrayQueue) IsEmpty() bool{
	return queue.Size() == 0
}

func (queue *LinkedListQueue) IsEmpty() bool{
	return queue.Size() == 0
}

func (queue *ArrayQueue) Size() int{
	if queue.front == -1 && queue.rear == -1{
		return 0
	}
	if queue.front <= queue.rear {
		return queue.rear - queue.front + 1
	}
	return len(queue.v) + queue.rear - queue.front + 1
}

func (queue *LinkedListQueue) Size() int{
	return queue.inserted
}

func (queue *ArrayQueue) Print(){
	for i := queue.front; i<queue.Size() + queue.front; i++{
		fmt.Printf("%d ", queue.v[i%len(queue.v)])
	}
	fmt.Println()
}

func (queue *LinkedListQueue) Print(){
	aux := queue.front
	for i:=0; i<queue.Size(); i++{
		fmt.Printf("%d ", aux.value)
		aux = aux.next
	}
	fmt.Println()
}


func main() {
    // ---------------- ArrayQueue ----------------
    fmt.Println("=== Testando ArrayQueue ===")
    aq := &ArrayQueue{}
    aq.Init(3)

    fmt.Println("Enfileirando: 1, 2, 3 (capacidade cheia)")
    aq.Enqueue(1)
    aq.Enqueue(2)
    aq.Enqueue(3)
    aq.Print()

    fmt.Println("Removendo dois elementos (para circularidade)")
    v, _ := aq.Dequeue()
    fmt.Println("Dequeue():", v)
    v, _ = aq.Dequeue()
    fmt.Println("Dequeue():", v)
    aq.Print()

    fmt.Println("Enfileirando: 4, 5 (usa espaço liberado, fila circular)")
    aq.Enqueue(4)
    aq.Enqueue(5)
    aq.Print()

    fmt.Println("Forçando doubleV (capacidade inicial = 3, vamos enfileirar mais)")
    aq.Enqueue(6) // aqui deve dobrar a capacidade
    aq.Enqueue(7)
    aq.Enqueue(8)
    aq.Print()
    fmt.Println("Capacidade após doubleV:", len(aq.v))
    fmt.Println("Tamanho da fila:", aq.Size())

    fmt.Println("Esvaziando completamente ArrayQueue:")
    for !aq.IsEmpty() {
        v, _ := aq.Dequeue()
        fmt.Println("Dequeue():", v)
    }
    aq.Print()

    fmt.Println("Reutilizando fila após esvaziar:")
    aq.Enqueue(100)
    aq.Enqueue(200)
    aq.Print()

    // ---------------- LinkedListQueue ----------------
    fmt.Println("\n=== Testando LinkedListQueue ===")
    lq := &LinkedListQueue{}

    fmt.Println("Enfileirando: 1, 2, 3")
    lq.Enqueue(1)
    lq.Enqueue(2)
    lq.Enqueue(3)
    lq.Print()

    fmt.Println("Removendo um elemento")
    v, _ = lq.Dequeue()
    fmt.Println("Dequeue():", v)
    lq.Print()

    fmt.Println("Enfileirando: 4, 5, 6")
    lq.Enqueue(4)
    lq.Enqueue(5)
    lq.Enqueue(6)
    lq.Print()
    fmt.Println("Tamanho da fila:", lq.Size())

    fmt.Println("Esvaziando completamente LinkedListQueue:")
    for !lq.IsEmpty() {
        v, _ := lq.Dequeue()
        fmt.Println("Dequeue():", v)
    }
    lq.Print()

    fmt.Println("Reutilizando fila após esvaziar:")
    lq.Enqueue(1000)
    lq.Enqueue(2000)
    lq.Print()
}
