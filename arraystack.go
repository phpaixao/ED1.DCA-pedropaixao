package main

import (
	"errors"
	"fmt"
)

type IStack interface {
	Push(value int)    //Feito
	Pop() (int, error) //Feito
	Top() (int, error) //Feito
	IsEmpty() bool	   //Feito
	Size() int         //Feito
}

type ArrayStack struct {
	top int
	vetor [] int
}

func (pilha *ArrayStack) doubleV(){
	newV := make([] int, len(pilha.vetor)*2)
	for i := 0; i<pilha.Size(); i++{
		newV[i] = pilha.vetor[i]
	}
	pilha.vetor = newV
}

func (pilha *ArrayStack) Push(value int) {
	if pilha.Size() == len(pilha.vetor){
		pilha.doubleV
	}
	pilha.vetor[pilha.top] = value
	pilha.top++
}

func (pilha *ArrayStack) Pop() (int, error) {
	if pilha.IsEmpty() {
		return -1, errors.New("Lista vazia.")
	}
	pilha.top--
	return pilha.vetor[pilha.top], nil
}

func (pilha *ArrayStack) Top() (int, error) {
	if pilha.IsEmpty() {
		return -1, errors.New("Lista vazia.")
	}
	return pilha.vetor[pilha.top-1], nil	
}

func (pilha *ArrayStack) IsEmpty() {
	return pilha.top == 0
}

func (pilha *ArrayStack) Size() {
	return pilha.top
}