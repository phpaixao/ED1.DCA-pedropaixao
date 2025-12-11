package main

import "errors"

type BinaryHeap struct {
	v                []int
	elementsInserted int
}

func (heap *BinaryHeap) IndexOfParent(index int) int {
	if index == 0 {
		return -1
	}
	return (index - 1) / 2
}

func (heap *BinaryHeap) indexOfChildren(index int) (int, int) {
	return 2*index + 1, 2*index + 2
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
	if smallest == index {
		return
	}

	heap.v[index], heap.v[smallest] = heap.v[smallest], heap.v[index]
	heap.bubbleDown(smallest)
}

func (heap *BinaryHeap) bubbleUp(index int) {
	for {
		parent := heap.IndexOfParent(index)
		if parent < 0 {
			return
		}
		if heap.v[parent] <= heap.v[index] {
			return
		}

		heap.v[parent], heap.v[index] = heap.v[index], heap.v[parent]
		index = parent
	}
}

func (heap *BinaryHeap) Add(value int) {
	if heap.elementsInserted == len(heap.v) {
		heap.v = append(heap.v, value)
	} else {
		heap.v[heap.elementsInserted] = value
	}

	heap.elementsInserted++
	heap.bubbleUp(heap.elementsInserted - 1)
}

func (heap *BinaryHeap) Poll() (int, error) {
	if heap.elementsInserted == 0 {
		return -1, errors.New("Heap já está vazia.")
	}

	removed := heap.v[0]
	heap.v[0] = heap.v[heap.elementsInserted-1]
	heap.elementsInserted--
	heap.bubbleDown(0)
	return removed, nil
}

func (heap *BinaryHeap) Remove(e int) error {
	if heap.elementsInserted == 0 {
		return errors.New("Heap já está vazia.")
	}

	index := -1
	for i := 0; i < heap.elementsInserted; i++ {
		if heap.v[i] == e {
			index = i
			break
		}
	}

	if index == -1 {
		return errors.New("Elemento não existe na heap.")
	}

	old := heap.v[index]
	heap.v[index] = heap.v[heap.elementsInserted-1]
	heap.elementsInserted--

	if heap.v[index] < old {
		heap.bubbleUp(index)
	} else {
		heap.bubbleDown(index)
	}

	return nil
}
