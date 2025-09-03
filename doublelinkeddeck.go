package main

import (
	"errors"
	"fmt"
)

type iDeck interface {
	push_front(val int)
	push_back(val int)
	pull_front() (int, error)
	pull_back() (int, error)
	size() int
	Front() (int, error)
	back() (int, error)
}

type Node struct{
	value int
	next *Node
	prev *Node
}

type DoubleLinkedQueue struct {
	head *Node
	tail *Node
	inserted int
}

func (deck *DoubleLinkedQueue) push_front(val int) {
	new_node := &Node{
		value: val,
		next: nil,
		prev: nil,
	}
	if deck.size() == 0{
		deck.head = new_node
		deck.tail = new_node
	} else {
		new_node.next = deck.head
		deck.head.prev = new_node
		deck.head = new_node
	}
	deck.inserted++
}

func (deck *DoubleLinkedQueue) push_back(val int) {
	new_node := &Node{
		value: val,
		next: nil,
		prev: nil,
	}
	if deck.size() == 0{
		deck.head = new_node
		deck.tail = new_node
	} else {
		new_node.prev = deck.tail
		deck.tail.next = new_node
		deck.tail = new_node
	}
	deck.inserted++
}

func (deck *DoubleLinkedQueue) pull_front() (int, error){
	if deck.size() == 0{
		return -1, errors.New("Fila vazia.")
	}
	if deck.size() == 1{
		val := deck.head.value
		deck.head, deck.tail = nil, nil
		deck.inserted--
		return val, nil
	}
	val := deck.head.value
	deck.head = deck.head.next
	deck.head.prev = nil
	deck.inserted--
	return val, nil
}

func (deck *DoubleLinkedQueue) pull_back() (int, error){
	if deck.size() == 0{
		return -1, errors.New("Fila vazia.")
	}
	if deck.size() == 1{
		val := deck.tail.value
		deck.head, deck.tail = nil, nil
		deck.inserted--
		return val, nil
	}
	val := deck.tail.value
	deck.tail = deck.tail.prev
	deck.tail.next = nil
	deck.inserted--
	return val, nil
}

func (deck *DoubleLinkedQueue) size() int {
	return deck.inserted
}

func (deck *DoubleLinkedQueue) Front() (int, error) {
	if deck.size() == 0{
		return -1, errors.New("Fila vazia.")
	}
	return deck.head.value, nil
}

func (deck *DoubleLinkedQueue) back() (int, error) {
	if deck.size() == 0{
		return -1, errors.New("Fila vazia.")
	}
	return deck.tail.value, nil
}

func (deck *DoubleLinkedQueue) Print() {
	aux := deck.head
	for i:=0; i<deck.size(); i++ {
		fmt.Printf("%d ", aux.value)
		aux = aux.next
	}
	fmt.Println()
}

func main() {
	deck := DoubleLinkedQueue{}

	fmt.Println("=== PUSH_BACK ===")
	deck.push_back(10)
	deck.push_back(20)
	deck.push_back(30)
	deck.Print()

	fmt.Println("=== PUSH_FRONT ===")
	deck.push_front(5)
	deck.Print()

	fmt.Println("=== FRONT / BACK ===")
	frontVal, _ := deck.Front()
	backVal, _ := deck.back()
	fmt.Println("Front:", frontVal, "Back:", backVal)

	fmt.Println("=== PULL_FRONT ===")
	val, _ := deck.pull_front()
	fmt.Println("Removido do front:", val)
	deck.Print()

	fmt.Println("=== PULL_BACK ===")
	val, _ = deck.pull_back()
	fmt.Println("Removido do back:", val)
	deck.Print()

	fmt.Println("=== PUSH / PULL misto ===")
	deck.push_back(40)
	deck.push_front(2)
	deck.push_back(50)
	deck.Print()

	for deck.size() > 0 {
		val, _ := deck.pull_front()
		fmt.Println("Removendo do front:", val)
		deck.Print()
	}

	_, err := deck.pull_front()
	if err != nil {
		fmt.Println("Erro esperado ao remover deque vazia:", err)
	}
}