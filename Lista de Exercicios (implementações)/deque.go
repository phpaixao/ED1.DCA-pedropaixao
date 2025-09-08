package main

import (
	"errors"
	"fmt"
)

type IDeque interface {
	EnqueueFront(value int)			
	EnqueueRear(value int)			
	DequeueFront() (int, error)		
	DequeueRear() (int, error)
	Front() (int, error)
	Rear() (int, error)
	IsEmpty() bool
	Size() int
}

type ArrayDeque struct{
	front int
	rear int
	v[] int
}

type Node struct{
	value int
	next *Node
	prev *Node
}

type DoubleLinkedListDeque struct{
	front *Node
	rear *Node
	inserted int
}

func (deque *ArrayDeque) Init(size int){
	deque.v = make([] int, size)
	deque.rear = -1
	deque.front = -1
}

func (deque *ArrayDeque) doubleV(){
	newV := make([] int, len(deque.v)*2)
	for i := deque.front; i<(len(deque.v) + deque.front); i++{
		newV[i - deque.front] = deque.v[i % len(deque.v)]
	}
	deque.v = newV
	deque.front = 0
	deque.rear = deque.Size() - 1
}

func (deque *ArrayDeque) EnqueueFront(value int){
	if len(deque.v) == deque.Size(){
		deque.doubleV()
	}
	if deque.Size() == 0{
		deque.front++
		deque.rear++
	} else if deque.front - 1 > -1{
		deque.front--
	} else if deque.front - 1 == -1{
		deque.front = len(deque.v) - 1
	}
	deque.v[deque.front] = value
}

func (deque *DoubleLinkedListDeque) EnqueueFront(value int){
	new_node := &Node{
		value: value,
		next: nil,
		prev: nil,
	}

	if deque.Size() == 0{
		deque.rear = new_node
	} else {
		new_node.next = deque.front
		deque.front.prev = new_node
	}
	deque.front = new_node
	deque.inserted++
}

func (deque *ArrayDeque) EnqueueRear(value int){
	if len(deque.v) == deque.Size(){
		deque.doubleV()
	}
	if deque.Size() == 0{
		deque.front++
		deque.rear++
	} else if deque.rear + 1 < len(deque.v){
		deque.rear++
	} else if deque.rear + 1 == len(deque.v){
		deque.rear = 0
	}
	deque.v[deque.rear] = value
}

func (deque *DoubleLinkedListDeque) EnqueueRear(value int){
	new_node := &Node{
		value: value,
		next: nil,
		prev: nil,
	}

	if deque.Size() == 0{
		deque.front = new_node
	} else {
		deque.rear.next = new_node
		new_node.prev = deque.rear
	}
	deque.rear = new_node
	deque.inserted++
}

func (deque *ArrayDeque) DequeueFront() (int, error){
	if deque.Size() == 0{
		return -1, errors.New("Fila vazia.")
	}
	val := deque.v[deque.front]
	if deque.Size() == 1{
		deque.front = -1
		deque.rear = -1
	} else if deque.front + 1 < len(deque.v){
		deque.front++
	} else if deque.front + 1 == len(deque.v){
		deque.front = 0
	}
	return val, nil	
}

func (deque *DoubleLinkedListDeque) DequeueFront() (int, error){
	if deque.Size() == 0{
		return -1, errors.New("Deque vazio.")
	}
	val := deque.front.value
	if deque.Size() == 1{
		deque.front = nil
		deque.rear = nil
	} else {
		deque.front = deque.front.next
		deque.front.prev = nil
	}
	deque.inserted--
	return val, nil	
}

func (deque *ArrayDeque) DequeueRear() (int, error){
	if deque.Size() == 0{
		return -1, errors.New("Deque vazio.")
	}
	val := deque.v[deque.rear]
	if deque.Size() == 1{
		deque.front = -1
		deque.rear = -1
		return val, nil
	}
	if deque.rear - 1 > -1{
		deque.rear--
	} else if deque.rear - 1 == -1{
		deque.rear = len(deque.v) - 1
	}
	return val, nil
}

func (deque *DoubleLinkedListDeque) DequeueRear() (int, error){
	if deque.Size() == 0{
		return -1, errors.New("Deque vazio.")
	}
	val := deque.rear.value
	if deque.Size() == 1{
		deque.front = nil
		deque.rear = nil
	} else {
		deque.rear = deque.rear.prev
		deque.rear.next = nil
	}
	deque.inserted--
	return val, nil
}

func (deque *ArrayDeque) Front() (int, error){
	if deque.Size() == 0{
		return -1, errors.New("Fila vazia.")
	}
	return deque.v[deque.front], nil
}

func (deque *DoubleLinkedListDeque) Front() (int, error){
	if deque.Size() == 0{
		return -1, errors.New("Fila vazia.")
	}
	return deque.front.value, nil
}

func (deque *ArrayDeque) Rear() (int, error){
	if deque.Size() == 0{
		return -1, errors.New("Fila vazia.")
	}
	return deque.v[deque.rear], nil
}

func (deque *DoubleLinkedListDeque) Rear() (int, error){
	if deque.Size() == 0{
		return -1, errors.New("Fila vazia.")
	}
	return deque.rear.value, nil
}

func (deque *ArrayDeque) IsEmpty() bool{
	return deque.Size() == 0
}

func (deque *DoubleLinkedListDeque) IsEmpty() bool{
	return deque.Size() == 0
}

func (deque *ArrayDeque) Size() int{
	if deque.front == -1 && deque.rear == -1{
		return 0
	}
	if deque.front <= deque.rear {
		return deque.rear - deque.front + 1
	}
	return len(deque.v) + deque.rear - deque.front + 1
}

func (deque *DoubleLinkedListDeque) Size() int{
	return deque.inserted
}

func (deque *ArrayDeque) Print(){
	for i := deque.front; i<deque.Size() + deque.front; i++{
		fmt.Printf("%d ", deque.v[i%len(deque.v)])
	}
	fmt.Println()
}

func (deque *DoubleLinkedListDeque) Print(){
	aux := deque.front
	for i:=0; i<deque.Size(); i++{
		fmt.Printf("%d ", aux.value)
		aux = aux.next
	}
	fmt.Println()
}

func main() {
	fmt.Println("=== Testando ArrayDeque ===")
	arr := &ArrayDeque{}
	arr.Init(5)

	fmt.Println("Inserindo 10, 20, 30 no rear")
	arr.EnqueueRear(10)
	arr.EnqueueRear(20)
	arr.EnqueueRear(30)
	fmt.Print("ArrayDeque atual: ")
	arr.Print()

	fmt.Println("Inserindo 5, 0 no front")
	arr.EnqueueFront(5)
	arr.EnqueueFront(0)
	fmt.Print("ArrayDeque atual: ")
	arr.Print()

	fmt.Println("Removendo do front e do rear")
	val, _ := arr.DequeueFront()
	fmt.Println("DequeueFront():", val)
	val, _ = arr.DequeueRear()
	fmt.Println("DequeueRear():", val)
	fmt.Print("ArrayDeque após remoções: ")
	arr.Print()

	fmt.Println("Inserindo 40, 50, 60 para testar expansão (doubleV)")
	arr.EnqueueRear(40)
	arr.EnqueueRear(50)
	arr.EnqueueRear(60)
	fmt.Print("ArrayDeque atual (depois do doubleV): ")
	arr.Print()

	fmt.Println("Verificando Front e Rear")
	front, _ := arr.Front()
	rear, _ := arr.Rear()
	fmt.Println("Front:", front, "Rear:", rear)

	fmt.Println("Esvaziando ArrayDeque")
	for !arr.IsEmpty() {
		val, _ := arr.DequeueFront()
		fmt.Println("DequeueFront():", val)
	}

	fmt.Println("\n=== Testando DoubleLinkedListDeque ===")
	dll := &DoubleLinkedListDeque{}

	fmt.Println("Inserindo 10, 20, 30 no rear")
	dll.EnqueueRear(10)
	dll.EnqueueRear(20)
	dll.EnqueueRear(30)
	fmt.Print("DoubleLinkedListDeque atual: ")
	dll.Print()

	fmt.Println("Inserindo 5, 0 no front")
	dll.EnqueueFront(5)
	dll.EnqueueFront(0)
	fmt.Print("DoubleLinkedListDeque atual: ")
	dll.Print()

	fmt.Println("Removendo do front e do rear")
	val, _ = dll.DequeueFront()
	fmt.Println("DequeueFront():", val)
	val, _ = dll.DequeueRear()
	fmt.Println("DequeueRear():", val)
	fmt.Print("DoubleLinkedListDeque após remoções: ")
	dll.Print()

	fmt.Println("Verificando Front e Rear")
	front, _ = dll.Front()
	rear, _ = dll.Rear()
	fmt.Println("Front:", front, "Rear:", rear)

	fmt.Println("Esvaziando DoubleLinkedListDeque")
	for !dll.IsEmpty() {
		val, _ := dll.DequeueFront()
		fmt.Println("DequeueFront():", val)
	}
}