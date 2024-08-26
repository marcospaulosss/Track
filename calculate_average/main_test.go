package main

import (
	"testing"
)

func TestCalculateAverage(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected float64
	}{
		{"Empty slice", []int{}, 0},
		{"Single element", []int{10}, 10},
		{"Multiple elements", []int{10, 20, 30, 40, 50}, 30},
		{"Negative numbers", []int{-10, -20, -30, -40, -50}, -30},
		{"Mixed numbers", []int{-10, 0, 10, 20, 30}, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculateAverage(tt.numbers)
			if result != tt.expected {
				t.Errorf("calculateAverage(%v) = %v; want %v", tt.numbers, result, tt.expected)
			}
		})
	}
}
