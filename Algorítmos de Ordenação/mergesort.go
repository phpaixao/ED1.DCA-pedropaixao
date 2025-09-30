package main

import (
	"fmt"
)

func merge(v []int, e []int, d []int) {
	i := 0
	index_e := 0
	index_d := 0
	for index_e < len(e) && index_d < len(d){
		if e[index_e] < d[index_d] {
			v[i] = e[index_e]
			index_e++
		} else {
			v[i] = d[index_d]
			index_d++
		}
		i++
	}

	for index_e < len(e) {
		v[i] = e[index_e]
		index_e++
		i++
	}

	for index_d < len(d) {
		v[i] = d[index_d]
		index_d++
		i++
	}
}

func MergeSort(v []int) {
	if len(v) > 1{
		meio := len(v)/2
		e := make([] int, meio)
		d := make([] int, len(v) - meio)

		i := 0
		for index_e := 0; index_e < len(e); index_e++{
			e[index_e] = v[i]
			i++
		}

		for index_d := 0; index_d < len(d); index_d++{
			d[index_d] = v[i]
			i++
		}

		MergeSort(e)
		MergeSort(d)
		merge(v, e, d)
	}
}

func main() {
	v := []int{9,4,3,6,3,2,5,7,1,8}
	
	MergeSort(v)
	
	fmt.Printf("%d", v)
}
