package main

import (
	"fmt"
	"math"
)

func SelectionSort(v []int){
	var i_menor int
	for i := 0; i < len(v)-1; i++ {
		i_menor = i
		for j := i+1; j < len(v); j++ {
			if v[j] < v[i_menor] {
				i_menor = j
			} 
		}
		v[i], v[i_menor] = v[i_menor], v[i]
	}
}

func BubbleSort(v []int){
	for i := 0; i<len(v)-1; i++{
		trocou := false
		for j := 0; j < len(v)-1-i; j++ {
			if v[j] > v[j+1] {
				trocou = true
				v[j], v[j+1] = v[j+1], v[j]
			}
		}
		if !trocou {return}
	}
}

func InsertionSort(v []int){
	for i := 1; i < len(v); i++ {
		for j := i ; j > 0; j-- {
			if v[j] < v[j-1] {
				v[j], v[j-1] = v[j-1], v[j]
			} else {
				break
			}
		} 
	}
}

func merge(v []int, e []int, d []int) {
	index_e := 0
	index_d := 0
	i := 0

	for index_e < len(e) && index_d < len(d) {
		if e[index_e] <= d[index_d] {
			v[i] = e[index_e]
			index_e++
		} else {
			v[i] = d[index_d]
			index_d++
		}
		i++
	}

	for index := index_e; index < len(e); index++ {
		v[i] = e[index]
		i++
	}

	for index := index_d; index < len(d); index++ {
		v[i] = d[index]
		i++
	} 
}

func MergeSort(v[] int) {
	// Se o vetor for de tam = 1, encerro a recursividade
	if len(v) <= 1 {
		return
	}

	// se len(v) = 5, mid = 2 (e[] tem 2 elementos)
	// d[] tem 3 elementos (len(v)-mid: 5-2 = 3)
	mid := len(v)/2
	e := make([]int, mid)
	d := make([]int, len(v)-mid)
	
	// e[] é preenchido de v[0] até v[mid-1]
	for i := 0; i < len(e); i++ {
		e[i] = v[i]
	}

	// d[] é preenchido de v[mid] até v[len(v)-1]
	for i := 0; i < len(d); i++ {
		d[i] = v[i+mid]
	}

	// Divido os vetores recursivamente até terem tam = 1
	MergeSort(e)
	MergeSort(d)

	// Quando os vetores finalmente tiverem tam = 1,
	// Faço o merge em v[] dos vetores e[] e d[]
	merge(v, e, d)
}

func QuickSort(v []int, ini int, fim int) {
	if ini >= fim {
		return
	}
	indexPivot := Partition(v, ini, fim)
	QuickSort(v, ini, indexPivot-1)
	QuickSort(v, indexPivot+1, fim)
}

func Partition(v []int, ini int, fim int) int {
	randIndex := rand.Intn(fim-ini+1) + ini

	v[randIndex], v[fim] = v[fim], v[randIndex]

	pivot := v[fim]
	pIndex := ini
	var i int
	for i = ini; i < fim; i++ {
		if v[i] < pivot {
			v[i], v[pIndex] = v[pIndex], v[i]
			pIndex++
		}
	}
	v[pIndex], v[i] = v[i], v[pIndex]
	return pIndex
}

func CountingSort(v []int) []int {
	if len(v) == 0 {
		return []int{}
	}
	// Encontro o maior e o menor valor do vetor
	menor := v[0]
	maior := v[0]
	for i := 1; i<len(v); i++ {
		if v[i] > maior {
			maior = v[i]
		}
		if v[i] < menor {
			menor = v[i]
		}
	}

	// Criação do vetor de contagem, que abrange
	// todos os números no intervalo {menor, maior}
	c := make([]int, maior-menor+1)
	
	// Preenchendo o vetor de contagem
	for i := 0; i<len(v); i++{
		c[v[i] - menor]++
	}

	// Soma cumulativa
	for i := 0; i<len(c)-1; i++{
		c[i+1] += c[i]
	}

	//Criando o novo vetor que será o ordenado
	ord := make([]int, len(v))

	//Colocando os elementos na ordem correta
	/*
	for i := 0; i<len(v); i++{
		ord[c[v[i]-menor]-1] = v[i]
		c[v[i]-menor]--
	}
	*/

	// Versao mais estável
	for i := len(v)-1; i >= 0; i-- {
		ord[c[v[i]-menor]-1] = v[i]
		c[v[i]-menor]--
	}

	// Retorna o novo vetor ordenado
	return ord
}

func main() {
	// Semeia o 'rand' global UMA VEZ para o QuickSort
	rand.Seed(time.Now().UnixNano())

	// Função helper para obter um slice novo
	getOriginal := func() []int {
		return []int{5, 2, 8, 1, 9, 4, 7, 3, 6, 0}
	}

	// --- Teste 1: SelectionSort ---
	fmt.Println("--- Testando SelectionSort ---")
	data := getOriginal()
	fmt.Println("Original:", data)
	SelectionSort(data)
	fmt.Println("Ordenado:", data)

	// --- Teste 2: BubbleSort ---
	fmt.Println("\n--- Testando BubbleSort ---")
	data = getOriginal()
	fmt.Println("Original:", data)
	BubbleSort(data)
	fmt.Println("Ordenado:", data)

	// --- Teste 3: InsertionSort ---
	fmt.Println("\n--- Testando InsertionSort ---")
	data = getOriginal()
	fmt.Println("Original:", data)
	InsertionSort(data)
	fmt.Println("Ordenado:", data)

	// --- Teste 4: MergeSort ---
	fmt.Println("\n--- Testando MergeSort ---")
	data = getOriginal()
	fmt.Println("Original:", data)
	MergeSort(data)
	fmt.Println("Ordenado:", data)

	// --- Teste 5: QuickSort ---
	fmt.Println("\n--- Testando QuickSort ---")
	data = getOriginal()
	fmt.Println("Original:", data)
	QuickSort(data, 0, len(data)-1)
	fmt.Println("Ordenado:", data)

	// --- Teste 6: CountingSort (NOVO TESTE) ---
	fmt.Println("\n--- Testando CountingSort ---")
	data = getOriginal()
	fmt.Println("Original:", data)
	// CountingSort retorna um NOVO slice
	ordenadoData := CountingSort(data)
	fmt.Println("Ordenado:", ordenadoData)
	fmt.Println("Original (não deve mudar):", data)

	// --- Teste 7: Casos de Borda (Era o 6) ---
	fmt.Println("\n--- Testando Casos de Borda ---")

	// Fatias de teste de borda
	emptySlice := []int{}
	singleElementSlice := []int{42}
	reversedSlice := []int{10, 9, 8, 7, 6, 5}

	// Criando cópias para testar (para MergeSort e QuickSort)
	emptyCopy := make([]int, len(emptySlice))
	copy(emptyCopy, emptySlice)
	singleCopy := make([]int, len(singleElementSlice))
	copy(singleCopy, singleElementSlice)
	reversedCopy := make([]int, len(reversedSlice))
	copy(reversedCopy, reversedSlice)

	// Testando MergeSort (que lida bem com cópias)
	MergeSort(emptyCopy)
	MergeSort(singleCopy)
	MergeSort(reversedCopy)
	fmt.Println("Vazio (MergeSort):", emptyCopy)
	fmt.Println("Um Elemento (MergeSort):", singleCopy)
	fmt.Println("Invertido (MergeSort):", reversedCopy)

	// Recriando cópias para QuickSort
	emptyCopy = make([]int, len(emptySlice))
	copy(emptyCopy, emptySlice)
	singleCopy = make([]int, len(singleElementSlice))
	copy(singleCopy, singleElementSlice)
	reversedCopy = make([]int, len(reversedSlice))
	copy(reversedCopy, reversedSlice)

	// Testando QuickSort
	QuickSort(emptyCopy, 0, len(emptyCopy)-1)
	QuickSort(singleCopy, 0, len(singleCopy)-1)
	QuickSort(reversedCopy, 0, len(reversedCopy)-1)
	fmt.Println("Vazio (QuickSort):", emptyCopy)
	fmt.Println("Um Elemento (QuickSort):", singleCopy)
	fmt.Println("Invertido (QuickSort):", reversedCopy)

	// Testando CountingSort (retorna novo slice, não precisa de cópia)
	fmt.Println("--- Casos de Borda (CountingSort) ---")
	resultadoVazio := CountingSort(emptySlice)
	resultadoUnico := CountingSort(singleElementSlice)
	resultadoInvertido := CountingSort(reversedSlice)
	fmt.Println("Vazio (CountingSort):", resultadoVazio)
	fmt.Println("Um Elemento (CountingSort):", resultadoUnico)
	fmt.Println("Invertido (CountingSort):", resultadoInvertido)
}