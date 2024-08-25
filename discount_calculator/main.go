package main

import "fmt"

func calculateDiscount(price float64, discount float64) float64 {
	if discount > 1 || discount < 0 {
		return 0 // Corrigido para retornar 0 se o desconto for inválido
	}
	return price - price*discount
}

func main() {
	// Testes esperados:
	// Discounted price (expected 80): 80
	// Discounted price (expected 0): 0
	// Discounted price (expected 0): 0
	// Discounted price (expected 100): 100
	fmt.Println("Discounted price (expected 80):", calculateDiscount(100, 0.2))
	fmt.Println("Discounted price (expected 0):", calculateDiscount(100, -0.2))
	fmt.Println("Discounted price (expected 0):", calculateDiscount(100, -1))
	fmt.Println("Discounted price (expected 100):", calculateDiscount(100, 0))
}
