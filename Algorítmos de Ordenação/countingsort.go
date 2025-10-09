package main

import(
	"fmt"
)

func CountingSort(v []int) []int{
	// Passo 1 -> N
	menor, maior := v[0], v[0]
	for i := 1; i<len(v); i++{
		if v[i] < menor{ 
			menor = v[i]
		}
		if v[i] > maior{
			maior = v[i]
		}
	}

	// Passo 2
	// Vetor de contagem
	c := make([]int, maior - menor + 1)

	// Passo 3 -> N
	for i := 0; i<len(v); i++{
		c[v[i] - menor]++
	}

	// Passo 4 -> K
	for i := 0; i<len(c)-1; i++{
		c[i+1] += c[i]
	}

	//Passo 5 
	ord := make([]int, len(v))

	//Passo 6 -> N
	for i := 0; i<len(v); i++{
		ord[c[v[i]-menor]-1] = v[i]
		c[v[i]-menor]--
	}

	// O(3N + K) --> O(N + K)
	// K = maior - menor + 1
	// Algorítmo é bom quando (K <= N) (Eh bom que soh)
	return ord
}

func main() {
	// Exemplo de vetor para ordenar
	v := []int{5, 2, 9, 1, 5, 6}

	// Exibindo o vetor original
	fmt.Println("Vetor original:", v)

	// Chamando o CountingSort
	ordenado := CountingSort(v)

	// Exibindo o vetor ordenado
	fmt.Println("Vetor ordenado:", ordenado)
}