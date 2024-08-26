package main

import (
	"fmt"
	"sync"
)

func calculateSum(nums []int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for _, num := range nums {
		sum += num
	}
	results <- sum
}

func totalSum(numbers []int) int {
	results := make(chan int, len(numbers))
	var wg sync.WaitGroup
	chunkSize := len(numbers) / 2

	// Adição de uma condicional que garante que chunkSize nunca seja zero. Isso garante que sempre haja pelo menos uma goroutine para processar os números.
	// isso evita travamentos nos testes que passam um slice vazio.
	if chunkSize == 0 {
		chunkSize = 1
	}

	for i := 0; i < len(numbers); i += chunkSize {
		end := i + chunkSize
		if end > len(numbers) {
			end = len(numbers)
		}
		wg.Add(1)
		go calculateSum(numbers[i:end], results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	totalSum := 0
	for sum := range results {
		totalSum += sum
	}

	return totalSum /// Foi removido a divisão por sum o que retornava uma média ao invés da soma total
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	result := totalSum(nums)
	fmt.Println("Total Sum:", result) // Esperado: 21 (1 + 2 + 3 + 4 + 5 + 6)
}
