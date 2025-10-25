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

func main() {
	vetor := []int{2, 1, 0, 4, 3, 10, 1, 4, 2}
	InsertionSort(vetor)
	fmt.Println(vetor)
}