package main

import (
	"slices"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	var tests = []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Неотсортированный массив",
			input:    []int{40, 15, 16, 150, 100, 0, 43, 77},
			expected: []int{0, 15, 16, 40, 43, 77, 100, 150},
		},
		{
			name:     "Отсортированный массив",
			input:    []int{0, 15, 16, 40, 43, 77, 100, 150},
			expected: []int{0, 15, 16, 40, 43, 77, 100, 150},
		},
		{
			name:     "Массив с повторяющимися значениями",
			input:    []int{40, 15, 16, 15, 100, 0, 43, 40},
			expected: []int{0, 15, 15, 16, 40, 40, 43, 100},
		},
		{
			name:     "Пустой массив",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Массив с одним элементом",
			input:    []int{10},
			expected: []int{10},
		},
		{
			name:     "Массив с двумя элементами",
			input:    []int{20, 10},
			expected: []int{10, 20},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			bubbleSort(tt.input)

			if !slices.Equal(tt.input, tt.expected) {
				t.Errorf("BubbleSort = %v; expected %v", tt.input, tt.expected)
			}
		})
	}
}

func bubbleSort(arr []int) {
	len := len(arr) - 1
	for i := 0; i < len; i++ {
		for j := 0; j < len-i; j++ {

			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
}
