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

type ArrayDequeue struct {
	v [] int
	front int
	rear int
}

func (deck *ArrayDequeue) init(size int) {
	deck.v = make([] int, size)
	deck.front, deck.rear = -1, -1
} 

func (deck *ArrayDequeue) doubleV() {
	newV := make([] int, len(deck.v)*2)
	for i:=deck.front; i<(len(deck.v) + deck.front); i++{
		newV[i-deck.front] = deck.v[i % len(deck.v)]
	}
	deck.v = newV
	deck.front = 0
	deck.rear = deck.size()-1
}

func (deck *ArrayDequeue) push_front(val int){
	if deck.size() == len(deck.v){
		deck.doubleV()
	}
	if deck.size() == 0{
		deck.front++
		deck.rear++
	} else {
		if deck.front - 1 == -1{
			deck.front = len(deck.v)-1
		} else {
			deck.front--
		}
	}
	deck.v[deck.front] = val
}

func (deck *ArrayDequeue) push_back(val int){
	if deck.size() == len(deck.v){
		deck.doubleV()
	}
	if deck.size() == 0{
		deck.front++
		deck.rear++
	} else {
		deck.rear = (deck.rear+1) % len(deck.v)
	}
	deck.v[deck.rear] = val
}

func (deck *ArrayDequeue) pull_front() (int, error){
	if deck.size() == 0{
		return -1, errors.New("Lista vazia.")
	}
	val := deck.v[deck.front]
	if deck.front == deck.rear {
		deck.rear, deck.front = -1, -1
	} else {
		deck.front = (deck.front+1) % len(deck.v)
	}
	return val, nil
}

func (deck *ArrayDequeue) pull_back() (int, error){
	if deck.size() == 0{
		return -1, errors.New("Lista vazia.")
	}
	if deck.rear == deck.front{
		val := deck.v[deck.rear]
		deck.front = -1
		deck.rear = -1

		return val, nil
	}
	val := deck.v[deck.rear]
	if deck.rear - 1 == -1 {
		deck.rear = len(deck.v) - 1
	} else {
		deck.rear--
	}
	return val, nil
}

func (deck *ArrayDequeue) size() int {
	if deck.rear == -1 && deck.front == -1{
		return 0
	}
	if deck.rear >= deck.front {
		return deck.rear - deck.front + 1
	}
	return len(deck.v) - deck.front + deck.rear + 1 
}

func (deck *ArrayDequeue) Front() (int, error) {
	if deck.size() == 0{
		return -1, errors.New("Lista vazia.")
	}
	return deck.v[deck.front], nil
}

func (deck *ArrayDequeue) back() (int, error) {
	if deck.size() == 0{
		return -1, errors.New("Lista vazia.")
	}
	return deck.v[deck.rear], nil
}

func (deck *ArrayDequeue) Print(){
	for i:=deck.front; i<(deck.size()+deck.front); i++{
		fmt.Printf("%d ", deck.v[i % len(deck.v)])
	}
	fmt.Println()
}

func main() {
	deck := ArrayDequeue{}
	deck.init(3) // inicializa com capacidade pequena para testar crescimento

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