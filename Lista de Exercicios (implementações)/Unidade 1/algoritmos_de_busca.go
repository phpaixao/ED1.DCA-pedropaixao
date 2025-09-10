package main

import (
	"fmt"
)

func bin_search(val int, list []int, start int, end int) int {
	if start > end {
		return start
	}
	mid := (end + start) / 2
	if list[mid] == val {
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
		return start
	}
	mid := (start + end) / 2
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
	fmt.Println("Busca 1   ->", bin_search(1, asc, 0, len(asc)-1))  // posição do 1
	fmt.Println("Busca 11  ->", bin_search(11, asc, 0, len(asc)-1)) // posição do 11
	fmt.Println("Busca 7   ->", bin_search(7, asc, 0, len(asc)-1))  // posição do 7
	fmt.Println("Busca 4   ->", bin_search(4, asc, 0, len(asc)-1))  // deveria ficar entre 3 e 5
	fmt.Println("Busca 0   ->", bin_search(0, asc, 0, len(asc)-1))  // antes do primeiro
	fmt.Println("Busca 12  ->", bin_search(12, asc, 0, len(asc)-1)) // depois do último
	fmt.Println("Busca em vazio ->", bin_search(5, []int{}, 0, -1)) // vetor vazio

	fmt.Println("\n=== Testes vetor decrescente ===")
	fmt.Println("Busca 20  ->", rev_bin_search(20, desc, 0, len(desc)-1)) // posição do 20
	fmt.Println("Busca 2   ->", rev_bin_search(2, desc, 0, len(desc)-1))  // posição do 2
	fmt.Println("Busca 10  ->", rev_bin_search(10, desc, 0, len(desc)-1)) // posição do 10
	fmt.Println("Busca 18  ->", rev_bin_search(18, desc, 0, len(desc)-1)) // deveria ficar entre 20 e 17
	fmt.Println("Busca 25  ->", rev_bin_search(25, desc, 0, len(desc)-1)) // antes do primeiro
	fmt.Println("Busca 0   ->", rev_bin_search(0, desc, 0, len(desc)-1))  // depois do último
	fmt.Println("Busca em vazio ->", rev_bin_search(5, []int{}, 0, -1))   // vetor vazio
}
