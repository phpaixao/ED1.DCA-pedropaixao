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
	Init()
}

Type Tuple struct {
	key int
	value string
}

type HashTable struct {
	buckets [][]Tuple
	elementsInserted int
}