package main

import (
	"fmt"
)

func bin_search(val int, list []int, start int, end int) int {
	if start > end{
		return -1
	}	
	mid := (end+start)/2
	if list[mid] == val{
		return mid
	} else {
		if val < list[mid] {
			return bin_search(val, list, start, mid-1)
		} else {
			return bin_search(val, list, mid+1, end)
		}
	}
}

func rev_bin_search(val int, list []int, start int, end int) int {
	if start > end {
		return -1
	}
	mid := (start+end)/2
	if list[mid] == val {
		return mid
	} else {
		if val > list[mid] {
			return rev_bin_search(val, list, start, mid-1)
		} else {
			return rev_bin_search(val, list, mid+1, end)
		}
	}
}


func main() {
	asc := []int{1, 3, 5, 7, 9, 11}
	desc := []int{20, 17, 15, 10, 8, 5, 2}

	fmt.Println("=== Testes vetor crescente ===")
	fmt.Println("Busca 1  ->", bin_search(1, asc, 0, len(asc)-1))   // extremo esquerdo
	fmt.Println("Busca 11 ->", bin_search(11, asc, 0, len(asc)-1))  // extremo direito
	fmt.Println("Busca 7  ->", bin_search(7, asc, 0, len(asc)-1))   // meio
	fmt.Println("Busca 4  ->", bin_search(4, asc, 0, len(asc)-1))   // inexistente
	fmt.Println("Busca em vazio ->", bin_search(5, []int{}, 0, -1)) // vetor vazio

	fmt.Println("\n=== Testes vetor decrescente ===")
	fmt.Println("Busca 20 ->", rev_bin_search(20, desc, 0, len(desc)-1)) // extremo esquerdo
	fmt.Println("Busca 2  ->", rev_bin_search(2, desc, 0, len(desc)-1))  // extremo direito
	fmt.Println("Busca 10 ->", rev_bin_search(10, desc, 0, len(desc)-1)) // meio
	fmt.Println("Busca 18 ->", rev_bin_search(18, desc, 0, len(desc)-1)) // inexistente
	fmt.Println("Busca em vazio ->", rev_bin_search(5, []int{}, 0, -1))  // vetor vazio
}
