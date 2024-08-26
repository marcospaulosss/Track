package main

import (
	"fmt"
)

func calculateAverage(numbers []int) float64 {
	if len(numbers) == 0 {
		return 0
	}
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return float64(sum) / float64(len(numbers)) // Conversão do sum e len(numbers) para float64
}

func main() {
	nums := []int{10, 20, 30, 40, 50}
	average := calculateAverage(nums)
	fmt.Println("Average:", average)
}
