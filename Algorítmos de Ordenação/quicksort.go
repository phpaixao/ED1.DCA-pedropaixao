package main

import (
	"fmt"
	"math/rand"
	"time"
)

func QuickSort(v [] int, ini int, fim int) {
	if ini > fim {return}
	indexPivot := Partition(v, ini, fim)
	QuickSort(v, ini, indexPivot - 1)
	QuickSort(v, indexPivot + 1, fim)
}

func Partition(v [] int, ini int, fim int) int {
	source := rand.NewSource(time.Now().UnixNano())
	// Criar um gerador de números aleatórios com a fonte
	r := rand.New(source)

	// Gerar um valor aleatório entre ini e fim (inclusive)
	randIndex := r.Intn(fim-ini+1) + ini

	// Trocar os elementos v[randIndex] e v[fim]
	v[randIndex], v[fim] = v[fim], v[randIndex]
	
	pivot := v[fim]
	pIndex := ini
	var i int
	for i = ini; i<fim; i++{
		if v[i] <= pivot {
			v[pIndex], v[i] = v[i], v[pIndex]
			pIndex++
		} 
	}
	v[pIndex], v[i] = v[i], v[pIndex]
}

func main(){

}
