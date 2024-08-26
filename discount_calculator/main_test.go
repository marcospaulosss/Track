package main

import (
	"testing"
)

func TestCalculateDiscount(t *testing.T) {
	tests := []struct {
		name     string
		price    float64
		discount float64
		expected float64
	}{
		{"Valid discount", 100, 0.2, 80},
		{"Negative discount", 100, -0.2, 0},
		{"Discount greater than one", 100, 1.5, 0},
		{"Zero discount", 100, 0, 100},
		{"Full discount", 100, 1, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculateDiscount(tt.price, tt.discount)
			if result != tt.expected {
				t.Errorf("calculateDiscount(%v, %v) = %v; want %v", tt.price, tt.discount, result, tt.expected)
			}
		})
	}
}
