package main

import "fmt"

func printCountNegativeNonMultiple3Result() {
	matrix := [][]int{
		{1, -2, 3},
		{4, -5, 6},
		{7, -8, 9},
		{-10, 11, -12},
		{13, 14, -15},
	}

	fmt.Println("Task 5 - variant 14")
	fmt.Println("Test case")
	PrintMatrix(matrix)
	fmt.Println("Result: ")
	fmt.Print(countNegativeNonMultiple3(matrix))
}

func countNegativeNonMultiple3(matrix [][]int) int {
	count := 0
	for i := 4; i < 5; i++ { // accessing the 5th row
		for _, val := range matrix[i] {
			if val < 0 && val%3 != 0 {
				count++
			}
		}
	}
	return count
}

func PrintMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%4d", matrix[i][j])
		}
		fmt.Println()
	}
}
