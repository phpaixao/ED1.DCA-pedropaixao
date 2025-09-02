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
	front() (int, error)
	back() (int, error)
}

type ArrayDequeue struct {
	v [] int
	front int
	rear int
	inserted int
}

func (deck *ArrayDequeue) push_front(val int){
	if deck.front == -1 && deck.rear == -1{
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
	deck.inserted++
}

func (deck *ArrayDequeue) push_back(val int){
	if deck.front == -1 && deck.rear == -1{
		deck.front++
		deck.rear++
	} else {
		deck.rear = (deck.rear+1) % len(deck.v)
	}
	deck.v[deck.rear] = val
	deck.inserted++
}

func (deck *ArrayDequeue) pull_front() (int, error){
	if deck.front == -1 && deck.rear == -1{
		return -1, errors.New("Lista vazia.")
	}
	val := deck.v[deck.front]
	if deck.front == deck.rear {
		deck.rear, deck.front = -1, -1
	} else {
		deck.front = (deck.front+1) % len(deck.v)
	}
	deck.inserted--
	return val, nil
}

func (deck *ArrayDequeue) pull_back() (int, error){
	if deck.front == -1 && deck.rear == -1{
		return -1, errors.New("Lista vazia.")
	}
	val := deck.v[deck.front]
	
}