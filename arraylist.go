package main

import (
	"fmt"
	"errors"
)

type List interface{
	Tamanho() int; //Feito
	Get(index int) (int, error); //Feito
	Append(value int); //Feito
	Insert(index int, value int) error; //Feito
	Remove(index int) error; //Feito
}

type ArrayList struct{
	vetor []int;
	inserted int;
}

func (lista *ArrayList) Init(size int){
	lista.vetor = make([]int, size);
}

func (lista *ArrayList) Tamanho() int{
	return lista.inserted;
}

func (lista *ArrayList) Get(index int) (int, error){
	if index < 0 || index >= lista.inserted {
		return -1, errors.New(fmt.Sprintf("Index out of range: %d", index));
	} 
	return lista.vetor[index], nil;
}

func (lista *ArrayList) doubleV() {
	newV := make([]int, len(lista.vetor)*2);
	for i := 0; i < len(lista.vetor); i++{
		newV[i] = lista.vetor[i];
	}
	lista.vetor = newV;
}

func (lista *ArrayList) Append(value int){
	if lista.inserted == len(lista.vetor){
		lista.doubleV(); 
	} 
	lista.vetor[lista.inserted] = value;
	lista.inserted++;
}

func (lista *ArrayList) Insert(index int, value int) error {
	if index < 0 || index > lista.inserted{
		return errors.New(fmt.Sprintf("Index out of range: %d", index));
	}
	if lista.inserted == len(lista.vetor){
		lista.doubleV();
	}
	for i := lista.inserted; i>index; i--{
		lista.vetor[i] = lista.vetor[i-1];
	}
	lista.vetor[index] = value;
	lista.inserted++;
	return nil;
}

func (lista *ArrayList) Remove(index int) error {
	if index < 0  || index >= lista.inserted{
		return errors.New(fmt.Sprintf("Index out of range: %d", index));
	}
	for i := index; i < lista.inserted-1; i++{
		lista.vetor[i] = lista.vetor[i+1];
	}
	lista.inserted--;
	lista.vetor[lista.Tamanho()] = 0;
	return nil;
}

func (lista *ArrayList) Print() {
	for i := 0; i<lista.Tamanho(); i++{
		fmt.Printf("%d ", lista.vetor[i]);
	}
	fmt.Println();
}

func main(){
	l := &ArrayList{}
	l.Init(10);
	for i := 0; i<50; i++{
		l.Append(i);
	}
	l.Print();

	l.Insert(1, 100);

	l.Print();

	l.Remove(1);

	l.Print();
}