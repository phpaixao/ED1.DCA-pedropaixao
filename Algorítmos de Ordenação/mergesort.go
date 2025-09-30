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


func main() {
	v := []int{9, 4, 3, 6, 3, 2, 5, 7, 1, 8}
	e := []int{3, 3, 4, 6, 9}
	d := []int{1, 2, 5, 7, 8}

	merge(v, e, d)
	fmt.Printf("%d", v)
}
