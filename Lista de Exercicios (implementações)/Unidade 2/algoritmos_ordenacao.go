package main

import (
	"fmt"
	"math"
)

func SelectionSort(v []int){
	var i_menor int
	for i := 0; i < len(v)-1; i++ {
		i_menor = 0
		for j := i+1; j < len(v); j++ {
			if v[i] > v[j] {
				i_menor = j
			} 
		}
		v[i], v[i_menor] = v[i_menor], v[i]
	}
}

func BubbleSort(v []int){
	for i := 0; i<len(v)-1; i++{
		trocou := false
		for j := 0; j < len(v)-1-i; j++ {
			if v[i] > v[i+1] {
				trocou = true
				v[i], v[i+1] = v[i+1], v[i]
			}
		}
		if !trocou {return}
	}
}

func InsertionSort(v []int){
	for i := 1; i < len(v); i++ {
		for j := i ; j > 0; j-- {
			if v[j] < v[j-1] {
				v[j], v[j-1] = v[j-1], v[j]
			} else {
				break
			}
		} 
	}
}

func merge(v []int, e []int, d []int) {
	index_e := 0
	index_d := 0
	i := 0

	for index_e < len(e) && index_d < len(d) {
		if e[index_e] <= d[index_d] {
			v[i] = e[index_e]
			index_e++
		} else {
			v[i] = d[index_d]
			index_d++
		}
		i++
	}

	for index := index_e; index < len(e); index++ {
		v[i] = e[index]
		i++
	}

	for index := index_d; index < len(d); index++ {
		v[i] = d[index]
		i++
	} 
}

func MergeSort(v[] int) {
	// Se o vetor for de tam = 1, encerro a recursividade
	if len(v) == 1 {
		return
	}

	// se len(v) = 5, mid = 2 (e[] tem 2 elementos)
	// d[] tem 3 elementos (len(v)-mid: 5-2 = 3)
	mid := len(v)/2
	e := make([]int, mid)
	d := make([]int, len(v)-mid)
	
	// e[] é preenchido de v[0] até v[mid-1]
	for i := 0; i < len(e); i++ {
		e[i] = v[i]
	}

	// d[] é preenchido de v[mid] até v[len(v)-1]
	for i := 0; i < len(d); i++ {
		d[i] = v[i+mid]
	}

	// Divido os vetores recursivamente até terem tam = 1
	MergeSort(e)
	MergeSort(d)

	// Quando os vetores finalmente tiverem tam = 1,
	// Faço o merge em v[] dos vetores e[] e d[]
	merge(v, e, d)
}

func main() {
	vetor := []int{2, 1, 0, 4, 3, 10, 1, 4, 2}
	MergeSort(vetor)
	fmt.Println(vetor)
}