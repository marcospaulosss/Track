package main

import (
	"testing"
)

func TestTotalSum(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected int
	}{
		{name: "Positive numbers", numbers: []int{1, 2, 3, 4, 5, 6}, expected: 21},
		{name: "Including zero", numbers: []int{0, 1, 2, 3, 4, 5}, expected: 15},
		{name: "Single element", numbers: []int{10}, expected: 10},
		{name: "Empty slice", numbers: []int{}, expected: 0},
		{name: "Negative numbers", numbers: []int{-1, -2, -3, -4}, expected: -10},
		{name: "Mixed numbers", numbers: []int{-1, 0, 1}, expected: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := totalSum(tt.numbers)
			if result != tt.expected {
				t.Errorf("totalSum(%v) = %d; expected %d", tt.numbers, result, tt.expected)
			}
		})
	}
}
