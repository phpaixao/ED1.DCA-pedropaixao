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
	return (fila.rear - fila.front + len(fila.v)) % len(fila.v) + 1
}

func (fila *ArrayQueue) Front() (int, error){
	if fila.front == -1 && fila.rear == -1 {
		return -1, errors.New("A lista está vazia.")
	}
	return fila.v[fila.front], nil
}