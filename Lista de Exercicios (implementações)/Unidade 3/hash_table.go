package main

import (
	"fmt"
	"errors"
)

type Map interface {
	Put(key int, value string)
	Get(key int) (string, error)
	Remove(key int) error
	Size() int
	LoadFactor() float32
	Init(size int)
}

type Tuple struct {
	key int
	value string
}

type HashTable struct {
	buckets [][]Tuple
	elementsInserted int
}

var emptyTuple = Tuple{}

func put(buckets [][]Tuple, key int, value string) [][]Tuple {
	bucket := key % len(buckets)
	for i, tuple := range buckets[bucket] {
		if tuple.key == key {
			buckets[bucket][i].value = value
			return buckets 	
		}
	}
	buckets[bucket] = append(buckets[bucket], Tuple{key, value})
	return buckets
}

func (table *HashTable) increaseBuckets() {
	increasedBuckets := make([][]Tuple, 2*len(table.buckets))
	for i := 0; i<len(table.buckets); i++ {
		for _, tuple := range table.buckets[i] {
			put(increasedBuckets, tuple.key, tuple.value)
		}
	}
	table.buckets = increasedBuckets
}

func (table *HashTable) Put(key int, value string) {
	if table.LoadFactor() > 0.75 {
		table.increaseBuckets()
	}
	contido := false
	bucket := key % len(table.buckets)
	for _, tuple := range table.buckets[bucket] {
		if tuple.key == key {
			contido = true
			break
		}
	}
	table.buckets = put(table.buckets, key, value)
	if !contido {
		table.elementsInserted++
	}
}

func (table *HashTable) Get(key int) (string, error) {
	bucket := key % len(table.buckets)
	for _, tuple := range table.buckets[bucket] {
		if tuple.key == key {
			return tuple.value, nil
		}
	}
	return "", errors.New("A tabela não possui a determinada chave")
}

func (table *HashTable) Remove(key int) error {
	bucket := key % len(table.buckets)
	index := -1
	for i, tuple := range table.buckets[bucket] {
		if tuple.key == key {
			index = i
			break
		}
	}
	if index == -1 {
		return errors.New("A tabela não possui a determinada chave")
	} else {
		table.buckets[bucket] = append(table.buckets[bucket][:index], table.buckets[bucket][index+1:]...)
		table.elementsInserted--
		return nil
	}
}

func (table *HashTable) LoadFactor() float32 {
	return float32(table.elementsInserted) / float32(len(table.buckets))
}

func (table *HashTable) Size() int {
	return table.elementsInserted
}

func (table *HashTable) Init(size int) {
	table.buckets = make([][]Tuple, size)
	table.elementsInserted = 0
}

func main() {
	ht := &HashTable{}
	ht.Init(2) // Começa pequeno para testar o resize

	fmt.Println("Put(1, A)")
	ht.Put(1, "A")
	fmt.Printf("Size: %d (Esp: 1)\n", ht.Size())

	fmt.Println("Put(1, B) - Atualização")
	ht.Put(1, "B")
	fmt.Printf("Size: %d (Esp: 1) -> Não deve aumentar!\n", ht.Size())

	v, _ := ht.Get(1)
	fmt.Println("Get(1):", v, "(Esp: B)")

	fmt.Println("Put(2, C)")
	ht.Put(2, "C")
	
	fmt.Println("Put(3, D) - Deve causar Resize")
	ht.Put(3, "D")
	
	fmt.Printf("Size Final: %d\n", ht.Size())
	fmt.Printf("LoadFactor: %.2f\n", ht.LoadFactor())
}