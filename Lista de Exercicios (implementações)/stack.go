package main

import (
	"fmt"
	"errors"
)

type IStack interface{
	Push(value int)
	Pop() (int, error)
	Peek() (int, error)
	IsEmpty() bool
	Size() int
}

type ArrayStack struct{
	v [] int
	top int
}

type Node struct{
	value int
	next *Node
}

type LinkedStack struct{
	top *Node
	inserted int
}

func (stack *ArrayStack) Init(size int){
	stack.v = make([] int, size)
}

func (stack *ArrayStack) doubleV(){
	newV := make([] int, stack.Size()*2)
	for i:=0; i<stack.Size(); i++{
		newV[i] = stack.v[i]
	}
	stack.v = newV
}

func (stack *ArrayStack) Push(value int){
	if len(stack.v) == stack.Size(){
		stack.doubleV()
	}
	stack.v[stack.top] = value
	stack.top++
}

func (stack *LinkedStack) Push(value int){
	new_node := &Node{
		value: value,
		next: nil,
	}
	if stack.IsEmpty(){
		stack.top = new_node
	} else {
		new_node.next = stack.top
		stack.top = new_node
	}
	stack.inserted++
}

func (stack *ArrayStack) Pop() (int, error) {
	if stack.IsEmpty(){
		return -1, errors.New("Pilha está vazia.")
	}
	stack.top--
	return stack.v[stack.top], nil
}

func (stack *LinkedStack) Pop() (int, error) {
	if stack.IsEmpty(){
		return -1, errors.New("Pilha está vazia.")
	}
	val := stack.top.value
	stack.top = stack.top.next
	stack.inserted--
	return val, nil
}

func (stack *ArrayStack) Peek() (int, error){
	if stack.IsEmpty(){
		return -1, errors.New("Pilha está vazia.")
	}
	return stack.v[stack.top-1], nil
}

func (stack *LinkedStack) Peek() (int, error){
	if stack.IsEmpty(){
		return -1, errors.New("Pilha está vazia.")
	}
	return stack.top.value, nil
}

func (stack *ArrayStack) IsEmpty() bool {
	return stack.Size() == 0
}

func (stack *LinkedStack) IsEmpty() bool {
	return stack.Size() == 0
}

func (stack *ArrayStack) Size() int {
	return stack.top
}

func (stack *LinkedStack) Size() int {
	return stack.inserted
}

func (stack *ArrayStack) Print(){
	for i := stack.Size()-1; i>0; i--{
		fmt.Printf("%d ", stack.v[i])
	}
	fmt.Println()
}

func (stack *LinkedStack) Print(){
	aux := stack.top
	for i := 0; i<stack.Size(); i++{
		fmt.Printf("%d ", aux.value)
		aux = aux.next
	}
	fmt.Println()
}

func main() {
	fmt.Println("=== Testando ArrayStack ===")
	arrStack := &ArrayStack{}
	arrStack.Init(2)

	fmt.Println("Empilhando valores: 10, 20, 30")
	arrStack.Push(10)
	arrStack.Push(20)
	arrStack.Push(30)
	fmt.Print("ArrayStack atual: ")
	arrStack.Print()

	top, _ := arrStack.Peek()
	fmt.Println("Topo da pilha (Peek):", top)

	val, _ := arrStack.Pop()
	fmt.Println("Pop():", val)
	fmt.Print("ArrayStack após Pop: ")
	arrStack.Print()

	fmt.Println("Empilhando 40")
	arrStack.Push(40)
	fmt.Print("ArrayStack atual: ")
	arrStack.Print()

	fmt.Println("Tamanho da pilha:", arrStack.Size())
	fmt.Println("Está vazia?", arrStack.IsEmpty())

	fmt.Println("\n=== Testando LinkedStack ===")
	linkedStack := &LinkedStack{}

	fmt.Println("Empilhando valores: 10, 20, 30")
	linkedStack.Push(10)
	linkedStack.Push(20)
	linkedStack.Push(30)
	fmt.Print("LinkedStack atual: ")
	linkedStack.Print()

	top, _ = linkedStack.Peek()
	fmt.Println("Topo da pilha (Peek):", top)

	val, _ = linkedStack.Pop()
	fmt.Println("Pop():", val)
	fmt.Print("LinkedStack após Pop: ")
	linkedStack.Print()

	fmt.Println("Empilhando 40")
	linkedStack.Push(40)
	fmt.Print("LinkedStack atual: ")
	linkedStack.Print()

	fmt.Println("Tamanho da pilha:", linkedStack.Size())
	fmt.Println("Está vazia?", linkedStack.IsEmpty())

	// Testando Pop até a pilha esvaziar
	fmt.Println("\nEsvaziando ArrayStack:")
	for !arrStack.IsEmpty() {
		val, _ := arrStack.Pop()
		fmt.Println("Pop():", val)
	}

	fmt.Println("Esvaziando LinkedStack:")
	for !linkedStack.IsEmpty() {
		val, _ := linkedStack.Pop()
		fmt.Println("Pop():", val)
	}
	
	fmt.Println("\n=== Testando operações em pilhas vazias ===")

	emptyArrayStack := &ArrayStack{}
	emptyArrayStack.Init(2)

	val, err := emptyArrayStack.Pop()
	if err != nil {
		fmt.Println("Pop em ArrayStack vazio retornou erro:", err)
	} else {
		fmt.Println("Pop em ArrayStack vazio retornou:", val)
	}

	val, err = emptyArrayStack.Peek()
	if err != nil {
		fmt.Println("Peek em ArrayStack vazio retornou erro:", err)
	} else {
		fmt.Println("Peek em ArrayStack vazio retornou:", val)
	}

	emptyLinkedStack := &LinkedStack{}

	val, err = emptyLinkedStack.Pop()
	if err != nil {
		fmt.Println("Pop em LinkedStack vazio retornou erro:", err)
	} else {
		fmt.Println("Pop em LinkedStack vazio retornou:", val)
	}

	val, err = emptyLinkedStack.Peek()
	if err != nil {
		fmt.Println("Peek em LinkedStack vazio retornou erro:", err)
	} else {
		fmt.Println("Peek em LinkedStack vazio retornou:", val)
	}
}
