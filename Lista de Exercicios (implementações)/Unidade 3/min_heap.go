package main

// Implementação de um Min-heap

import (
	"errors"
)

type BinaryHeap struct {
	v [] int
	elementsInserted int
}

func (heap *BinaryHeap) IndexOfParent(index int) int {
	if index == 0 {return -1}
	return (index-1)/2
}

func (heap *BinaryHeap) indexOfChildren(index int) (int, int) {
	left_child := 2*index + 1
	right_child := 2*index + 2
	return left_child, right_child
}

func (heap *BinaryHeap) bubbleDown(index int) {
	left, right := heap.indexOfChildren(index)
	smallest := index
	if left < heap.elementsInserted && heap.v[left] < heap.v[smallest] {
		smallest = left
	}
	if right < heap.elementsInserted && heap.v[right] < heap.v[smallest] {
		smallest = right
	}
	if smallest == index {return}
	heap.v[index], heap.v[smallest] = heap.v[smallest], heap.v[index]
	heap.bubbleDown(smallest)
}

func (heap *BinaryHeap) bubbleUp(index int) {
	parent := heap.IndexOfParent(index)
	if parent < 0 {
		return
	}
	if heap.v[parent] <= heap.v[index] {
		return
	}
	heap.v[parent], heap.v[index] = heap.v[index], heap.v[parent]
	heap.bubbleUp(parent)
}

func (heap *BinaryHeap) Add(value int) {
	if heap.elementsInserted == len(heap.v) {
		heap.v = append(heap.v, value)
	} else {
		heap.v[heap.elementsInserted] = value
	}
	heap.elementsInserted++
	heap.bubbleUp(heap.elementsInserted-1)
}

func (heap *BinaryHeap) Poll() (int, error) {
	if heap.elementsInserted == 0 {
		return -1, errors.New("Heap already empty.")
	}
	last := heap.elementsInserted-1
	removed := heap.v[0]
	heap.v[0] = heap.v[last]
	heap.elementsInserted--
	heap.bubbleDown(0)
	return removed, nil
}

func (heap *BinaryHeap) Remove(e int) error {
	if heap.elementsInserted == 0 {
		return errors.New("Heap already empty.")
	}
	index := -1
	for i := 0; i<heap.elementsInserted; i++ {
		if heap.v[i] == e {
			index = i
			break
		}
	}
	if index == -1 {
		return errors.New("No such element in the Heap.")
	}
	last := heap.elementsInserted-1
	heap.v[index] = heap.v[last]
	heap.elementsInserted--
	heap.bubbleUp(index)
	heap.bubbleDown(index)
	return nil
}

func main() {
	// Pequeno teste para validar
	h := &BinaryHeap{}
	h.Add(10)
	h.Add(5)
	h.Add(20)
	h.Add(2) // 2 deve subir para a raiz

	val, _ := h.Poll()
	fmt.Println("Removido (Min):", val) // Deve ser 2
    
    // Validando o estado atual
    fmt.Println("Próximo Poll deve ser 5. Removendo...")
    val, _ = h.Poll()
    fmt.Println("Removido:", val) // Deve ser 5
}