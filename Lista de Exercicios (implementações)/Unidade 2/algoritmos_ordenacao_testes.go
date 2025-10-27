package main

import (
	"fmt"
	"math"
	"time"
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

func QuickSortSemRandomizacao(v []int, ini int, fim int) {
	if ini >= fim {
		return
	}
	indexPivot := PartitionSemRandomizacao(v, ini, fim)
	QuickSortSemRandomizacao(v, ini, indexPivot-1)
	QuickSortSemRandomizacao(v, indexPivot+1, fim)
}

func PartitionSemRandomizacao(v []int, ini int, fim int) int {
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

func generateSlice(n int, disposition string) []int {
	vetor := make([]int, n)
	for i := 0; i<n; i++ {
		vetor[i] = rand.Intn(n*10)
	}

	switch disposition {
	case "aleatorio":
		return vetor
	case "crescente":
		for i := 0; i<n; i++ {
			vetor[i] = i
		}
		return vetor
	case "decrescente":
		for i := 0; i<n; i++ {
			vetor[i] = n - i - 1
		}
		return vetor
	case "counting_bom":
		for i := 0; i<n; i++ {
			vetor[i] = rand.Intn(n)
		}
		return vetor
	case "counting_ruim":
		for i := 0; i<n; i++ {
			vetor[i] = rand.Intn(n)
		}
		vetor[n/2] = 50000000
		return vetor
	default:
		return vetor
	}
}

func runTest(label string, data []int, sortFunc func([]int)) {
	copia := make([]int, len(data))
	copy(copia, data)
	fmt.Printf("Rodando: %-45s", label)
	start := time.Now()
	sortFunc(copia)
	duration := time.Since(start)
	fmt.Printf(" -> Duração: %v\n", duration)
}

func runTestQuick(label string, data []int, sortFunc func([]int, int, int)) {
	copia := make([]int, len(data))
	copy(copia, data)
	fmt.Printf("Rodando: %-45s", label)
	start := time.Now()
	sortFunc(copia, 0, len(copia)-1)
	duration := time.Since(start)
	fmt.Printf(" -> Duração: %v\n", duration)
}

func runTestCounting(label string, data []int) {
	copia := make([]int, len(data))
	copy(copia, data)
	fmt.Printf("Rodando: %-45s", label)
	start := time.Now()
	_ = CountingSort(copia)
	duration := time.Since(start)
	fmt.Printf(" -> Duração: %v\n", duration)
}

func main() {
	// Semeia o gerador aleatório UMA VEZ
	rand.Seed(time.Now().UnixNano())

	fmt.Println("--- INICIANDO EXPERIMENTOS ---")

	// --- FATO 1: Vetores pequenos (n=50) ---
	fmt.Println("\n--- FATO 1: Vetores pequenos (n=50) ---")
	const nPequeno = 50
	slicePequenoAleatorio := generateSlice(nPequeno, "aleatorio")
	slicePequenoBom := generateSlice(nPequeno, "counting_bom")

	runTest("SelectionSort (O(n^2))", slicePequenoAleatorio, SelectionSort)
	runTest("MergeSort (O(n log n))", slicePequenoAleatorio, MergeSort)
	runTestCounting("CountingSort (O(n+k), k=n)", slicePequenoBom)
	fmt.Println("(Resultado: Todos os tempos são muito baixos e parecidos)")

	// --- FATO 2, 3, 4: Vetores grandes (n=10000) e Disposição ---
	const nGrande = 10000 // 10 mil elementos
	fmt.Printf("\n--- FATOS 2, 3, 4: Vetores grandes (n=%d) ---\n", nGrande)

	sliceGrandeAleatorio := generateSlice(nGrande, "aleatorio")
	sliceGrandeCrescente := generateSlice(nGrande, "crescente")
	sliceGrandeDecrescente := generateSlice(nGrande, "decrescente")

	// FATO 2: InsertionSort (O(n^2) vs O(n))
	fmt.Println("\n--- FATO 2: Disposição afeta O(n^2) (InsertionSort) ---")
	runTest("InsertionSort (Aleatório)", sliceGrandeAleatorio, InsertionSort)
	runTest("InsertionSort (Crescente - Melhor Caso)", sliceGrandeCrescente, InsertionSort)
	runTest("InsertionSort (Decrescente - Pior Caso)", sliceGrandeDecrescente, InsertionSort)
	fmt.Println("(Resultado: Melhor Caso (crescente) é quase instantâneo, Pior Caso (decrescente) é MUITO lento)")

	// FATO 3: MergeSort (Sempre O(n log n))
	fmt.Println("\n--- FATO 3: Disposição NÃO afeta O(n log n) (MergeSort) ---")
	runTest("MergeSort (Aleatório)", sliceGrandeAleatorio, MergeSort)
	runTest("MergeSort (Crescente)", sliceGrandeCrescente, MergeSort)
	runTest("MergeSort (Decrescente)", sliceGrandeDecrescente, MergeSort)
	fmt.Println("(Resultado: Todos os tempos são rápidos e muito parecidos)")

	// FATO 4: QuickSort (Pior Caso vs Randomização)
	fmt.Println("\n--- FATO 4: QuickSort (Pior Caso vs Randomização) ---")
	runTestQuick("QuickSort (Sem Randomização, Pior Caso)", sliceGrandeCrescente, QuickSortSemRandomizacao)
	runTestQuick("QuickSort (Com Randomização)", sliceGrandeCrescente, QuickSort)
	fmt.Println("(Resultado: Sem randomização é lento (O(n^2)), Com randomização é rápido (O(n log n)))")

	// --- FATO 5: CountingSort (Bom vs Ruim - k vs n) ---
	fmt.Println("\n--- FATO 5: CountingSort (k vs n) ---")
	sliceCountingBom := generateSlice(nGrande, "counting_bom")
	sliceCountingRuim := generateSlice(nGrande, "counting_ruim")

	runTestCounting("CountingSort (Bom, k=n)", sliceCountingBom)
	runTestCounting("CountingSort (Ruim, k=50M)", sliceCountingRuim)
	fmt.Println("(Resultado: Caso 'Bom' (k=n) é muito rápido. Caso 'Ruim' (k grande) é EXTREMAMENTE lento ou falha)")

	fmt.Println("\n--- EXPERIMENTOS CONCLUÍDOS ---")
}