package main

import (
	"fmt"
	"math"
)

// Implementação Out-Of-Place
func SelectionSort(v []int) []int {
	newV := make([]int, len(v))
	max := math.MaxInt32
	for i := 0; i < len(v); i++ {
		menor := 0
		for j := 1; j < len(v); j++ {
			if v[j] < v[menor] {
				menor = j
			}
		}
		newV[i] = v[menor]
		v[menor] = max
	}
	return newV
}

// Implementação In-Place
func SelectionSortIP(v []int) {
	for i := 0; i < len(v)-1; i++ {
		menor := i
		for j := i+1; j < len(v); j++ {
			if v[j] < v[menor] {
				menor = j
			}
		}
		v[i], v[menor] = v[menor], v[i]
	}
}

func main() {
	// Testanto ordenação out of place
	vetor := []int{2, 1, 0, 4, 3, 10, 1, 4, 2}
	fmt.Printf("%d", SelectionSort(vetor))
	fmt.Println()

	// Testando ordenação in place
	vetor2 := []int{2, 1, 0, 4, 3, 10, 1, 4, 2}
	SelectionSortIP(vetor2)
	fmt.Printf("%d", vetor2)
}
