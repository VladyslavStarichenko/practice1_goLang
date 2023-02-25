package main

import (
	"fmt"
)

func printSumNonZeroSecondColumn(matrix [][]int) {
	fmt.Println("Task5 - variant 18")
	fmt.Println("Matrix:")
	PrintMatrix(matrix)
	fmt.Printf("Sum of non-zero elements in the second column: %d\n", sumNonZeroSecondColumn(matrix))
}

func sumNonZeroSecondColumn(matrix [][]int) int {
	sum := 0
	for _, row := range matrix {
		if len(row) >= 2 && row[1] != 0 {
			sum += row[1]
		}
	}
	return sum
}
