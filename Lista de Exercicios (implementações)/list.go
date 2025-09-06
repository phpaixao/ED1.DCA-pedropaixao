package main

import (
	"fmt"
	"errors"
)

type IList interface {
	Add(value int)							//Feito
	AddOnIndex(value int, index int) error	//Feito
	RemoveOnIndex(index int) error			//Feito
	Get(index int) (int, error)				//Feito
	Set(value int, index int) error			//Feito
	Size() int								//Feito
}

type ArrayList struct{
	v [] int
	inserted int
}

type Node1 struct{
	value int
	next *Node1
}

type Node2 struct{
	value int
	next *Node2
	prev *Node2
}

type LinkedList struct{
	head *Node1
	inserted int
}

type DoubleLinkedList struct{
	head *Node2
	tail *Node2
	inserted int
}

func (list *ArrayList) Init(size int){
	list.v = make([] int, size)
}

func (list *ArrayList) doubleV(){
	newV := make([] int, len(list.v)*2)
	for i:=0; i<len(list.v); i++{
		newV[i] = list.v[i]
	}
	list.v = newV
}

func (list *ArrayList) Add(value int){
	if list.Size() == len(list.v){
		list.doubleV()
	}
	list.v[list.inserted] = value
	list.inserted++
}

func (list *LinkedList) Add(value int){
	new_node := &Node1{
		value: value,
		next: nil,
	}
	if list.Size() == 0{
		list.head = new_node
	} else {
		aux := list.head
		for i:=0; i<list.Size()-1; i++{
			aux = aux.next
		}
		aux.next = new_node
	}
	list.inserted++
}

func (list *DoubleLinkedList) Add(value int){
	new_node := &Node2{
		value: value,
		next: nil,
		prev: nil,
	}
	if list.Size() == 0{
		list.head, list.tail = new_node, new_node
	} else {
		new_node.prev = list.tail
		list.tail.next = new_node
		list.tail = new_node
	}
	list.inserted++
}

func (list *ArrayList) AddOnIndex(value int, index int) error {
	if index<0 || index > list.Size(){
		return errors.New("Invalid index.")
	} 
	if len(list.v) == list.Size(){
		list.doubleV()
	}
	for i:=list.Size(); i>index; i--{
		list.v[i] = list.v[i-1]
	}
	list.v[index] = value
	list.inserted++
	return nil
}

func (list *LinkedList) AddOnIndex(value int, index int) error {
	if index < 0 || index > list.Size(){
		return errors.New("Invalid index.")
	}
	new_node := &Node1{
		value: value,
		next: nil,
	}
	if list.Size() == 0{
		list.head = new_node
		list.inserted++
		return nil
	}
	aux := list.head
	for i:=0; i<index-1; i++{
		aux = aux.next
	}
	new_node.next = aux.next
	aux.next = new_node
	list.inserted++
	return nil
}

func (list *DoubleLinkedList) AddOnIndex(value int, index int) error {
	if index < 0 || index > list.Size() {
		return errors.New("Invalid Index.")
	}
	
	new_node := &Node2{
		value: value,
		next: nil,
		prev: nil,
	}

	if list.Size() == 0{
		list.head, list.tail = new_node, new_node
		list.inserted++
		return nil
	}
	if index == 0{
		new_node.next = list.head
		list.head.prev = new_node
		list.head = new_node
		list.inserted++
		return nil
	}
	if index == list.Size(){
		new_node.prev = list.tail
		list.tail.next = new_node
		list.tail = new_node
		list.inserted++
		return nil
	}

	if index <= (list.Size()-1)/2{
		aux := list.head
		for i:=0; i<index-1; i++{
			aux = aux.next
		}
		new_node.next = aux.next
		new_node.prev = aux
		aux.next.prev = new_node
		aux.next = new_node
	} else {
		aux := list.tail
		for i:=list.Size()-1; i>index; i--{
			aux = aux.prev
		}
		new_node.prev = aux.prev
		new_node.next = aux
		aux.prev.next = new_node
		aux.prev = new_node
	}
	list.inserted++
	return nil
}

func (list *ArrayList) RemoveOnIndex(index int) error {
	if index < 0 || index > list.Size()-1{
		return errors.New("Invalid index.")
	}
	for i:=index; i<list.Size()-1; i++{
		list.v[i] = list.v[i+1]
	}
	list.inserted--
	return nil
}

func (list *LinkedList) RemoveOnIndex(index int) error {
	if index < 0 || index > list.Size()-1{
		return errors.New("Invalid index.")
	}
	if index == 0{
		list.head = list.head.next
		list.inserted--
		return nil
	}
	aux := list.head
	for i:=0; i<index-1; i++{
		aux = aux.next
	}
	aux.next = aux.next.next
	list.inserted--
	return nil
}

func (list *DoubleLinkedList) RemoveOnIndex(index int) error {
	if index < 0 || index >= list.Size(){
		return errors.New("Invalid index.")
	}
	if list.Size() == 1{
		list.head, list.tail = nil, nil
		list.inserted--
		return nil
	}
	if index == 0{
		list.head = list.head.next
		list.head.prev = nil
		list.inserted--
		return nil
	}
	if index == list.Size()-1{
		list.tail = list.tail.prev
		list.tail.next = nil
		list.inserted--
		return nil
	}
	if index <= (list.Size()-1)/2 {
		aux := list.head
		for i:=0; i<index-1; i++{
			aux = aux.next
		}
		aux.next = aux.next.next
		aux.next.prev = aux
	} else {
		aux := list.tail
		for i:=list.Size()-1; i>index+1; i--{
			aux = aux.prev
		}
		aux.prev = aux.prev.prev
		aux.prev.next = aux
	}
	list.inserted--
	return nil
}

func (list *ArrayList) Get(index int) (int, error){
	if index < 0 || index >= list.Size(){
		return -1, errors.New("Invalid index.")
	}
	return list.v[index], nil
}

func (list *LinkedList) Get(index int) (int, error){
	if index < 0 || index >= list.Size(){
		return -1, errors.New("Invalid index.")
	}
	aux := list.head
	for i:=0; i<index; i++{
		aux = aux.next
	}
	return aux.value, nil
}

func (list *DoubleLinkedList) Get(index int) (int, error){
	if index < 0 || index >= list.Size(){
		return -1, errors.New("Invalid index.")
	}
	var aux *Node2
	if index <= (list.Size()-1)/2{
		aux = list.head
		for i:=0; i<index; i++{
			aux = aux.next
		}
	} else {
		aux = list.tail
		for i:=list.Size()-1; i>index; i--{
			aux = aux.prev
		}
	}
	return aux.value, nil
}

func (list *ArrayList) Set(value int, index int) error {
	if index < 0 || index >= list.Size(){
		return errors.New("Invalid index.")
	}
	list.v[index] = value
	return nil
}

func (list *LinkedList) Set(value int, index int) error {
	if index < 0 || index >= list.Size(){
		return errors.New("Invalid index.")
	}
	aux := list.head
	for i:=0; i<index; i++{
		aux = aux.next
	}
	aux.value = value
	return nil
}

func (list *DoubleLinkedList) Set(value int, index int) error {
	if index < 0 || index >= list.Size(){
		return errors.New("Invalid index.")
	}
	var aux *Node2
	if index <= (list.Size()-1)/2{
		aux = list.head
		for i:=0; i<index; i++{
			aux = aux.next
		}
	} else {
		aux = list.tail
		for i:=list.Size()-1; i>index; i--{
			aux = aux.prev
		}
	}
	aux.value = value
	return nil
}

func (list *ArrayList) Size() int {
	return list.inserted
}

func (list *LinkedList) Size() int {
	return list.inserted
}

func (list *DoubleLinkedList) Size() int {
	return list.inserted
}

func (list *ArrayList) Print() {
	for i:=0; i<list.Size(); i++{
		fmt.Printf("%d ", list.v[i])
	}
	fmt.Println()
}

func (list *LinkedList) Print() {
	aux := list.head
	for i:=0; i<list.Size(); i++{
		fmt.Printf("%d ", aux.value)
		aux = aux.next
	}
	fmt.Println()
}

func (list *DoubleLinkedList) Print() {
	aux := list.head
	for i:=0; i<list.Size(); i++{
		fmt.Printf("%d ", aux.value)
		aux = aux.next
	}
	fmt.Println()
}

func main() {
	fmt.Println("=== Testando ArrayList ===")
	arr := &ArrayList{}
	arr.Init(3)

	arr.Add(10)
	arr.Add(20)
	arr.Add(30)
	fmt.Print("ArrayList após Add: ")
	arr.Print()

	arr.AddOnIndex(15, 1)
	fmt.Print("ArrayList após AddOnIndex(15, 1): ")
	arr.Print()

	arr.Set(25, 2)
	fmt.Print("ArrayList após Set(25, 2): ")
	arr.Print()

	val, _ := arr.Get(2)
	fmt.Println("Get(2):", val)

	arr.RemoveOnIndex(1)
	fmt.Print("ArrayList após RemoveOnIndex(1): ")
	arr.Print()

	fmt.Println("Tamanho ArrayList:", arr.Size())

	fmt.Println("\n=== Testando LinkedList ===")
	ll := &LinkedList{}
	ll.Add(10)
	ll.Add(20)
	ll.Add(30)
	fmt.Print("LinkedList após Add: ")
	ll.Print()

	ll.AddOnIndex(15, 1)
	fmt.Print("LinkedList após AddOnIndex(15, 1): ")
	ll.Print()

	ll.Set(25, 2)
	fmt.Print("LinkedList após Set(25, 2): ")
	ll.Print()

	val, _ = ll.Get(2)
	fmt.Println("Get(2):", val)

	ll.RemoveOnIndex(1)
	fmt.Print("LinkedList após RemoveOnIndex(1): ")
	ll.Print()

	fmt.Println("Tamanho LinkedList:", ll.Size())

	fmt.Println("\n=== Testando DoubleLinkedList ===")
	dll := &DoubleLinkedList{}
	dll.Add(10)
	dll.Add(20)
	dll.Add(30)
	dll.Add(40)
	dll.Add(50)
	fmt.Print("DoubleLinkedList após Add: ")
	dll.Print()

	// Testando AddOnIndex no meio e no fim para forçar otimização
	dll.AddOnIndex(15, 1)  // percorre a partir do head
	dll.AddOnIndex(45, 6)  // percorre a partir do tail
	fmt.Print("DoubleLinkedList após AddOnIndex(15,1) e AddOnIndex(45,6): ")
	dll.Print()

	// Testando Set usando índices do fim
	dll.Set(25, 2)   // perto do head
	dll.Set(48, 6)   // perto do tail
	fmt.Print("DoubleLinkedList após Set(25,2) e Set(48,6): ")
	dll.Print()

	// Testando Get usando índices do fim
	val, _ = dll.Get(2)  // do meio, vai do head
	fmt.Println("Get(2):", val)
	val, _ = dll.Get(6)  // do final, vai do tail
	fmt.Println("Get(6):", val)

	// Testando RemoveOnIndex usando índices do fim
	dll.RemoveOnIndex(1)  // próximo ao head
	dll.RemoveOnIndex(5)  // próximo ao tail
	fmt.Print("DoubleLinkedList após RemoveOnIndex(1) e RemoveOnIndex(5): ")
	dll.Print()

	fmt.Println("Tamanho DoubleLinkedList:", dll.Size())
}
