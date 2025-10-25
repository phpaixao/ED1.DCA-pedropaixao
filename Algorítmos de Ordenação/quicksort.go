package main

import (
	"fmt"
	"math/rand"
	"time"
)

func QuickSort(v [] int, ini int, fim int) {
	if ini >= fim {return}
	indexPivot := Partition(v, ini, fim)
	QuickSort(v, ini, indexPivot - 1)
	QuickSort(v, indexPivot + 1, fim)
}

func Partition(v [] int, ini int, fim int) int {
	// Gerar um valor aleatório entre ini e fim (inclusive)
	randIndex := rand.Intn(fim-ini+1) + ini

	// Trocar os elementos v[randIndex] e v[fim]
	v[randIndex], v[fim] = v[fim], v[randIndex]
	
	pivot := v[fim]
	pIndex := ini
	var i int
	for i = ini; i<fim; i++{
		if v[i] <= pivot {
			v[pIndex], v[i] = v[i], v[pIndex]
			pIndex++
		} 
	}
	v[pIndex], v[i] = v[i], v[pIndex]
	return pIndex
}

// --- FUNÇÃO MAIN PARA TESTE ---
func main() {
	// 1. Inicializar a semente de rand para evitar resultados repetitivos na aleatoriedade do Partition
	// Embora a função Partition tente fazer isso, é boa prática fazer aqui.
	rand.Seed(time.Now().UnixNano())

	// 2. Definir o slice de teste
	data := []int{5, 2, 8, 1, 9, 4, 7, 3, 6, 0}
	
	// Criar uma cópia para demonstração
	dataCopy := make([]int, len(data))
	copy(dataCopy, data)
	
	fmt.Println("Slice Original:", dataCopy)

	// 3. Chamar o QuickSort
	// ini é 0, fim é len(data) - 1
	QuickSort(dataCopy, 0, len(dataCopy)-1)

	// 4. Imprimir o resultado
	fmt.Println("Slice Ordenado:", dataCopy)

	// --- Segundo Teste (Caso de Borda) ---
	emptySlice := []int{}
	QuickSort(emptySlice, 0, -1) // -1 é o 'fim' correto para um slice vazio
	fmt.Println("\nSlice Vazio:", emptySlice)

	singleElementSlice := []int{42}
	QuickSort(singleElementSlice, 0, 0)
	fmt.Println("Slice com 1 Elemento:", singleElementSlice)

	reversedSlice := []int{10, 9, 8, 7, 6, 5}
	fmt.Println("\nSlice Invertido Original:", reversedSlice)
	QuickSort(reversedSlice, 0, len(reversedSlice)-1)
	fmt.Println("Slice Invertido Ordenado:", reversedSlice)
}