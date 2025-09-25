package main

import(
	"fmt"
)

func InsertionSort(v []int) {
	for i := 1; i < len(v); i++{
		for j := i; j > 0; j--{
			if v[j] < v[j-1]{
				v[j], v[j-1] = v[j-1], v[j]
			} else {
				break
			}
		}
	}
}

func main(){
	vector := [] int{2, 1, 0, 4, 3, 10, 1, 4, 2}
	InsertionSort(vector)
	fmt.Printf("%d", vector)
}