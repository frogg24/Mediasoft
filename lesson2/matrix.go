package main

import (
	"math/rand"
)

func generateMatrix(m int, n int) [][]int {
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	for i := range matrix {
		for j := range matrix[i] {
			for {
				temp := rand.Intn(1000)
				if !isInMatrix(matrix, temp) {
					matrix[i][j] = temp
					break
				}
			}
		}
	}
	return matrix
}

func isInMatrix(matrix [][]int, number int) bool {
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == number {
				return true
			}
		}
	}
	return false
}
