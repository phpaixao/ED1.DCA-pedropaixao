package main

import(
	"fmt"
)

func BubbleSort(v []int) {
	for i := 0; i<len(v)-1; i++{
		flag := false
		for j := 0; j<(len(v)-i-1); j++{
			if v[j] > v[j+1]{
				v[j], v[j+1] = v[j+1], v[j]
				flag = true
			}
		}
		// Teste pra saber quantas varreduras foram feitas
		fmt.Println("Executado!") 
		if !flag {return}
	}
}

func main(){
	vector := [] int{2, 1, 0, 4, 3, 10, 1, 4, 2}
	BubbleSort(vector)
	fmt.Printf("%d", vector)
}